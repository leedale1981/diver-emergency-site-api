package controller

var (
	sitesController SitesController
	usersController UsersController
)

func Startup() {
	sitesController.registerRoutes()
	usersController.registerRoutes()
}
