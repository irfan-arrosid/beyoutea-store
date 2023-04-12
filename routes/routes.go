package routes

import (
	"kstyle-be-techtest/handlers"

	"github.com/labstack/echo/v4"
)

func MemberRoutes() *echo.Echo {
	e := echo.New()
	e.POST("/member", handlers.CreateMember)
	e.PUT("/member/:id", handlers.UpdateMember)
	e.GET("/member/:id", handlers.GetMemberByID)

	return e
}
