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

func (service *Service) GetCarByID(id string) {
}

func (service *Service) CreateCar(name, image, description string) {
}

func (service *Service) UpdateCar(id, name, image, description string) {
}

func (service *Service) DeleteCar(id string) {
}
