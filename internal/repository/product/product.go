package product

import (
	"awesomeProject/internal/entity"
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

type Repo interface {
	GetAll(ctx context.Context, page int, limit int) ([]entity.Product, error)
	GetByID(ctx context.Context, id int64) (entity.Product, error)
}

type repo struct {
	db *pgx.Conn
}

func NewRepo(db *pgx.Conn) repo {
	return repo{db: db}
}

func (r repo) GetAll(ctx context.Context, page int, limit int) ([]entity.Product, error) {
	result := make([]entity.Product, 0)

	const query = "SELECT * FROM products"
	err := pgxscan.Select(ctx, r.db, &result, query)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r repo) GetByID(ctx context.Context, id int64) (result entity.Product, err error) {
	const query = "SELECT * FROM products WHERE id=$1 AND name=$2"
	err = pgxscan.Get(ctx, r.db, &result, query, id, "product A")
	return result, err
}
