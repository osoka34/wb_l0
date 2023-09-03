package payment

import "wb_l0/internal/models"

type Repository interface {
	Select(params *models.SelectParams) (*models.PaymentModel, error)
	Insert(params *models.InsertParams) error
	SelectAll() (*[]models.PaymentModel, error)
}
