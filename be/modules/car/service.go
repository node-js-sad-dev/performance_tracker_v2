package car

import (
	"context"
	"fmt"
	"performance_tracker_v2_be/db/main-db/models"
	"strings"

	"performance_tracker_v2_be/core"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	Pool    *pgxpool.Pool
	Context context.Context
}

func (service *Service) GetAllCars(
	pagination *core.Pagination,
	sort *core.Sort,
	filters map[string][]string,
) ([]models.Car, error) {
	type FilterRule struct {
		DBColumn string
		Operator string
		IsFuzzy  bool
	}

	filterRules := map[string]FilterRule{
		"name": {DBColumn: "name", Operator: "ILIKE", IsFuzzy: true},
		"id":   {DBColumn: "id", Operator: "=", IsFuzzy: false},
	}

	var queryBuilder strings.Builder
	queryBuilder.WriteString(`SELECT id, name, image, description, "createdAt" FROM cars`)

	var args []interface{}
	var whereClauses []string
	argCounter := 1

	for key, values := range filters {
		rule, isAllowed := filterRules[key]

		if !isAllowed || len(values) == 0 {
			continue
		}

		var orConditions []string
		for _, val := range values {
			if val == "" {
				continue
			}

			var dbValue string

			if rule.IsFuzzy {
				dbValue = "%" + val + "%"
			} else {
				dbValue = val
			}

			orConditions = append(orConditions, fmt.Sprintf("%s %s $%d", rule.DBColumn, rule.Operator, argCounter))
			args = append(args, dbValue)
			argCounter++
		}

		if len(orConditions) > 0 {
			whereClauses = append(whereClauses, "("+strings.Join(orConditions, " OR ")+")")
		}
	}

	if len(whereClauses) > 0 {
		queryBuilder.WriteString(" WHERE " + strings.Join(whereClauses, " AND "))
	}

	sortMap := map[string]string{
		"name":      "name",
		"createdAt": `"createdAt"`,
		"id":        "id",
	}

	sortBy, ok := sortMap[sort.SortBy]
	if !ok {
		sortBy = `"createdAt"`
	}

	sortOrder := "DESC"
	if strings.ToUpper(sort.SortOrder) == "ASC" {
		sortOrder = "ASC"
	}

	queryBuilder.WriteString(fmt.Sprintf(" ORDER BY %s %s", sortBy, sortOrder))

	queryBuilder.WriteString(fmt.Sprintf(" LIMIT $%d OFFSET $%d", argCounter, argCounter+1))
	args = append(args, pagination.Limit, (pagination.Page-1)*pagination.Limit)

	rows, err := service.Pool.Query(service.Context, queryBuilder.String(), args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cars := make([]models.Car, 0)
	for rows.Next() {
		var car models.Car
		if err := rows.Scan(&car.ID, &car.Name, &car.Image, &car.Description, &car.CreatedAt); err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}

	return cars, nil
}

func (service *Service) GetCarByID(id string) {}

func (service *Service) CreateCar(name, image, description string) {}

func (service *Service) UpdateCar(id, name, image, description string) {}

func (service *Service) DeleteCar(id string) {}
