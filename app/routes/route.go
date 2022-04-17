package routes

import (
	"rest-go-stock/controllers/product_types"
	"rest-go-stock/controllers/products"
	"rest-go-stock/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouteControllerList struct {
	JWTMiddleware  middleware.JWTConfig
	UserController users.UserController
	ProductController products.ProductController
	ProductTypeController product_types.ProductTypeController
}

func (ctrl *RouteControllerList) RouteRegister(e *echo.Echo) {
	jwt := middleware.JWTWithConfig(ctrl.JWTMiddleware)
	
	e.DELETE("/user/:id", ctrl.UserController.DeleteUserByID, jwt)
	e.POST("register", ctrl.UserController.RegisterUser)
	e.POST("login", ctrl.UserController.LoginUser)
	e.PUT("user/password/:id", ctrl.UserController.UpdatePasswordByID, jwt)
	e.POST("user/checking", ctrl.UserController.CheckingUser)
	e.GET("user/:id", ctrl.UserController.GetByID)
	e.POST("user", ctrl.UserController.GetByEmail)
	e.GET("users", ctrl.UserController.GetAllUsers)
	e.PUT("user/:id", ctrl.UserController.UpdateUserByID, jwt)

	e.POST("/product_type",  ctrl.ProductTypeController.AddProductType)
	// e.GET("/product_type/:id", ctrl.ProductTypeController.GetByID)

	e.POST("/product", ctrl.ProductController.Add)

}