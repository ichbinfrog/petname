package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// Instance is a structure that stores the router
type Instance struct {
	Router *mux.Router
	API    map[string]API
}

// SetupRouter generates the initial router configurations
// for the petname API server
func (i *Instance) SetupRouter() {
	i.Router = mux.NewRouter()
	i.API = make(map[string]API)
	i.SetupAPI("default", false, "{{ .Adverb }}{{ .Adjective }}{{ .Name }}", "~")

	i.Router.
		Name("GetHealth").
		Path("/health").
		HandlerFunc(HealthGet)

	i.Router.
		Name("GetPetname").
		Path("/get/{api}").
		HandlerFunc(i.GetPetname)

	i.Router.
		Name("AddSeed").
		Path("/api/{api}/add").
		HandlerFunc(i.AddSeed)

	i.Router.
		Name("ReloadAPI").
		Path("/api/{api}/reload").
		HandlerFunc(i.ReloadAPI)

	i.Router.
		Name("RemoveSeed").
		Path("/api/{api}/remove").
		HandlerFunc(i.RemoveSeed)

	i.Router.
		Name("GetInfoAPI").
		Path("/api/{api}").
		HandlerFunc(i.GetInfoAPI)

	i.Router.
		Name("AddAPI").
		Path("/api").
		HandlerFunc(i.AddAPI)
}

// Start loads router configuration and starts the http listening
func (i *Instance) Start(port int) {
	i.SetupRouter()
	http.ListenAndServe(fmt.Sprintf(":%d", port), i.Router)
}
