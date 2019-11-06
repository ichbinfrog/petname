package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

const (
	prt = "8000"
)

var (
	i = &Instance{}
)

func init() {
	i.SetupRouter()

	go func() {
		http.ListenAndServe(":8000", i.Router)
	}()
}

func handleReturn(t *testing.T, p string, prt string, code int) {
	r, err := http.Get("http://localhost:" + prt + p)
	if err != nil {
		t.Errorf("[%s] Encountered error while querying , see %s", p, err.Error())
	}
	if r.StatusCode != code {
		t.Errorf("[%s] Status code should be %d", p, code)
	}
}

func handleGet(t *testing.T, path string, port string, code int, fail bool) {
	r, err := http.Get("http://localhost:" + port + path)
	if err != nil {
		t.Errorf("[%s] Encountered error while performing GET, see %s", path, err.Error())
	}

	if r.StatusCode != code {
		t.Errorf("[%s] Status code should be %d", path, code)
	}

	defer r.Body.Close()
	body, openErr := ioutil.ReadAll(r.Body)
	if openErr != nil {
		t.Errorf("[%s] Response body should be readable by client", path)
	}
	if !fail && string(body) == "" {
		t.Errorf("[%s] Get request should return non empty string", path)
	}
	fmt.Printf("[%s] Get request answered with %s", path, string(body))
}

func TestReturn(t *testing.T) {
	// Test Health handler return code
	handleReturn(t, "/v1/health", prt, http.StatusOK)

	// Test Error handler return code
	handleReturn(t, "/v1/holla", prt, http.StatusNotFound)
}

func TestGet(t *testing.T) {
	// Test get handler for non existing api
	// handleGet(t, "/v1", "/api/0000000000", prt, http.StatusNotFound, true)
	// Test get handler for the default api
	handleGet(t, "/get/default", prt, http.StatusOK, false)
}
