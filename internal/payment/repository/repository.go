package repository

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jmoiron/sqlx"
	"wb_l0/internal/cconstants"
	"wb_l0/internal/models"
	"wb_l0/internal/payment"
)

type postgresRepository struct {
	poolDb *pgxpool.Pool
	db     *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB, poolDb *pgxpool.Pool) payment.Repository {
	return &postgresRepository{db: db, poolDb: poolDb}
}

func (p *postgresRepository) Insert(params *models.InsertParams) error {
	var (
		query = `
			INSERT INTO %[1]s
				(order_uid, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
			VALUES
				%[2]s;
			`
		values []any = []any{
			cconstants.PaymentDB, params.SqlValues,
		}
	)

	query = fmt.Sprintf(query, values...)

	if _, err := p.db.Exec(query); err != nil {
		return err
	}
	return nil
}

func (p *postgresRepository) Select(params *models.SelectParams) (*models.PaymentModel, error) {
	var (
		payment models.PaymentModel
		query   = `
			SELECT
				order_uid, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee
			FROM
				%[1]s
			WHERE
				order_uid = '%[2]s';
		`
		values []any = []any{
			cconstants.PaymentDB, params.OrderUid,
		}
	)

	query = fmt.Sprintf(query, values...)

	if err := p.db.Get(&payment, query); err != nil {
		return nil, err
	}

	return &payment, nil
}

func (p *postgresRepository) SelectAll() (*[]models.PaymentModel, error) {
	var (
		payment []models.PaymentModel
		query   = `
			SELECT
				order_uid, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee
			FROM
				%[1]s;
		`
		values []any = []any{
			cconstants.PaymentDB,
		}
	)

	query = fmt.Sprintf(query, values...)

	if err := p.db.Select(&payment, query); err != nil {
		return nil, err
	}

	return &payment, nil
}
