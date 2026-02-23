package car

import (
	"context"
	"performance_tracker_v2_be/core/handler"
	"performance_tracker_v2_be/core/service"
	"performance_tracker_v2_be/db/main-db/models"
	"performance_tracker_v2_be/helpers"

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
) ([]models.Car, error) {
	rows, err := service.GetEntityList(handler.GetEntityListPayload{
		Pool:         s.Pool,
		Context:      ctx,
		TableName:    "cars",
		Pagination:   pagination,
		Sort:         sort,
		Filters:      helpers.StructToMap[[]string](filters),
		FilterRules:  s.GetFilters(),
		SelectFields: []string{"id", "name", "image", "description", "created_at"},
	})

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]models.Car, 0)
	for rows.Next() {
		var car models.Car
		if err := rows.Scan(&car.ID, &car.Name, &car.Image, &car.Description, &car.CreatedAt); err != nil {
			return nil, err
		}
		result = append(result, car)
	}

	return result, nil
}

func (s *Service) GetTotalCount(ctx context.Context, filters *GetFilters) (int64, error) {
	return service.GetEntityCount(handler.GetEntityCountPayload{
		Pool:        s.Pool,
		Context:     ctx,
		TableName:   "cars",
		Filters:     helpers.StructToMap[[]string](filters),
		FilterRules: s.GetFilters(),
	})
}

func (s *Service) GetByID(ctx context.Context, id int64) (*models.Car, error) {
	car := s.Pool.QueryRow(ctx, `
		SELECT id, name, image, description, created_at
		FROM cars
		WHERE id = $1
	`, id)

	var result models.Car

	if err := car.Scan(&result.ID, &result.Name, &result.Image, &result.Description, &result.CreatedAt); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *Service) Create(ctx context.Context, payload *CreateRequest) (int64, error) {
	result := s.Pool.QueryRow(ctx, `
		INSERT INTO cars (name, image, description)
		VALUES ($1, $2, $3) returning id
	`, payload.Name, payload.Image, payload.Description)

	var id int64

	if err := result.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Service) UpdateByID(ctx context.Context, id int64, payload *UpdateRequestParsed) error {
	updates := helpers.StructToMap[handler.IOptionalField](payload)

	return service.UpdateEntity(handler.UpdateEntityByIdPayload{
		ID:       id,
		Executor: s.Pool,
		Context:  ctx,
		Updates:  updates,
		Table:    "cars",
	})
}

func (s *Service) DeleteByID(ctx context.Context, id int64) error {
	_, err := s.Pool.Query(ctx, `
		delete from cars where id = $1
	`, id)

	return err
}
