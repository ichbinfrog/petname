package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/ichbinfrog/petname/pkg/response"
	"net/http"
	"strconv"
	"sync"
)

// GetPetname returns petname when queried
func (i *Instance) GetPetname(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	query := r.URL.Query()
	var err error
	nb := 1

	if query["amount"] != nil && len(query["amount"]) == 1 {
		nb, err = strconv.Atoi(query["amount"][0])
		if err != nil {
			http.Error(w, response.QueryAmountInvalid, http.StatusBadRequest)
		}

		if nb < 0 {
			http.Error(w, response.QueryAmountInvalid, http.StatusBadRequest)
		}
	}

	if a, ok := i.API[mux.Vars(r)["api"]]; ok {
		var g []string
		var wg sync.WaitGroup
		q := make(chan string, 1)

		for i := 0; i < nb; i++ {
			wg.Add(1)
			go func() {
				s, err := a.Generator.Get()
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					q <- ""
				}
				q <- s
			}()
		}

		go func() {
			for s := range q {
				g = append(g, s)
				wg.Done()
			}
		}()

		wg.Wait()
		json.NewEncoder(w).Encode(g)
		return
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}
