package delivery

import "wb_l0/internal/models"

type Repository interface {
	Select(params *models.SelectParams) (*models.DeliveryModel, error)
	Insert(params *models.InsertParams) error
	SelectAll() (*[]models.DeliveryModel, error)
}
