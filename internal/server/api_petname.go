package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// GetPetname returns petname when queried
func (i *Instance) GetPetname(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	query := r.URL.Query()
	nb := 1

	if query["amount"] != nil && len(query["amount"]) == 1 {
		nb, err := strconv.Atoi(query["amount"][0])
		if err != nil {
			http.Error(w, "Amount parameter must be positive", http.StatusBadRequest)
		}

		if nb < 0 {
			http.Error(w, "/get/{api} : amount must be > 0", http.StatusBadRequest)
		}
	}

	if a, ok := i.API[mux.Vars(r)["api"]]; ok {
		g := make([]string, nb)
		var err error
		for i := 0; i < nb; i++ {
			g[i], err = a.Generator.Get()

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
		json.NewEncoder(w).Encode(g)
		return
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}
