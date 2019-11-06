package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func (i *Instance) GetPetname(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	for _, a := range i.API {
		if a.Name == mux.Vars(r)["api"] {
			json.NewEncoder(w).Encode(a.Generator.Get())
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
