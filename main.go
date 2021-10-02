package main

import (
	"net/http"

	controller "github.com/leedale1981/diver-emergency-site-api/controllers"
)

func main() {
	controller.Startup()
	http.ListenAndServe(":8000", nil)
}
