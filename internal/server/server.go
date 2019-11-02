package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Instance is a structure that stores the router
type Instance struct {
	Router *mux.Router
	API    []API
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
	i.Router.Path("/api").HandlerFunc(i.handleAPI)
	i.Router.Path("/get/{api}").HandlerFunc(i.handleGet)
	i.Router.Path("/get/{api}/{nb:[0-9]+}").HandlerFunc(i.handleGetMul)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusOK), http.StatusOK)
}

func handleReload(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusTeapot), http.StatusTeapot)
}

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusOK), http.StatusOK)
}

func handleError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
}

func (i *Instance) handleGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	for _, a := range i.API {
		if a.Name == mux.Vars(r)["api"] {
			json.NewEncoder(w).Encode(a.Generator.Get())
			return
		}
	}
	http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
}

func (i *Instance) handleGetMul(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	nb, _ := strconv.Atoi(mux.Vars(r)["nb"])
	// if err != nil {
	// 	http.Error(w, "/get/{api}/{nb} : err",
	// 		http.StatusInternalServerError)
	// }
	if nb <= 0 {
		http.Error(w, "/get/{api}/{nb} : nb must be > 0",
			http.StatusInternalServerError)
	}

	for _, a := range i.API {
		if a.Name == mux.Vars(r)["api"] {
			g := make([]string, nb)
			for i := 0; i < nb; i++ {
				g[i] = a.Generator.Get()
			}
			json.NewEncoder(w).Encode(g)
			return
		}
	}
	http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
}

func (i *Instance) handleAPI(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	err := ""
	if len(q) < 1 {
		err = "Query param cannot be empty"
	}

	if q["lock"] == nil || len(q["lock"]) != 1 {
		err = "Lock query param cannot be empty"
	}

	if q["name"] == nil || len(q["name"]) != 1 {
		err = "Name query param cannot be empty"
	}

	if q["template"] == nil || len(q["template"]) != 1 {
		err = "Template query param cannot be empty"
	}

	if q["separator"] == nil || len(q["separator"]) != 1 {
		err = "Separator query param cannot be empty"
	}

	if len(err) != 0 {
		http.Error(w, err, http.StatusInternalServerError)
		return
	}

	lock, _ := strconv.ParseBool(q["lock"][0])
	separator := []rune(q["separator"][0])
	if i.SetupAPI(q["name"][0], lock, q["template"][0], separator[0]) {
		w.Write([]byte("Successful insert"))
	} else {
		http.Error(w, "Failed insert due to duplicate", http.StatusInternalServerError)
	}
}

// Start loads router configuration and starts the http listening
func (i *Instance) Start(port int) {
	i.SetupRouter()
	i.SetupAPI("v1", false, "{{.Name}}{{.Adjective}}{{.Adjective}}", '+')
	http.ListenAndServe(fmt.Sprintf(":%d", port), i.Router)
}
