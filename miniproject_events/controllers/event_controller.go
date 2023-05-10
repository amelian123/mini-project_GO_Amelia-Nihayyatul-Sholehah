package controllers

import (
	"net/http"
	"saya/config"
	"saya/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all event
func GetEventsController(c echo.Context) error {
	var events []models.Event

	if err := config.DB.Find(&events).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all event",
		"event":   events,
	})
}

// get event by id
func GetEventController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var event models.Event
	if err = config.DB.Where("id = ?", id).First(&event).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get event by id",
		"event":   event,
	})
}

// create new event
func CreateEventController(c echo.Context) error {
	event := models.Event{}
	c.Bind(&event)

	if err := config.DB.Save(&event).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new event",
		"event":   event,
	})
}

// delete Category by id
func DeleteEventController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var event models.Event
	if err := config.DB.First(&event, "id = ? ", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Event not found",
		})
	}

	if err := config.DB.Delete(&event).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to delete event",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully deleted event data",
		"event":   event,
	})
}

// update event by id
func UpdateEventController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var event models.Event
	if err := config.DB.Where("id = ?", id).First(&event).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{
			"message": "event not found",
		})
	}
	c.Bind(&event)
	if err := config.DB.Model(&event).Updates(event).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update event data",
		"event":   event,
	})
}
