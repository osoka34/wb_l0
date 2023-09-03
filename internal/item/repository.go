package item

import "wb_l0/internal/models"

type Repository interface {
	Insert(params *models.InsertParams) error
	Select(params *models.SelectParams) (*[]models.ItemModel, error)
	SelectAll(params *models.SelectParams) (*[]models.ItemModel, error)
}
