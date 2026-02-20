package car

import (
	"context"
	"performance_tracker_v2_be/core"
	"performance_tracker_v2_be/db/main-db/models"
	"performance_tracker_v2_be/helpers"

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

func (service *Service) GetAllCars(
	context context.Context,
	pagination *core.Pagination,
	sort *core.Sort,
	filters *GetCarsFilter,
) ([]models.Car, error) {
	rows, err := core.GetEntityList(core.GetEntityListPayload{
		Pool:         service.Pool,
		Context:      context,
		TableName:    "cars",
		Pagination:   pagination,
		Sort:         sort,
		Filters:      helpers.StructToQueryFilters(filters),
		FilterRules:  service.GetFilters(),
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

func (service *Service) GetTotalCarsCount(context context.Context, filters *GetCarsFilter) (int64, error) {
	return core.GetEntityCount(core.GetEntityCountPayload{
		Pool:        service.Pool,
		Context:     context,
		TableName:   "cars",
		Filters:     helpers.StructToQueryFilters(filters),
		FilterRules: service.GetFilters(),
	})
}

func (service *Service) GetCarByID(context context.Context, id int) (*models.Car, error) {
	car := service.Pool.QueryRow(context, `
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

func (service *Service) CreateCar(context context.Context, payload *CreateCarRequest) (int, error) {
	result := service.Pool.QueryRow(context, `
		INSERT INTO cars (name, image, description)
		VALUES ($1, $2, $3) returning id
	`, payload.Name, payload.Image, payload.Description)

	var id int

	if err := result.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (service *Service) UpdateCar(context context.Context, id int, payload *UpdateCarRequest) error {
	updates := core.StructToMap[core.OptionalBodyField[interface{}]](payload)

	return core.UpdateEntity(core.UpdateEntityByIdPayload{
		ID:      id,
		Pool:    service.Pool,
		Context: context,
		Updates: updates,
		Table:   "cars",
	})
}

func (service *Service) DeleteCar(context context.Context, id int) error {
	_, err := service.Pool.Query(context, `
		delete from cars where id = $1
	`, id)

	return err
}
