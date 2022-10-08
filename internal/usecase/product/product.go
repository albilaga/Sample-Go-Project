package product

import (
	"awesomeProject/internal/entity"
	"awesomeProject/internal/repository/product"
	"context"
)

type ProductUseCase interface {
	GetProducts(ctx context.Context, page int, limit int) ([]entity.Product, error)
}

type usecase struct {
	repo product.Repo
}

func NewUseCase(repo product.Repo) usecase {
	return usecase{repo: repo}
}

func (u usecase) GetProducts(ctx context.Context, page int, limit int) ([]entity.Product, error) {
	return u.repo.GetAll(ctx, page, limit)
}
