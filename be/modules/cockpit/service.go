package cockpit

import (
	"context"
	"performance_tracker_v2_be/core"
	"performance_tracker_v2_be/db/main-db/models"
	"performance_tracker_v2_be/helpers"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	Pool *pgxpool.Pool
}

func (service *Service) GetFilters() map[string]core.FilterRule {
	return map[string]core.FilterRule{
		"name": {DBColumn: "name", Operator: "ILIKE", IsFuzzy: true},
		"id":   {DBColumn: "id", Operator: "=", IsFuzzy: false},
	}
}

func (service *Service) GetList(
	ctx context.Context,
	pagination *core.Pagination,
	sort *core.Sort,
	filters *GetFilters,
) ([]models.Cockpit, error) {
	rows, err := core.GetEntityList(core.GetEntityListPayload{
		Pool:        service.Pool,
		Context:     ctx,
		TableName:   "cockpits",
		Pagination:  pagination,
		Sort:        sort,
		Filters:     helpers.StructToMap[[]string](filters),
		FilterRules: service.GetFilters(),
		SelectFields: []string{
			"id", "name", "is_default", "created_at",
		},
	})

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]models.Cockpit, 0)
	for rows.Next() {
		var cockpit models.Cockpit
		if err := rows.Scan(&cockpit.ID, &cockpit.Name, &cockpit.IsDefault, &cockpit.CreatedAt); err != nil {
			return nil, err
		}
		result = append(result, cockpit)
	}

	return result, nil
}

func (service *Service) GetTotalCount(ctx context.Context, filters *GetFilters) (int64, error) {
	return core.GetEntityCount(core.GetEntityCountPayload{
		Pool:        service.Pool,
		Context:     ctx,
		TableName:   "cockpits",
		Filters:     helpers.StructToMap[[]string](filters),
		FilterRules: service.GetFilters(),
	})
}

func (service *Service) GetById(ctx context.Context, id int) (*models.Cockpit, error) {
	cockpit := service.Pool.QueryRow(ctx, `
		SELECT id, name, is_default, created_at
		FROM cockpits
		WHERE id = $1
	`, id)

	var result models.Cockpit
	if err := cockpit.Scan(&result.ID, &result.Name, &result.IsDefault, &result.CreatedAt); err != nil {
		return nil, err
	}

	return &result, nil
}

func (service *Service) Create(ctx context.Context, request CreateRequest) (int64, error) {
	var id int64

	err := pgx.BeginFunc(ctx, service.Pool, func(tx pgx.Tx) error {
		if request.IsDefault {
			_, err := tx.Exec(ctx, `
				UPDATE cockpits
				SET is_default = FALSE
				WHERE is_default = TRUE
			`)

			if err != nil {
				return err
			}
		}

		return tx.QueryRow(ctx, `
			INSERT INTO cockpits (name, is_default)
			VALUES ($1, $2) returning id
		`, request.Name, request.IsDefault).Scan(&id)
	})

	return id, err
}

func (service *Service) UpdateById(ctx context.Context, id int64, request UpdateRequestParsed) error {
	return pgx.BeginFunc(ctx, service.Pool, func(tx pgx.Tx) error {

		return nil
	})
}
