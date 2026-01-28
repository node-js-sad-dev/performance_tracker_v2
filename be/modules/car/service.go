package car

import (
	"performance_tracker_v2_be/core"
	"performance_tracker_v2_be/db/main-db/models"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	Pool    *pgxpool.Pool
	Context *gin.Context
}

func (service *Service) GetAllCars(
	pagination *core.Pagination,
	sort *core.Sort,
	filters map[string][]string,
) ([]models.Car, error) {

	filterRules := map[string]core.FilterRule{
		"name": {DBColumn: "name", Operator: "ILIKE", IsFuzzy: true},
		"id":   {DBColumn: "id", Operator: "=", IsFuzzy: false},
	}

	rows, err := core.GetEntityList(core.GetEntityListPayload{
		Pool:         service.Pool,
		Context:      service.Context,
		TableName:    "cars",
		Pagination:   pagination,
		Sort:         sort,
		Filters:      filters,
		FilterRules:  filterRules,
		SelectFields: []string{"id", "name", "image", "description", "createdAt"},
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

func (service *Service) GetCarByID(id string) {}

func (service *Service) CreateCar(name, image, description string) {}

func (service *Service) UpdateCar(id, name, image, description string) {}

func (service *Service) DeleteCar(id string) {}
