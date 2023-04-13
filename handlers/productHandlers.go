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
		return c.String(http.StatusBadRequest, "Failed to create product")
	}

	return c.JSON(http.StatusOK, product)
}

func GetAllProduct(c echo.Context) error {
	var product []entities.Product

	if err := database.DB.Model(&entities.Product{}).Preload("Reviews").Find(&product).Error; err != nil {
		return c.String(http.StatusBadRequest, "Product is empty")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"products": product,
	})
}

//Select product by ID_PRODUCT dan data review dari product tersebut. Datareview yang harus ditampilkan USERNAME, GENDER, SKINTYPE,SKINCOLOR, DESC_REVIEW, JUMLAH_LIKE_REVIEW.

func GetProductById(c echo.Context) error {
	productId := c.Param("id")
	var product entities.Product
	var review []entities.Review_product
	var member []entities.Member

	if err := database.DB.Model(&entities.Product{}).Preload("Reviews").First(&product, productId).Error; err != nil {
		return c.String(http.StatusBadRequest, "Product not found")
	}

	if err := database.DB.Model(&entities.Review_product{}).Preload("Members").Find(&review).Error; err != nil {
		return c.String(http.StatusBadRequest, "Review not found")
	}

	if err := database.DB.Model(&entities.Member{}).Preload("Reviews").Find(&member).Error; err != nil {
		return c.String(http.StatusBadRequest, "Member not found")
	}

	return c.JSON(http.StatusOK, product)

	// Member field is still NULL :(
}

func CreateProductReview(c echo.Context) error {
	productReview := new(entities.Review_product)
	if err := c.Bind(&productReview); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	if err := database.DB.Create(&productReview).Error; err != nil {
		return c.String(http.StatusBadRequest, "Failed to create product review")
	}

	return c.JSON(http.StatusOK, productReview)
}
