package handlers

import (
	"kstyle-be-techtest/database"
	"kstyle-be-techtest/models/entities"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateProduct(c echo.Context) error {
	product := new(entities.Product)
	if err := c.Bind(&product); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	if err := database.DB.Create(&product).Error; err != nil {
		return c.String(http.StatusBadRequest, "Failed to create member")
	}

	return c.JSON(http.StatusOK, product)
}

func GetAllProduct(c echo.Context) error {
	var product []entities.Product

	if err := database.DB.Find(&product).Error; err != nil {
		return c.String(http.StatusBadRequest, "Member not found")
	}

	return c.JSON(http.StatusOK, product)
}
