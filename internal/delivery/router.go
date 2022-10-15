package delivery

import (
	"awesomeProject/internal/delivery/product"
	product3 "awesomeProject/internal/repository/product"
	product2 "awesomeProject/internal/usecase/product"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"os"
)

func Routes(r *gin.Engine) {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:root@localhost:5432/ecommerce")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	repo := product3.NewRepo(conn)
	usecase := product2.NewUseCase(repo)

	r.GET("/products", product.GetProducts(usecase))
	r.GET("/products/:id", product.GetProduct(usecase))

}
