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
	apiv1.GET("/magic-card", magicCardController.GetAllMagicCard)
	apiv1.PUT("/magic-card/:id", magicCardController.UpdateMagicCard)
	apiv1.DELETE("/magic-card/:id", magicCardController.DeleteMagicCard)
	apiv1.POST("/magic-card/senyawa/:id", magicCardController.CreateSenyawaAndIntegrateToMagicCard)
}
