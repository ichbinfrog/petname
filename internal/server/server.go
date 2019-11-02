package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Instance is a structure that stores the router
type Instance struct {
	Router *mux.Router
}

// SetupRouter uses mux to setup up routing paths for endpoints
// /health for health endpoints (liveness probe intended)
// /ready for readiness probes
// /reload for dynamic configuration reloads
func (i *Instance) SetupRouter() {
	i.Router = mux.NewRouter()
	i.Router.NotFoundHandler = http.HandlerFunc(handleError)
	i.Router.Path("/health").HandlerFunc(handleHealth)
	i.Router.Path("/ready").HandlerFunc(handleReadiness)
	i.Router.Path("/reload").HandlerFunc(handleReload)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusOK),
		http.StatusOK)
}

func handleReload(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusTeapot),
		http.StatusTeapot)
}

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusOK),
		http.StatusOK)
}

func handleError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusForbidden),
		http.StatusForbidden)
}

// Start loads router configuration and starts the http listening
func (i *Instance) Start(port int) {
	i.SetupRouter()
	fmt.Printf("Serving on port %d", port)
	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%d", port), i.Router))
}
