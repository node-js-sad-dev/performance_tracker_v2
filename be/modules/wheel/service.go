package wheel

import (
	"context"
	"performance_tracker_v2_be/core/handler"
	"performance_tracker_v2_be/core/service"
	"performance_tracker_v2_be/db/main-db/models"
	"performance_tracker_v2_be/helpers"

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
) ([]models.Wheel, error) {
	rows, err := service.GetEntityList(handler.GetEntityListPayload{
		Pool:        s.Pool,
		Context:     ctx,
		TableName:   "wheels",
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

	result := make([]models.Wheel, 0)
	for rows.Next() {
		var wheel models.Wheel
		if err := rows.Scan(&wheel.ID, &wheel.Name, &wheel.IsDefault, &wheel.CreatedAt); err != nil {
			return nil, err
		}
		result = append(result, wheel)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Service) GetTotalCount(ctx context.Context, filters *GetFilters) (int64, error) {
	return service.GetEntityCount(handler.GetEntityCountPayload{
		Pool:        s.Pool,
		Context:     ctx,
		TableName:   "wheels",
		Filters:     helpers.StructToMap[[]string](filters),
		FilterRules: s.GetFilters(),
	})
}

func (s *Service) GetByID(ctx context.Context, id int64) (*models.Wheel, error) {
	row := s.Pool.QueryRow(ctx, `
		SELECT id, name, is_default, created_at
		FROM wheels
		WHERE id = $1
	`, id)

	var result models.Wheel
	if err := row.Scan(&result.ID, &result.Name, &result.IsDefault, &result.CreatedAt); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *Service) Create(ctx context.Context, payload *CreateRequest) (int64, error) {
	var id int64

	err := pgx.BeginFunc(ctx, s.Pool, func(tx pgx.Tx) error {
		if payload.IsDefault {
			if _, err := tx.Exec(ctx, `
				UPDATE wheels
				SET is_default = FALSE
				WHERE is_default = TRUE
			`); err != nil {
				return err
			}
		}

		return tx.QueryRow(ctx, `
			INSERT INTO wheels (name, is_default)
			VALUES ($1, $2)
			RETURNING id
		`, payload.Name, payload.IsDefault).Scan(&id)
	})
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Service) UpdateByID(ctx context.Context, id int64, payload *UpdateRequestParsed) error {
	return pgx.BeginFunc(ctx, s.Pool, func(tx pgx.Tx) error {
		var entityFromDB models.Wheel
		if err := tx.QueryRow(ctx, `
			SELECT id, name, is_default, created_at
			FROM wheels
			WHERE id = $1
		`, id).Scan(&entityFromDB.ID, &entityFromDB.Name, &entityFromDB.IsDefault, &entityFromDB.CreatedAt); err != nil {
			return err
		}

		if payload.IsDefault.GetIsSet() && !payload.IsDefault.GetIsNull() && payload.IsDefault.GetValue() == true && !entityFromDB.IsDefault {
			if _, err := tx.Exec(ctx, `
				UPDATE wheels
				SET is_default = FALSE
				WHERE is_default = TRUE AND id != $1
			`, id); err != nil {
				return err
			}
		}

		updates := helpers.StructToMap[handler.IOptionalField](payload)

		return service.UpdateEntity(handler.UpdateEntityByIdPayload{
			ID:       id,
			Executor: tx,
			Context:  ctx,
			Updates:  updates,
			Table:    "wheels",
		})
	})
}

func (s *Service) DeleteByID(ctx context.Context, id int64) error {
	_, err := s.Pool.Exec(ctx, `
		DELETE FROM wheels
		WHERE id = $1
	`, id)

	return err
}
