package routes

import (
	"kstyle-be-techtest/handlers"

	"github.com/labstack/echo/v4"
)

func MemberRoutes() *echo.Echo {
	e := echo.New()

	e.POST("/members", handlers.CreateMember)
	e.PUT("/members/:id", handlers.UpdateMember)
	e.DELETE("/members/:id", handlers.DeleteMember)
	e.GET("/members", handlers.GetAllMember)

	e.POST("/products", handlers.CreateProduct)
	e.GET("/products", handlers.GetAllProduct)

	return e
}
