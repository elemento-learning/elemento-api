package routes

import (
	"elemento-api/app/controllers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// RouteMagicCard is a function to define the routes for magic card
func RouteModule(apiv1 *echo.Group, db *gorm.DB) {
	moduleController := controllers.NewModulController(db)

	apiv1.POST("/module", moduleController.CreateNewModul)
	apiv1.GET("/module/:id", moduleController.GetModulById)
	apiv1.POST("/module/bab", moduleController.CreateBabAndIntegrateToModul)
}
