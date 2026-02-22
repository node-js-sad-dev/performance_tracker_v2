package game

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

func (service *Service) GetList(
	ctx context.Context,
	pagination *core.Pagination,
	sort *core.Sort,
	filters *GetFilters,
) ([]models.Game, error) {
	rows, err := core.GetEntityList(core.GetEntityListPayload{
		Pool:         service.Pool,
		Context:      ctx,
		TableName:    "games",
		Pagination:   pagination,
		Sort:         sort,
		Filters:      helpers.StructToMap[[]string](filters),
		FilterRules:  service.GetFilters(),
		SelectFields: []string{"id", "name", "image", "created_at"},
	})

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]models.Game, 0)
	for rows.Next() {
		var game models.Game
		if err := rows.Scan(&game.ID, &game.Name, &game.Image, &game.CreatedAt); err != nil {
			return nil, err
		}
		result = append(result, game)
	}

	return result, nil
}

func (service *Service) GetTotalCount(ctx context.Context, filters *GetFilters) (int64, error) {
	return core.GetEntityCount(core.GetEntityCountPayload{
		Pool:        service.Pool,
		Context:     ctx,
		TableName:   "games",
		Filters:     helpers.StructToMap[[]string](filters),
		FilterRules: service.GetFilters(),
	})
}

func (service *Service) GetByID(ctx context.Context, id int64) (*models.Game, error) {
	game := service.Pool.QueryRow(ctx, `
		SELECT id, name, image, created_at
		FROM games
		WHERE id = $1
	`, id)

	var result models.Game
	if err := game.Scan(&result.ID, &result.Name, &result.Image, &result.CreatedAt); err != nil {
		return nil, err
	}

	return &result, nil
}

func (service *Service) Create(ctx context.Context, payload *CreateRequest) (int64, error) {
	var id int64
	err := service.Pool.QueryRow(ctx, `
		INSERT INTO games (name, image)
		VALUES ($1, $2)
		RETURNING id
	`, payload.Name, payload.Image).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (service *Service) UpdateByID(ctx context.Context, id int64, payload *UpdateRequestParsed) error {
	updates := helpers.StructToMap[core.IOptionalField](payload)

	return core.UpdateEntity(core.UpdateEntityByIdPayload{
		ID:       id,
		Executor: service.Pool,
		Context:  ctx,
		Updates:  updates,
		Table:    "cars",
	})
}

func (service *Service) DeleteByID(ctx context.Context, id int64) error {
	_, err := service.Pool.Exec(ctx, `
		DELETE FROM games
		WHERE id = $1
	`, id)

	return err
}
