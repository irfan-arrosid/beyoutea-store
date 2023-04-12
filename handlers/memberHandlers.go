package handlers

import (
	"kstyle-be-techtest/database"
	"kstyle-be-techtest/models/entities"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateMember(c echo.Context) error {
	member := new(entities.Member)
	if err := c.Bind(&member); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	if err := database.DB.Create(&member).Error; err != nil {
		return c.String(http.StatusBadRequest, "Failed to create member")
	}

	return c.JSON(http.StatusOK, member)
}

func UpdateMember(c echo.Context) error {

	memberId := c.Param("id")
	var member entities.Member

	if err := database.DB.First(&member, memberId).Error; err != nil {
		return c.String(http.StatusBadRequest, "Member not found")
	}

	// newMember := new(entities.Member)
	if err := c.Bind(&member); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	if err := database.DB.Save(&member).Error; err != nil {
		return c.String(http.StatusBadRequest, "Failed to update member")
	}

	return c.JSON(http.StatusOK, member)
}

func GetMemberByID(c echo.Context) error {
	var member entities.Member
	memberId := c.Param("id")

	if err := database.DB.First(&member, memberId).Error; err != nil {
		return c.String(http.StatusBadRequest, "Member not found")
	}

	return c.JSON(http.StatusOK, member)
}
