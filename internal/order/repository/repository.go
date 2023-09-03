package repository

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jmoiron/sqlx"
	"wb_l0/internal/cconstants"
	"wb_l0/internal/models"
	"wb_l0/internal/order"
)

type postgresRepository struct {
	poolDb *pgxpool.Pool
	db     *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB, poolDb *pgxpool.Pool) order.Repository {
	return &postgresRepository{db: db, poolDb: poolDb}
}

func (p *postgresRepository) Insert(params *models.InsertParams) error {
	var (
		query = `
			INSERT INTO %[1]s
				(order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
			VALUES
				%[2]s;

			`
		values []any = []any{
			cconstants.OrdersDB, params.SqlValues,
		}
	)

	query = fmt.Sprintf(query, values...)

	if _, err := p.db.Exec(query); err != nil {
		return err
	}
	return nil
}

func (p *postgresRepository) Select(params *models.SelectParams) (*models.OrderModel, error) {
	var (
		orderModel models.OrderModel
		query      = `
			SELECT
				order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard
			FROM
				%[1]s
			WHERE
				order_uid = '%[2]s';
		`
		values []any = []any{
			cconstants.OrdersDB, params.OrderUid,
		}
	)

	query = fmt.Sprintf(query, values...)

	if err := p.db.Get(&orderModel, query); err != nil {
		return nil, err
	}

	return &orderModel, nil
}

func (p *postgresRepository) SelectAll() (*[]models.OrderModel, error) {

	var (
		orderModel []models.OrderModel
		query      = `
			SELECT
				order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard
			FROM
				%[1]s;
		`
		values []any = []any{
			cconstants.OrdersDB,
		}
	)

	query = fmt.Sprintf(query, values...)

	if err := p.db.Select(&orderModel, query); err != nil {
		return nil, err
	}

	return &orderModel, nil
}
