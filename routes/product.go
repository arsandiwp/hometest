package routes

import (
	"nutech/handlers"
	"nutech/pkg/middleware"
	"nutech/pkg/mysql"
	"nutech/repositories"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Group) {
	productRepository := repositories.RepositoryProduct(mysql.DB)
	h := handlers.HandlerProduct(productRepository)

	e.POST("/product", middleware.UploadFile(h.CreateProduct))
	e.GET("/products", h.FindProduct)
	e.GET("/product/:id", h.GetProduct)
	e.PATCH("/product/:id", middleware.UploadFile(h.UpdateProduct))
	e.DELETE("/product/:id", h.DeleteProduct)
}
