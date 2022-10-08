package product

import (
	"awesomeProject/internal/entity"
	"context"
	"fmt"
)

type Repo interface {
	GetAll(ctx context.Context, page int, limit int) ([]entity.Product, error)
}

type repo struct {
	//db *sql.DB
}

func NewRepo() repo {
	return repo{}
}

func (r repo) GetAll(ctx context.Context, page int, limit int) ([]entity.Product, error) {
	return []entity.Product{{
		Name:        "pen",
		Price:       10000,
		Quantity:    100,
		ImageUrl:    "https://www.example.com",
		Description: "Bolpoint",
		Sku:         "1430011",
	},
		{
			Name:        "tv",
			Price:       5_000_000,
			Quantity:    100,
			ImageUrl:    "https://www.example.com",
			Description: "Televisi 43 inch",
			Sku:         "143012",
		}}, fmt.Errorf("error")
}
