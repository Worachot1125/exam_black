package controller

import (
	"app/app/controller/emergency_report"
	"app/app/controller/emergency_type"
	"app/app/controller/login"
	"app/app/controller/logout"
	"app/app/controller/role"
	"app/app/controller/user"
	"app/app/controller/user_role"
	"app/config"
)

type Controller struct {
	LoginCtl            *login.Controller  // Assuming LoginController is in the student package
	LogoutCtl           *logout.Controller // Uncomment if you have a logout controller
	UserCtl             *user.Controller
	UserRoleCtl         *user_role.Controller
	Emergency_reportCtl *emergency_report.Controller
	Emergency_TypeCtl   *emergency_type.Controller
	RoleCtl             *role.Controller

	// Other controllers...
}

func New() *Controller {
	db := config.GetDB()
	return &Controller{

		LoginCtl:            login.NewController(db), // Assuming LoginController is in the student package
		LogoutCtl:           logout.NewController(db),
		UserCtl:             user.NewController(db),
		Emergency_reportCtl: emergency_report.NewController(db),
		Emergency_TypeCtl:   emergency_type.NewController(db),
		RoleCtl:             role.NewController(db),
		UserRoleCtl:         user_role.NewController(db), // Uncomment if you have a logout controller
		// Other controllers...
	}
}
