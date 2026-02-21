package lap

import (
	"context"
	"performance_tracker_v2_be/core"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	Pool *pgxpool.Pool
}

func (service *Service) GetAllLaps(
	ctx context.Context,
	pagination *core.Pagination,
	sort *core.Sort,
	filters *GetLapsFilter,
) ([]GetListLap, error) {
	// todo: get real values from db, mock for now

	return []GetListLap{
		{
			ID:    1,
			Date:  "2024-01-01",
			Game:  "Game 1",
			Car:   "Car 1",
			Track: "Track 1",
			Time:  "1:00.000",
			Clear: true,
		},
		{
			ID:    2,
			Date:  "2024-01-02",
			Game:  "Game 2",
			Car:   "Car 2",
			Track: "Track 2",
			Time:  "2:00.000",
			Clear: false,
		},
	}, nil
}

func (service *Service) GetTotalLapsCount(ctx context.Context, filters *GetLapsFilter) (int64, error) {
	// todo: get real count from db, mock for now

	return 2, nil
}
