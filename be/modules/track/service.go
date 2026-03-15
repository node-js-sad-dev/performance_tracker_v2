package track

import (
	"context"
	"errors"
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
) ([]models.Track, error) {
	rows, err := service.GetEntityList(handler.GetEntityListPayload{
		Pool:         s.Pool,
		Context:      ctx,
		TableName:    "tracks",
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

	result := make([]models.Track, 0)
	for rows.Next() {
		var track models.Track
		if err := rows.Scan(&track.ID, &track.Name, &track.Image, &track.Description, &track.CreatedAt); err != nil {
			return nil, err
		}
		result = append(result, track)
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
		TableName:   "tracks",
		Filters:     helpers.StructToMap[[]string](filters),
		FilterRules: s.GetFilters(),
	})
}

func (s *Service) GetByID(ctx context.Context, id int64) (*models.Track, error) {
	row := s.Pool.QueryRow(ctx, `
		SELECT id, name, image, description, created_at
		FROM tracks
		WHERE id = $1
	`, id)

	var result models.Track
	if err := row.Scan(&result.ID, &result.Name, &result.Image, &result.Description, &result.CreatedAt); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *Service) Create(ctx context.Context, payload *CreateRequestParsed) (int64, error) {
	var id int64
	if err := s.Pool.QueryRow(ctx, `
		INSERT INTO tracks (name, image, description)
		VALUES ($1, $2, $3)
		RETURNING id
	`, payload.Name, payload.Image, payload.Description).Scan(&id); err != nil {
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
		Table:    "tracks",
	})
}

func (s *Service) DeleteByID(ctx context.Context, id int64) error {
	_, err := s.Pool.Exec(ctx, `
		DELETE FROM tracks
		WHERE id = $1
	`, id)

	return err
}

func (s *Service) GetByName(ctx context.Context, name string, idToSkip *int64) (*models.Track, error) {
	query := `
		SELECT id, name, image, description, created_at
		FROM tracks
		WHERE name = $1
	`

	args := []any{name}
	if idToSkip != nil {
		query += " AND id != $2"
		args = append(args, *idToSkip)
	}

	row := s.Pool.QueryRow(ctx, query, args...)

	var result models.Track
	if err := row.Scan(&result.ID, &result.Name, &result.Image, &result.Description, &result.CreatedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &result, nil
}
