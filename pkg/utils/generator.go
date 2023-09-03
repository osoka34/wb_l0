package utils

import (
	"math/rand"
	"wb_l0/internal/models"
)

func GenerateRandomOrder(OrderUid string) *models.OrderModel {
	return &models.OrderModel{
		OrderUid:          OrderUid,
		TrackNumber:       GenerateRandomString(10),
		Entry:             GenerateRandomString(10),
		Delivery:          GenerateRandomDelivery(OrderUid),
		Payment:           GenerateRandomPayment(OrderUid),
		Items:             GenerateRandomItems(OrderUid),
		Locale:            GenerateRandomString(10),
		InternalSignature: GenerateRandomString(10),
		CustomerId:        GenerateRandomString(10),
		DeliveryService:   GenerateRandomString(10),
		Shardkey:          GenerateRandomString(10),
		SmId:              GenerateRandomInt(),
		DateCreated:       GetMoscowTime(),
		OofShard:          GenerateRandomString(10),
	}
}

func GenerateRandomDelivery(OrderUid string) *models.DeliveryModel {
	return &models.DeliveryModel{
		OrderUid: OrderUid,
		Name:     GenerateRandomString(10),
		Phone:    GenerateRandomString(10),
		Zip:      GenerateRandomString(10),
		City:     GenerateRandomString(10),
		Address:  GenerateRandomString(10),
		Region:   GenerateRandomString(10),
		Email:    GenerateRandomString(10),
	}
}

func GenerateRandomPayment(OrderUid string) *models.PaymentModel {
	return &models.PaymentModel{
		OrderUid:     OrderUid,
		Transaction:  GenerateRandomString(10),
		RequestId:    GenerateRandomString(10),
		Currency:     GenerateRandomString(10),
		Provider:     GenerateRandomString(10),
		Amount:       GenerateRandomInt(),
		PaymentDt:    GenerateRandomInt(),
		Bank:         GenerateRandomString(10),
		DeliveryCost: GenerateRandomInt(),
		GoodsTotal:   GenerateRandomInt(),
		CustomFee:    GenerateRandomInt(),
	}
}

func GenerateRandomItems(OrderUid string) *[]models.ItemModel {
	var (
		count = rand.Intn(5)
		items = make([]models.ItemModel, 0, count)
	)

	for i := 0; i < count; i++ {
		item := models.ItemModel{
			OrderUid:    OrderUid,
			ChrtId:      GenerateRandomInt(),
			TrackNumber: GenerateRandomString(10),
			Price:       GenerateRandomInt(),
			Rid:         GenerateRandomString(10),
			Name:        GenerateRandomString(10),
			Sale:        GenerateRandomInt(),
			Size:        GenerateRandomString(10),
			TotalPrice:  GenerateRandomInt(),
			NmId:        GenerateRandomInt(),
			Brand:       GenerateRandomString(10),
			Status:      GenerateRandomInt(),
		}
		items = append(items, item)

	}
	return &items
}

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func GenerateRandomInt() int {
	return rand.Intn(100000)
}
