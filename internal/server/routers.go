package server

import (
	"github.com/gorilla/mux"
)

// Instance is a structure that stores the router
type Instance struct {
	Router *mux.Router
	API    []API
}

func (i *Instance) SetupRouter() {
	i.Router = mux.NewRouter()
	i.SetupAPI("default", false, "{{ .Adverb }}{{ .Adjective }}{{ .Name }}", "~")

	i.Router.
		Name("GetHealth").
		Path("/v1/health").
		HandlerFunc(HealthGet)

	i.Router.
		Name("GetPetname").
		Path("/get/{api}").
		HandlerFunc(i.GetPetname)

	i.Router.
		Name("AddSeed").
		Path("/v1/api/{api}/add").
		HandlerFunc(AddSeed)

	i.Router.
		Name("ReloadAPI").
		Path("/v1/api/{api}/reload").
		HandlerFunc(ReloadAPI)

	i.Router.
		Name("RemoveSeed").
		Path("/v1/api/{api}/remove").
		HandlerFunc(RemoveSeed)

	i.Router.
		Name("GetAllAPI").
		Path("/v1/api/{api}").
		HandlerFunc(GetAllAPI)

}
