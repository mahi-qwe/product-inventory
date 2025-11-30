package routes

import (
	"github.com/gin-gonic/gin"

	"product-inventory/internal/handlers"
)

func Register(r *gin.Engine) {
	r.POST("/products", handlers.CreateProduct)
	r.GET("/products", handlers.ListProducts)

	r.POST("/stock/add", handlers.AddStock)
	r.POST("/stock/remove", handlers.RemoveStock)

	r.GET("/stock/report", handlers.StockReport)
}
