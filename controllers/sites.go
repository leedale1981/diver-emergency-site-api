package controller

import "net/http"

type SitesController struct {
}

func (sitesController SitesController) registerRoutes() {
	http.HandleFunc("/sites", sitesController.handleSites)
}

func (sitesController SitesController) handleSites(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello sites!"))
}
