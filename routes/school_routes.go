package routes

import (
	"elemento-api/app/controllers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteSchool(apiv1 *echo.Group, db *gorm.DB) {
	schoolController := controllers.NewSchoolController(db)

	apiv1.POST("/school", schoolController.CreateNewSchool)
	apiv1.GET("/school", schoolController.GetAllSchool)
	apiv1.GET("/school/:id", schoolController.GetSchoolById)
	apiv1.POST("/school/class/:id", schoolController.IntegrateClassToSchool)
	apiv1.GET("/school/class/:id", schoolController.GetClassBySchoolId)
}
