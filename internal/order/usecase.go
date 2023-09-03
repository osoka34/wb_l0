package order

import "wb_l0/internal/models"

type Usecase interface {
	CreateOrder(params *models.OrderModel) (models.Response, error)
	LoadCache()
	GetOrder(params *models.GetParams) (*models.Response, error)
}
