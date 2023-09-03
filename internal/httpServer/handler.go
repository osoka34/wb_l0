package httpServer

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	serverLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
	"os"
	"time"
	"wb_l0/internal/cconstants"
	"wb_l0/internal/models"
	"wb_l0/pkg/storage"
	"wb_l0/pkg/utils"

	deliveryRepository "wb_l0/internal/delivery/repository"
	itemRepository "wb_l0/internal/item/repository"
	paymentRepository "wb_l0/internal/payment/repository"

	orderHttp "wb_l0/internal/order/controller"
	orderRepository "wb_l0/internal/order/repository"
	orderUsecase "wb_l0/internal/order/usecase"

	natsCustom "wb_l0/internal/nats"
)

func (s *Server) MapHandlers(ctx context.Context, app *fiber.App, logger *zap.SugaredLogger) error {

	ctx, cancel := context.WithCancel(ctx)

	// -------------------------------------------------------------------------------------

	db, err := storage.InitPsqlDB(s.cfg)
	if err != nil {
		logger.Fatalf("err is: %v", err)
	}

	// -------------------------------------------------------------------------------------

	poolDb, err := storage.InitConnectionPoolPsqlDB(s.cfg)
	if err != nil {
		logger.Fatalf("err is: %v", err)
	}

	// -------------------------------------------------------------------------------------

	nts := natsCustom.NewNats(s.cfg, logger)

	// -------------------------------------------------------------------------------------

	itemRepo := itemRepository.NewPostgresRepository(db, poolDb)
	paymentRepo := paymentRepository.NewPostgresRepository(db, poolDb)
	deliveryRepo := deliveryRepository.NewPostgresRepository(db, poolDb)
	orderRepo := orderRepository.NewPostgresRepository(db, poolDb)

	// -------------------------------------------------------------------------------------

	orderUC := orderUsecase.NewOrderUsecase(orderRepo, paymentRepo, deliveryRepo, itemRepo, logger, s.cfg)

	// -------------------------------------------------------------------------------------

	orderHandlers := orderHttp.NewOrderHandler(orderUC, s.cfg, logger)

	// -------------------------------------------------------------------------------------

	app.Use(serverLogger.New())
	if _, ok := os.LookupEnv("LOCAL"); !ok {
		app.Use(recover.New())
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	// -------------------------------------------------------------------------------------

	orderHttp.MapOrderRoutes(app, orderHandlers)

	// -------------------------------------------------------------------------------------

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:

				var (
					order    *models.OrderModel
					orderUID = utils.GenerateRandomString(20)
				)

				order = utils.GenerateRandomOrder(orderUID)
				err := nts.PublishMessage(s.cfg.Nats.Topic, *order)
				if err != nil {
					logger.Errorf("error in publish message: %v", err)
					cancel()
				}

				time.Sleep(time.Second * 7)
			}
		}
	}(ctx)

	// -------------------------------------------------------------------------------------

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				var (
					order *models.OrderModel
				)

				order, err = nts.SubscribeAndReceiveMessage(s.cfg.Nats.Topic)
				if err != nil {
					logger.Errorf("error in receive message: %v", err)
					cancel()
				}
				_, err := orderUC.CreateOrder(order)
				if err != nil {
					logger.Errorf("error in create order: %v", err)
					cancel()
				}
				time.Sleep(time.Second * 6)
			}
		}
	}(ctx)

	// -------------------------------------------------------------------------------------

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				cancel()
			default:
				orderUC.LoadCache()
				time.Sleep(cconstants.GoRoutineSleepSeconds * time.Second)

			}
		}
	}(ctx)

	// -------------------------------------------------------------------------------------

	return nil
}
