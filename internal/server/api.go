package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ichbinfrog/petname/pkg/response"
	"net/http"
	"strconv"
	"sync"
)

// AddAPI adds an API endpoint
func (i *Instance) AddAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	param := r.URL.Query()
	err := ""
	if len(param) < 1 {
		err = err + response.QueryEmptyParam.Error()
	}

	if param["lock"] == nil || len(param["lock"]) != 1 {
		err = err + response.QueryEmptyLock.Error()
	}

	if param["name"] == nil || len(param["name"]) != 1 {
		err = err + response.QueryEmptyName.Error()
	}

	if param["template"] == nil || len(param["template"]) != 1 {
		err = err + response.QueryEmptyTemplate.Error()
	}

	if len(err) != 0 {
		http.Error(w, err, http.StatusBadRequest)
		return
	}

	lock, errConv := strconv.ParseBool(param["lock"][0])
	if errConv != nil {
		http.Error(w, errConv.Error(), http.StatusBadRequest)
		return
	}

	if i.SetupAPI(param["name"][0], lock, param["template"][0]) {
		w.Write([]byte("Successful insert"))
	} else {
		http.Error(w, response.APIAddDuplicateError.Error(), http.StatusBadRequest)
	}
}

// GetInfoAPI returns informations about a specific API
func (i *Instance) GetInfoAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	for _, a := range i.API {
		if a.Name == mux.Vars(r)["api"] {
			w.Write([]byte(fmt.Sprintf("%+v\n", a)))
			return
		}
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

// ReloadAPI cleans the Used binary tree for a specific API
func (i *Instance) ReloadAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if a, ok := i.API[mux.Vars(r)["api"]]; ok {
		a.Generator.Used = new(sync.Map)
		return
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

const (
	paramAdj  = "adj"
	paramAdv  = "adv"
	paramName = "name"
)

// AddSeed adds a seed to a specific api endpoint
// note that duplicate seed is explicitly allowed in order to allow
// for increasing odds as well as to allow some names to pop up twice
func (i *Instance) AddSeed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	param := r.URL.Query()
	if len(param) < 2 {
		http.Error(w, response.SeedAddParamRequired.Error(), http.StatusBadRequest)
		return
	}
	seedType := param["type"]
	if seedType != nil {
		value := param["value"]
		if value == nil || len(value) < 1 {
			http.Error(w, response.SeedAddValueRequired.Error(), http.StatusBadRequest)
			return
		}

		if a, ok := i.API[mux.Vars(r)["api"]]; ok {
			if seedType[0] == paramAdj {
				a.Generator.Adjectives = append(a.Generator.Adjectives, value...)
			} else if seedType[0] == paramAdv {
				a.Generator.Adverbs = append(a.Generator.Adverbs, value...)
			} else if seedType[0] == paramName {
				a.Generator.Names = append(a.Generator.Names, value...)
			} else {
				http.Error(w, response.SeedAddTypeRequired.Error(), http.StatusBadRequest)
			}
		} else {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	}
}

// RemoveSeed removes a seed to a specific api endpoint
func (i *Instance) RemoveSeed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	param := r.URL.Query()
	if len(param) < 2 {
		http.Error(w, response.SeedRmParamRequired.Error(), http.StatusBadRequest)
		return
	}

	seedType := param["type"]
	if seedType != nil {
		value := param["value"]
		if value == nil || len(value) < 1 {
			http.Error(w, response.SeedRmValueRequired.Error(), http.StatusBadRequest)
			return
		}

		if a, ok := i.API[mux.Vars(r)["api"]]; ok {
			if seedType[0] == paramAdj {
				a.Generator.Adjectives = removeSlice(a.Generator.Adjectives, value)
			} else if seedType[0] == paramAdv {
				a.Generator.Adverbs = removeSlice(a.Generator.Adverbs, value)
			} else if seedType[0] == paramName {
				a.Generator.Names = removeSlice(a.Generator.Names, value)
			} else {
				http.Error(w, response.SeedRmTypeRequired.Error(), http.StatusBadRequest)
			}
		} else {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	}
}

func removeSlice(a []string, k []string) []string {
	for _, v := range k {
		a = removeValue(a, v)
	}
	return a
}

func removeValue(a []string, k string) []string {
	for i, v := range a {
		if v == k {
			copy(a[i:], a[i+1:])
			a[len(a)-1] = ""
			return a[:len(a)-1]
		}
	}
	return a
}
