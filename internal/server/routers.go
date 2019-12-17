package server

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

// Instance is a structure that stores the router
type Instance struct {
	Router *fasthttprouter.Router
	API    map[string]API
}

// SetupRouter generates the initial router configurations
// for the petname API server
func (i *Instance) SetupRouter() {
	i.Router = fasthttprouter.New()
	i.API = make(map[string]API)
	i.SetupAPI("default", false, "{{ Adverb }}-{{ Adjective }}-{{ Name }}")

	i.Router.GET("/health", HealthGet)
	i.Router.GET("/get/:api", i.GetPetname)
	i.Router.GET("/api/:api/add", i.AddSeed)
	i.Router.GET("/api/:api/reload", i.ReloadAPI)
	i.Router.GET("/api/:api/remove", i.RemoveSeed)
	i.Router.GET("/api/:api", i.GetInfoAPI)
	i.Router.GET("/api", i.AddAPI)
}

// Start loads router configuration and starts the http listening
func (i *Instance) Start(port int) {
	i.SetupRouter()
	fasthttp.ListenAndServe(fmt.Sprintf(":%d", port), i.Router.Handler)
}
