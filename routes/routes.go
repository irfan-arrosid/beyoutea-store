package routes

import (
	"kstyle-be-techtest/handlers"

	"github.com/labstack/echo/v4"
)

func MemberRoutes() *echo.Echo {
	e := echo.New()

	e.POST("/member", handlers.CreateMember)
	e.PUT("/member/:id", handlers.UpdateMember)
	e.DELETE("/member/:id", handlers.DeleteMember)
	e.GET("/member", handlers.GetAllMember)

	return e
}
