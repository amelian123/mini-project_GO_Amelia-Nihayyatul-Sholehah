package controllers

import (
	"net/http"
	"saya/config"
	"saya/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all categories
func GetOrdersController(c echo.Context) error {
	var orders []models.Order

	if err := config.DB.Find(&orders).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all order",
		"order":   orders,
	})
}

// get category by id
func GetOrderController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var order models.Order
	if err = config.DB.Where("id = ?", id).First(&order).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get order by id",
		"order":   order,
	})
}

// create new Order
func CreateOrderController(c echo.Context) error {
	order := models.Order{}
	c.Bind(&order)

	if err := config.DB.Create(&order).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new order",
		"order":   order,
	})
}

// delete Order by id
func DeleteOrderController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var order models.Order
	if err := config.DB.First(&order, "id = ? ", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Order not found",
		})
	}

	if err := config.DB.Delete(&order).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to delete order",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully deleted order data",
		"order":   order,
	})
}

// update category by id
func UpdateOrderController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var order models.Order
	if err := config.DB.Where("id = ?", id).First(&order).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{
			"message": "order not found",
		})
	}
	c.Bind(&order)
	if err := config.DB.Model(&order).Updates(order).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update order data",
		"order":   order,
	})
}
