package controller

import "net/http"

type UsersController struct {
}

func (usersController UsersController) registerRoutes() {
	http.HandleFunc("/users", usersController.handleUsers)
}

func (usersController UsersController) handleUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello users!"))
}
