package routes

import (
	"app/app/controller"

	"github.com/gin-gonic/gin"
)

func User_role(router *gin.RouterGroup) {
	// Get the *bun.DB instance from config
	ctl := controller.New() // Pass the *bun.DB to the controller
	//md := middleware.AuthMiddleware()
	// log := middleware.NewLogResponse()
	user_role := router.Group("")
	{
		user_role.POST("/create", ctl.UserRoleCtl.Create)
	}
}
