package routes

import (
	"elemento-api/app/controllers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteMagicCard(apiv1 *echo.Group, db *gorm.DB) {
	magicCardController := controllers.NewMagicCardController(db)

	apiv1.POST("/magic-card", magicCardController.CreateMagicCard)
	apiv1.GET("/magic-card/:id", magicCardController.GetMagicCardById)
}
