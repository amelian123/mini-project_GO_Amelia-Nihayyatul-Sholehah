package controllers

import (
	"net/http"
	"reflect"
	"saya/config"
	database "saya/lib/database"
	"saya/middlewares"
	"saya/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all admins
func GetAdminsController(c echo.Context) error {
	var admins []models.Admin

	if err := config.DB.Find(&admins).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"admins":  admins,
	})
}

// get admin by id
func GetAdminController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var admin models.Admin
	if err = config.DB.Where("id = ?", id).First(&admin).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get admin by id",
		"admin":   admin,
	})
}

// create new admin
func CreateAdminController(c echo.Context) error {
	admin := models.Admin{}
	c.Bind(&admin)

	if err := config.DB.Save(&admin).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new admin",
		"admin":   admin,
	})
}

// delete admin by id
func DeleteAdminController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var admin models.Admin
	if err := config.DB.First(&admin, "id = ? ", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "admin not found",
		})
	}

	if err := config.DB.Delete(&admin).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to delete admin data",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully deleted admin data",
		"data":    admin,
	})
}

// update admin by id
func UpdateAdminController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var admin models.Admin
	if err := config.DB.Where("id = ?", id).First(&admin).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{
			"message": "admin not found",
		})
	}
	c.Bind(&admin)
	if err := config.DB.Model(&admin).Updates(admin).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update admin data",
		"admin":   admin,
	})
}

func LoginAdminController(c echo.Context) error {
	admin := models.Admin{}
	c.Bind(&admin)

	result, err := database.LoginAdminController(admin)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	reflectValue := reflect.ValueOf(result)
	adminID := reflectValue.FieldByName("ID").Interface().(uint)
	adminName := reflectValue.FieldByName("Name").Interface().(string)
	adminEmail := reflectValue.FieldByName("Email").Interface().(string)
	adminRole := reflectValue.FieldByName("Role").Interface().(string)

	token, err := middlewares.CreateToken(int(adminID), adminName, adminRole)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	adminResponse := models.AdminResponse{ID: adminID, Name: adminName, Email: adminEmail, Token: token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login admin",
		"user":    adminResponse,
	})
}
