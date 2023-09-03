package order

import "wb_l0/internal/models"

type Repository interface {
	Insert(params *models.InsertParams) error
	Select(params *models.SelectParams) (*models.OrderModel, error)
	SelectAll() (*[]models.OrderModel, error)
}
