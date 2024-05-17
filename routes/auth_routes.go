package routes

import (
	"elemento-api/app/controllers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteAuth(apiv1 *echo.Group, db *gorm.DB) {
	authController := controllers.NewAuthController(db)

	apiv1.POST("/login", authController.Login)
	apiv1.POST("/register", authController.Register)
	apiv1.GET("/user", authController.GetUser)
	apiv1.POST("/refresh", authController.Refresh)
}
