package repository

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jmoiron/sqlx"
	"wb_l0/internal/cconstants"
	"wb_l0/internal/delivery"
	"wb_l0/internal/models"
)

type postgresRepository struct {
	poolDb *pgxpool.Pool
	db     *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB, poolDb *pgxpool.Pool) delivery.Repository {
	return &postgresRepository{db: db, poolDb: poolDb}
}

func (p *postgresRepository) Insert(params *models.InsertParams) error {
	var (
		query = `
			INSERT INTO %[1]s
				(order_uid, name, phone, zip, city, address, region, email)
			VALUES
				%[2]s;
			`

		values []any = []any{
			cconstants.DeliveryDB, params.SqlValues,
		}
	)

	query = fmt.Sprintf(query, values...)

	if _, err := p.db.Exec(query); err != nil {
		return err
	}
	return nil
}

func (p *postgresRepository) Select(params *models.SelectParams) (*models.DeliveryModel, error) {
	var (
		delivery models.DeliveryModel
		query    = `
			SELECT
				name, phone, zip, city, address, region, email, order_uid
			FROM
				%[1]s
			WHERE
				order_uid = '%[2]s';
		`
		values []any = []any{
			cconstants.DeliveryDB, params.OrderUid,
		}
	)

	query = fmt.Sprintf(query, values...)

	if err := p.db.Get(&delivery, query); err != nil {
		return nil, err
	}

	return &delivery, nil
}

func (p *postgresRepository) SelectAll() (*[]models.DeliveryModel, error) {
	var (
		delivery []models.DeliveryModel
		query    = `
			SELECT
				name, phone, zip, city, address, region, email, order_uid
			FROM
				%[1]s;
		`
		values []any = []any{
			cconstants.DeliveryDB,
		}
	)

	query = fmt.Sprintf(query, values...)

	if err := p.db.Select(&delivery, query); err != nil {
		return nil, err
	}

	return &delivery, nil
}
