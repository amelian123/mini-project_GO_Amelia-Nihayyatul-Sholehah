package routes

import (
	"saya/constants"
	"saya/controllers"
	m "saya/middlewares"

	mid "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	m.LogMiddleware(e)

	eUser := e.Group("/users")
	eUser.POST("/register", controllers.CreateUserController)
	eUser.POST("/login", controllers.LoginUserController)
	// 	Authenticated with JWT
	eUserJwt := eUser.Group("")
	eUserJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eUserJwt.GET("", controllers.GetUsersController)
	eUserJwt.GET("/:id", controllers.GetUserController)
	eUserJwt.PUT("/:id", controllers.UpdateUserController)
	eUserJwt.DELETE("/:id", controllers.DeleteUserController)

	eAdmin := e.Group("/admins")
	eAdmin.POST("/register", controllers.CreateAdminController)
	eAdmin.POST("/login", controllers.LoginAdminController)
	// 	Authenticated with JWT
	eAdminJwt := eAdmin.Group("")
	eAdminJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eAdminJwt.GET("", controllers.GetAdminsController)
	eAdminJwt.GET("/:id", controllers.GetAdminController)
	eAdminJwt.PUT("/:id", controllers.UpdateAdminController)
	eAdminJwt.DELETE("/:id", controllers.DeleteAdminController)

	//categories routes
	eCategory := e.Group("/categories")
	eCategory.GET("", controllers.GetCategoriesController)
	eCategory.GET("/:id", controllers.GetCategoryController)

	//categories routes admin
	eCategoryJwt := eCategory.Group("/admin")
	eCategoryJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eCategoryJwt.POST("", controllers.CreateCategoryController)
	eCategoryJwt.PUT("/:id", controllers.UpdateCategoryController)
	eCategoryJwt.DELETE("/:id", controllers.DeleteCategoryController)

	//event routes
	eEvents := e.Group("/events")
	eEvents.GET("", controllers.GetEventsController)
	eEvents.GET("/:id", controllers.GetEventController)

	//products routes admin
	eEventsJwt := eEvents.Group("/admin")
	eEventsJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eEventsJwt.POST("", controllers.CreateEventController)
	eEventsJwt.PUT("/:id", controllers.UpdateEventController)
	eEventsJwt.DELETE("/:id", controllers.DeleteEventController)

	//Orders routes admin
	eOrders := e.Group("/orders")
	eOrders.GET("", controllers.GetOrdersController)
	eOrders.GET("/:id", controllers.GetOrderController)

	//Orders routes user
	eOrdersJwt := eOrders.Group("/user")
	eOrdersJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eOrdersJwt.POST("", controllers.CreateOrderController)
	eEventsJwt.PUT("/:id", controllers.UpdateOrderController)
	eOrdersJwt.DELETE("/:id", controllers.DeleteOrderController)

	return e
}
