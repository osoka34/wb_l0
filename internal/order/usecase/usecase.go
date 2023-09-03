package usecase

import (
	"fmt"
	"go.uber.org/zap"
	"sync"
	"sync/atomic"
	"wb_l0/config"
	"wb_l0/internal/cconstants"
	"wb_l0/internal/delivery"
	"wb_l0/internal/item"
	"wb_l0/internal/models"
	"wb_l0/internal/order"
	"wb_l0/internal/payment"
	"wb_l0/pkg/utils"
)

type OrderUsecase struct {
	orderRepo    order.Repository
	paymentRepo  payment.Repository
	deliveryRepo delivery.Repository
	itemRepo     item.Repository
	logger       *zap.SugaredLogger
	cfg          *config.Config
	mu           sync.Mutex
	atmc         *atomic.Value
}

func NewOrderUsecase(orderRepo order.Repository, paymentRepo payment.Repository, deliveryRepo delivery.Repository, itemRepo item.Repository, logger *zap.SugaredLogger, cfg *config.Config) order.Usecase {
	return &OrderUsecase{
		orderRepo: orderRepo, paymentRepo: paymentRepo, deliveryRepo: deliveryRepo, itemRepo: itemRepo, logger: logger, cfg: cfg, mu: sync.Mutex{},
	}
}

func (o *OrderUsecase) CreateOrder(params *models.OrderModel) (models.Response, error) {
	var (
		err error
		//order    models.OrderModel = *params
		response models.Response
	)

	err = o.orderRepo.Insert(&models.InsertParams{
		SqlValues: utils.StructToSqlArray(*params),
	})
	if err != nil {
		response.Success = false
		response.Description = fmt.Sprint(err)
		response.ErrCode = cconstants.CantInsertOrder
		return response, err
	}
	err = o.paymentRepo.Insert(&models.InsertParams{
		SqlValues: utils.StructToSqlArray(*(params.Payment)),
	})
	if err != nil {
		response.Success = false
		response.Description = fmt.Sprint(err)
		response.ErrCode = cconstants.CantInsertPayment
		return response, err
	}
	err = o.deliveryRepo.Insert(&models.InsertParams{
		SqlValues: utils.StructToSqlArray(*(params.Delivery)),
	})
	if err != nil {
		response.Success = false
		response.Description = fmt.Sprint(err)
		response.ErrCode = cconstants.CantInsertDelivery
		return response, err
	}
	for _, val := range *params.Items {
		err = o.itemRepo.Insert(&models.InsertParams{
			SqlValues: utils.StructToSqlArray(val),
		})
		if err != nil {
			response.Success = false
			response.Description = fmt.Sprint(err)
			response.ErrCode = cconstants.CantInsertItem
			return response, err
		}
	}

	response.Success = true
	response.Description = "Order created"
	response.ErrCode = cconstants.Ok

	o.mu.Lock()
	m := o.atmc.Load().(map[string]models.OrderModel)
	m[params.OrderUid] = *params
	o.atmc.Store(m)
	o.mu.Unlock()

	return response, nil
}

func (o *OrderUsecase) LoadCache() {
	var (
		atmc     atomic.Value
		orderMap map[string]models.OrderModel = make(map[string]models.OrderModel)
		mu       sync.Mutex
		wg       sync.WaitGroup
	)

	orders, err := o.orderRepo.SelectAll()
	if err != nil {
		o.logger.Errorf("can't initialize OrderUsecase, because of no connection to db")
		return
	}
	wg.Add(len(*orders))
	for _, ordr := range *orders {
		ordr := ordr
		go func(order models.OrderModel) {
			defer wg.Done()
			payment, err := o.paymentRepo.Select(&models.SelectParams{
				ordr.OrderUid,
			})
			if err != nil {
				o.logger.Errorf("error while loading payment for order %s", ordr.OrderUid)
				return
			}
			ordr.Payment = payment
			delivery, err := o.deliveryRepo.Select(&models.SelectParams{
				ordr.OrderUid,
			})
			if err != nil {
				o.logger.Errorf("error while loading delivery for order %s", ordr.OrderUid)
				return
			}
			ordr.Delivery = delivery
			items, err := o.itemRepo.SelectAll(&models.SelectParams{
				ordr.OrderUid,
			})
			if err != nil {
				o.logger.Errorf("error while loading items for order %s", ordr.OrderUid)
				return
			}
			ordr.Items = items

			mu.Lock()
			orderMap[ordr.OrderUid] = ordr
			mu.Unlock()
		}(ordr)
	}

	wg.Wait()

	atmc.Store(orderMap)

	o.mu.Lock()
	o.atmc = &atmc
	o.mu.Unlock()
}

func (o *OrderUsecase) GetOrder(params *models.GetParams) (*models.Response, error) {
	var response models.Response

	m := o.atmc.Load().(map[string]models.OrderModel)
	order, ok := m[params.OrderUid]
	if !ok {
		response.Success = false
		response.Description = "Order not found"
		response.ErrCode = cconstants.OrderNotFound
		return &response, fmt.Errorf("order not found")
	}
	response.Success = true
	response.ErrCode = cconstants.Ok
	response.Data = order
	return &response, nil
}
