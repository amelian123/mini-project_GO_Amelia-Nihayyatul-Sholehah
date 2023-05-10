package middlewares

import (
	"net/http"
	"saya/constants"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(userID int, name, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userID"] = userID
	claims["name"] = name
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}

func ExtractTokenUserId(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userID"].(int)
		return userId
	}
	return 0
}

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid or missing jwt token")
		}
		claims, ok := user.Claims.(jwt.MapClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid jwt claims")
		}
		if role, ok := claims["role"].(string); !ok || role != "Admin" {
			return echo.NewHTTPError(http.StatusUnauthorized, "user is not an admin")
		}
		return next(c)
	}
}
