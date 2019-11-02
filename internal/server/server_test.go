package server

import (
	"net/http"
	"strconv"
	"testing"
)

func handle(t *testing.T, p string, prt int, code int) {
	r, err := http.Get("http://localhost:" + strconv.Itoa(prt) + p)
	if err != nil {
		t.Errorf("Encountered error while querying /health , see " + err.Error())
	}

	if r.StatusCode != code {
		t.Errorf("Status code should be " + strconv.Itoa(code) + " for path " + p)
	}
}

func TestServer(t *testing.T) {
	// Test initialization
	i := Instance{}
	prt := 8000

	// Tests server start up
	go func() {
		i.Start(prt)
	}()

	// Test Health handler return code
	handle(t, "/health", prt, http.StatusOK)

	// Test Reload handler return code
	handle(t, "/reload", prt, http.StatusTeapot)

	// Test Readiness probe return code
	handle(t, "/ready", prt, http.StatusOK)

	// Test Error handler return code
	handle(t, "/holla", prt, http.StatusForbidden)
}
