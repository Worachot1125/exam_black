package routes

import (
	"app/app/controller"

	"github.com/gin-gonic/gin"
)

func Role(router *gin.RouterGroup) {
	// Get the *bun.DB instance from config
	ctl := controller.New() // Pass the *bun.DB to the controller
	//md := middleware.AuthMiddleware()
	// log := middleware.NewLogResponse()
	role := router.Group("")
	{
		role.POST("/create", ctl.RoleCtl.Create)
		role.GET("/list", ctl.RoleCtl.List)
	}
}