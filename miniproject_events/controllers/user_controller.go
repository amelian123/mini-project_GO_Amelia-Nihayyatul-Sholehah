package controllers

import (
	"net/http"
	"reflect"
	"strconv"

	database "saya/lib/database"
	"saya/middlewares"
	"saya/models"

	"github.com/labstack/echo/v4"
)

// get all users
func GetUsersController(c echo.Context) error {
	users, err := database.GetUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := database.GetUser(uint(userID))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get user",
		"users":    user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	result, err := database.CreateUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create new user",
		"user":    result,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := models.User{}
	c.Bind(&user)

	result, err := database.UpdateUser(uint(userID), user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update user",
		"users":    result,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result, err := database.DeleteUser(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete user",
		"id":       result,
	})
}

func LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	result, err := database.LoginUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	reflectValue := reflect.ValueOf(result)
	userID := reflectValue.FieldByName("ID").Interface().(uint)
	userName := reflectValue.FieldByName("Name").Interface().(string)
	userEmail := reflectValue.FieldByName("Email").Interface().(string)
	userRole := reflectValue.FieldByName("Role").Interface().(string)

	token, err := middlewares.CreateToken(int(userID), userName, userRole)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userResponse := models.UserResponse{ID: userID, Name: userName, Email: userEmail, Token: token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login user",
		"user":    userResponse,
	})
}
