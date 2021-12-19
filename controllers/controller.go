package controller

var (
	sitesController    SitesController
	usersController    UsersController
	chambersController ChambersController
)

func Startup() {
	sitesController.registerRoutes()
	usersController.registerRoutes()
	chambersController.registerRoutes()
}
