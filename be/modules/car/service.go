package car

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CarService struct {
	Pool    *pgxpool.Pool
	Context context.Context
}

// todo -> pagination, sort interfaces
func (c *CarService) GetAllCars() {}

func (c *CarService) GetCarByID(id string) {}

func (c *CarService) CreateCar(name, image, description string) {}

func (c *CarService) UpdateCar(id, name, image, description string) {}

func (c *CarService) DeleteCar(id string) {}
