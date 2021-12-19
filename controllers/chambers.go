package controller

import (
	"encoding/json"
	"net/http"

	model "github.com/leedale1981/diver-emergency-site-api/models"
)

type ChambersController struct {
}

func (chambersController ChambersController) registerRoutes() {
	http.HandleFunc("/chambers", chambersController.handleSites)
	http.HandleFunc("/chambers/nearest/", chambersController.handleNearestSite)
}

func (chambersController ChambersController) handleSites(w http.ResponseWriter, r *http.Request) {
	chambers := model.GetMockChambers()
	encoder := json.NewEncoder(w)
	encoder.Encode(chambers)
}

func (chambersController ChambersController) handleNearestSite(w http.ResponseWriter, r *http.Request) {
	// latLongPattern, _ := regexp.Compile(`/chambers/nearest/(\d+)`)
	// matches := latLongPattern.FindStringSubmatch(r.URL.Path)

	// if len(matches) > 0 {
	// 	latlong, _ := strconv.Atoi(matches[1])
	// 	return chambers that match latLong value
	// }
}
