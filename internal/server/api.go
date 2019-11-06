package server

import (
	"net/http"
	"strconv"
)

func (i *Instance) AddAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
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
	if i.SetupAPI(q["name"][0], lock, q["template"][0], q["separator"][0]) {
		w.Write([]byte("Successful insert"))
	} else {
		http.Error(w, "Failed insert due to duplicate", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func AddSeed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetAllAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func ReloadAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func RemoveSeed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
