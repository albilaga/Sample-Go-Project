package delivery

import (
	"awesomeProject/internal/delivery/product"
	product3 "awesomeProject/internal/repository/product"
	product2 "awesomeProject/internal/usecase/product"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	repo := product3.NewRepo()
	usecase := product2.NewUseCase(repo)

	r.GET("/products", product.GetProducts(usecase))
}
