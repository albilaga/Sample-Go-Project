package main

import (
	"awesomeProject/internal/delivery"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	delivery.Routes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
