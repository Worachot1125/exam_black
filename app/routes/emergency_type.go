package routes

import (
	"app/app/controller"

	"github.com/gin-gonic/gin"
)

func Emergency_type(router *gin.RouterGroup) {
	// Get the *bun.DB instance from config
	ctl := controller.New() // Pass the *bun.DB to the controller
	// md := middleware.AuthMiddleware()
	// log := middleware.NewLogResponse()
	emergency_type := router.Group("")
	{
		emergency_type.POST("/create", ctl.Emergency_TypeCtl.Create)
		emergency_type.GET("/list", ctl.Emergency_TypeCtl.List)
	}
}
