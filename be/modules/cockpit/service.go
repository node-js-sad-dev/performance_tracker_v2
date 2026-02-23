package cockpit

import (
	"context"
	"errors"
	"performance_tracker_v2_be/core/handler"
	"performance_tracker_v2_be/core/service"
	"performance_tracker_v2_be/db/main-db/models"
	"performance_tracker_v2_be/helpers"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	Pool *pgxpool.Pool
}

func (s *Service) GetFilters() map[string]handler.FilterRule {
	return map[string]handler.FilterRule{
		"name": {DBColumn: "name", Operator: "ILIKE", IsFuzzy: true},
		"id":   {DBColumn: "id", Operator: "=", IsFuzzy: false},
	}
}

func (s *Service) GetList(
	ctx context.Context,
	pagination *handler.Pagination,
	sort *handler.Sort,
	filters *GetFilters,
) ([]models.Cockpit, error) {
	rows, err := service.GetEntityList(handler.GetEntityListPayload{
		Pool:        s.Pool,
		Context:     ctx,
		TableName:   "cockpits",
		Pagination:  pagination,
		Sort:        sort,
		Filters:     helpers.StructToMap[[]string](filters),
		FilterRules: s.GetFilters(),
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

func (s *Service) GetTotalCount(ctx context.Context, filters *GetFilters) (int64, error) {
	return service.GetEntityCount(handler.GetEntityCountPayload{
		Pool:        s.Pool,
		Context:     ctx,
		TableName:   "cockpits",
		Filters:     helpers.StructToMap[[]string](filters),
		FilterRules: s.GetFilters(),
	})
}

func (s *Service) GetByID(ctx context.Context, id int64) (*models.Cockpit, error) {
	cockpit := s.Pool.QueryRow(ctx, `
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

func (s *Service) Create(ctx context.Context, payload *CreateRequest) (int64, error) {
	var id int64

	err := pgx.BeginFunc(ctx, s.Pool, func(tx pgx.Tx) error {
		if payload.IsDefault {
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
		`, payload.Name, payload.IsDefault).Scan(&id)
	})

	return id, err
}

func (s *Service) UpdateByID(ctx context.Context, id int64, payload *UpdateRequestParsed) error {
	return pgx.BeginFunc(ctx, s.Pool, func(tx pgx.Tx) error {
		var entityFromDb models.Cockpit

		err := tx.QueryRow(ctx, `
			SELECT id, name, is_default, created_at
			FROM cockpits
			WHERE id = $1
		`, id).Scan(&entityFromDb.ID, &entityFromDb.Name, &entityFromDb.IsDefault, &entityFromDb.CreatedAt)

		if err != nil {
			return err
		}

		if &entityFromDb == nil {
			return errors.New("cockpit not found with id: " + strconv.FormatInt(id, 10))
		}

		if payload.IsDefault.GetIsSet() && payload.IsDefault.GetValue() == true && entityFromDb.IsDefault == false {
			_, err := tx.Exec(ctx, `
				UPDATE cockpits
				SET is_default = FALSE
				WHERE is_default = TRUE AND id != $1
			`, id)

			if err != nil {
				return err
			}
		}

		updates := helpers.StructToMap[handler.IOptionalField](payload)

		return service.UpdateEntity(handler.UpdateEntityByIdPayload{
			ID:       id,
			Executor: tx,
			Context:  ctx,
			Updates:  updates,
			Table:    "cockpits",
		})
	})
}

func (s *Service) DeleteByID(ctx context.Context, id int64) error {
	_, err := s.Pool.Query(ctx, `
		delete from cockpits where id = $1
	`, id)

	return err
}
