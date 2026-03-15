package lap

import (
	"context"
	"fmt"
	"math"
	"performance_tracker_v2_be/core/handler"
	coreService "performance_tracker_v2_be/core/service"
	"performance_tracker_v2_be/db/main-db/models"
	"performance_tracker_v2_be/helpers"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	Pool *pgxpool.Pool
}

type filterRule struct {
	Expression string
	Operator   string
	IsFuzzy    bool
}

func (s *Service) GetFilters() map[string]filterRule {
	return map[string]filterRule{
		"date":  {Expression: "TO_CHAR(l.created_at, 'YYYY-MM-DD')", Operator: "ILIKE", IsFuzzy: true},
		"game":  {Expression: "g.name", Operator: "ILIKE", IsFuzzy: true},
		"car":   {Expression: "c.name", Operator: "ILIKE", IsFuzzy: true},
		"track": {Expression: "t.name", Operator: "ILIKE", IsFuzzy: true},
		"time":  {Expression: "CAST(l.time AS TEXT)", Operator: "ILIKE", IsFuzzy: true},
		"clear": {Expression: "CAST(l.is_clear AS TEXT)", Operator: "=", IsFuzzy: false},
	}
}

func (s *Service) GetAllLaps(
	ctx context.Context,
	pagination *handler.Pagination,
	sort *handler.Sort,
	filters *GetLapsFilter,
) ([]GetListLap, error) {
	whereClause, args := s.buildWhereClause(filters)

	sortColumn := s.getSortColumn(sort.SortBy)
	sortOrder := "DESC"
	if strings.ToUpper(sort.SortOrder) == "ASC" {
		sortOrder = "ASC"
	}

	query := `
		SELECT l.id, l.created_at, g.name, c.name, t.name, l.time, l.is_clear
		FROM laps l
		INNER JOIN games g ON g.id = l.game_id
		INNER JOIN cars c ON c.id = l.car_id
		INNER JOIN tracks t ON t.id = l.track_id
	` + whereClause + fmt.Sprintf(`
		ORDER BY %s %s
		LIMIT $%d OFFSET $%d
	`, sortColumn, sortOrder, len(args)+1, len(args)+2)

	args = append(args, pagination.Limit, (pagination.Page-1)*pagination.Limit)

	rows, err := s.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]GetListLap, 0)
	for rows.Next() {
		var lap GetListLap
		var createdAt time.Time
		var lapTime float64

		if err := rows.Scan(&lap.ID, &createdAt, &lap.Game, &lap.Car, &lap.Track, &lapTime, &lap.Clear); err != nil {
			return nil, err
		}

		lap.Date = createdAt.Format("2006-01-02")
		lap.Time = formatLapTime(lapTime)
		result = append(result, lap)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Service) GetTotalLapsCount(ctx context.Context, filters *GetLapsFilter) (int64, error) {
	whereClause, args := s.buildWhereClause(filters)

	query := `
		SELECT COUNT(*)
		FROM laps l
		INNER JOIN games g ON g.id = l.game_id
		INNER JOIN cars c ON c.id = l.car_id
		INNER JOIN tracks t ON t.id = l.track_id
	` + whereClause

	var count int64
	if err := s.Pool.QueryRow(ctx, query, args...).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func (s *Service) GetByID(ctx context.Context, id int64) (*models.Lap, error) {
	row := s.Pool.QueryRow(ctx, `
		SELECT id, car_id, track_id, game_id, wheel_id, cockpit_id, pedals_id, time, is_clear, has_significant_errors, created_at
		FROM laps
		WHERE id = $1
	`, id)

	var lap models.Lap
	if err := row.Scan(
		&lap.ID,
		&lap.CarID,
		&lap.TrackID,
		&lap.GameID,
		&lap.WheelID,
		&lap.CockpitID,
		&lap.PedalsID,
		&lap.Time,
		&lap.IsClear,
		&lap.HasSignificantErrors,
		&lap.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &lap, nil
}

func (s *Service) Create(ctx context.Context, payload *CreateRequest) (int64, error) {
	var id int64

	err := s.Pool.QueryRow(ctx, `
		INSERT INTO laps (
			car_id,
			track_id,
			game_id,
			wheel_id,
			cockpit_id,
			pedals_id,
			time,
			is_clear,
			has_significant_errors
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id
	`,
		payload.CarID,
		payload.TrackID,
		payload.GameID,
		payload.WheelID,
		payload.CockpitID,
		payload.PedalsID,
		payload.Time,
		payload.IsClear,
		payload.HasSignificantErrors,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Service) UpdateByID(ctx context.Context, id int64, payload *UpdateRequestParsed) error {
	updates := helpers.StructToMap[handler.IOptionalField](payload)

	return coreService.UpdateEntity(handler.UpdateEntityByIdPayload{
		ID:       id,
		Executor: s.Pool,
		Context:  ctx,
		Updates:  updates,
		Table:    "laps",
	})
}

func (s *Service) DeleteByID(ctx context.Context, id int64) error {
	_, err := s.Pool.Exec(ctx, `
		DELETE FROM laps
		WHERE id = $1
	`, id)

	return err
}

func (s *Service) buildWhereClause(filters *GetLapsFilter) (string, []interface{}) {
	filterValues := helpers.StructToMap[[]string](filters)
	filterRules := s.GetFilters()
	keys := []string{"date", "game", "car", "track", "time", "clear"}

	clauses := make([]string, 0)
	args := make([]interface{}, 0)
	argCounter := 1

	for _, key := range keys {
		values := filterValues[key]
		rule, ok := filterRules[key]
		if !ok || len(values) == 0 {
			continue
		}

		orConditions := make([]string, 0)
		for _, value := range values {
			if value == "" {
				continue
			}

			argValue := value
			if rule.IsFuzzy {
				argValue = "%" + value + "%"
			}

			orConditions = append(orConditions, fmt.Sprintf("%s %s $%d", rule.Expression, rule.Operator, argCounter))
			args = append(args, argValue)
			argCounter++
		}

		if len(orConditions) > 0 {
			clauses = append(clauses, "("+strings.Join(orConditions, " OR ")+")")
		}
	}

	if len(clauses) == 0 {
		return "", args
	}

	return " WHERE " + strings.Join(clauses, " AND "), args
}

func (s *Service) getSortColumn(sortBy string) string {
	sortColumns := map[string]string{
		"id":    "l.id",
		"date":  "l.created_at",
		"game":  "g.name",
		"car":   "c.name",
		"track": "t.name",
		"time":  "l.time",
		"clear": "l.is_clear",
	}

	if sortColumn, ok := sortColumns[sortBy]; ok {
		return sortColumn
	}

	return "l.id"
}

func formatLapTime(seconds float64) string {
	totalMilliseconds := int64(math.Round(seconds * 1000))
	if totalMilliseconds < 0 {
		totalMilliseconds = 0
	}

	minutes := totalMilliseconds / 60000
	remainingMilliseconds := totalMilliseconds % 60000
	secs := remainingMilliseconds / 1000
	milliseconds := remainingMilliseconds % 1000

	return fmt.Sprintf("%d:%02d.%03d", minutes, secs, milliseconds)
}
