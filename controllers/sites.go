package controller

import "net/http"

type SitesController struct {
}

func (sitesController SitesController) registerRoutes() {
	http.HandleFunc("/sites", sitesController.handleSites)
	http.HandleFunc("/sites/nearest/", sitesController.handleNearestSite)
}

func (sitesController SitesController) handleSites(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello sites!"))
}

func (sitesController SitesController) handleNearestSite(w http.ResponseWriter, r *http.Request) {

}
