package controllers

import (
	"net/http"
	"saya/config"
	"saya/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all categories
func GetCategoriesController(c echo.Context) error {
	var categories []models.Category

	if err := config.DB.Find(&categories).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "success get all categories",
		"categories": categories,
	})
}

// get category by id
func GetCategoryController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var category models.Category
	if err = config.DB.Where("id = ?", id).First(&category).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Success get category by id",
		"category": category,
	})
}

// create new category
func CreateCategoryController(c echo.Context) error {
	category := models.Category{}
	c.Bind(&category)

	if err := config.DB.Save(&category).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success create new cetegory",
		"category": category,
	})
}

// delete Category by id
func DeleteCategoryController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var category models.Category
	if err := config.DB.First(&category, "id = ? ", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Category not found",
		})
	}

	if err := config.DB.Delete(&category).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to delete category",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully deleted category data",
		"data":    category,
	})
}

// update category by id
func UpdateCategoryController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var category models.Category
	if err := config.DB.Where("id = ?", id).First(&category).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{
			"message": "category not found",
		})
	}
	c.Bind(&category)
	if err := config.DB.Model(&category).Updates(category).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success update category data",
		"category": category,
	})
}
