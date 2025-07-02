package routes

import (
	"app/app/controller"
	"app/app/middleware"

	"github.com/gin-gonic/gin"
)

func Emergency_report(router *gin.RouterGroup) {
	// Get the *bun.DB instance from config
	ctl := controller.New() // Pass the *bun.DB to the controller
	md := middleware.AuthMiddleware()
	// log := middleware.NewLogResponse()
	emergency_report := router.Group("")
	{
		emergency_report.POST("/create", md, ctl.Emergency_reportCtl.Create)
		emergency_report.PATCH("/:id", md, ctl.Emergency_reportCtl.Update)
		emergency_report.GET("/list", ctl.Emergency_reportCtl.List)
		emergency_report.GET("/:id", ctl.Emergency_reportCtl.Get)
		emergency_report.DELETE("/:id", ctl.Emergency_reportCtl.Delete)
		emergency_report.GET("/user/:id", md, ctl.Emergency_reportCtl.GetByUserIDEmergency) // Get reports by user ID
	}
}
