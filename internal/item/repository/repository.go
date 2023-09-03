package repository

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jmoiron/sqlx"
	"wb_l0/internal/cconstants"
	"wb_l0/internal/item"
	"wb_l0/internal/models"
)

type postgresRepository struct {
	poolDb *pgxpool.Pool
	db     *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB, poolDb *pgxpool.Pool) item.Repository {
	return &postgresRepository{db: db, poolDb: poolDb}
}

func (p *postgresRepository) Insert(params *models.InsertParams) error {
	var (
		query = `
			INSERT INTO %[1]s
				(order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status)
			VALUES
				%[2]s;

			`
		values []any = []any{
			cconstants.ItemDB, params.SqlValues,
		}
	)

	query = fmt.Sprintf(query, values...)

	if _, err := p.db.Exec(query); err != nil {
		return err
	}
	return nil
}

func (p *postgresRepository) Select(params *models.SelectParams) (*[]models.ItemModel, error) {
	var (
		items []models.ItemModel
		query = `
			SELECT
				order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status
			FROM
				%[1]s
			WHERE
				order_uid = '%[2]s';
		`
		values []any = []any{
			cconstants.ItemDB, params.OrderUid,
		}
	)

	query = fmt.Sprintf(query, values...)

	if err := p.db.Select(&items, query); err != nil {
		return nil, err
	}

	return &items, nil
}

func (p *postgresRepository) SelectAll(params *models.SelectParams) (*[]models.ItemModel, error) {
	var (
		items []models.ItemModel
		query = `
			SELECT
				order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status
			FROM
				%[1]s;
		`
		values []any = []any{
			cconstants.ItemDB,
		}
	)

	query = fmt.Sprintf(query, values...)

	if err := p.db.Select(&items, query); err != nil {
		return nil, err
	}

	return &items, nil
}
