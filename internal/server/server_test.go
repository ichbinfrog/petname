package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
)

var (
	i   = Instance{}
	prt = 8000
)

func init() {
	go func() {
		i.Start(prt)
	}()
}

func handleReturn(t *testing.T, p string, prt int, code int) {
	r, err := http.Get("http://localhost:" + strconv.Itoa(prt) + p)
	if err != nil {
		t.Errorf("[%s] Encountered error while querying , see %s", p, err.Error())
	}

	if r.StatusCode != code {
		t.Errorf("[%s] Status code should be %d", p, code)
	}
}

func handleGet(t *testing.T, p string, api string, prt int, code int, fail bool) {
	r, err := http.Get("http://localhost:" + strconv.Itoa(prt) + p + api)
	if err != nil {
		t.Errorf("[%s%s] Encountered error while querying %s, see %s", p, api, p, err.Error())
	}

	if r.StatusCode != code {
		t.Errorf("[%s%s] Status code should be %d", p, api, code)
	}

	defer r.Body.Close()
	body, openErr := ioutil.ReadAll(r.Body)
	if openErr != nil {
		t.Errorf("[%s%s] Response body should be readable by client", p, api)
	}
	if !fail && string(body) == "" {
		t.Errorf("[%s%s] Get request should return non empty string", p, api)
	}
	fmt.Printf("[%s%s] Get request answered with %s", p, api, string(body))
}

func handleGetMul(t *testing.T, p string, api string, nb string, prt int, code int, fail bool) {
	r, err := http.Get("http://localhost:" + strconv.Itoa(prt) + p + api + "/" + nb)
	if err != nil {
		t.Errorf("[%s%s/%s] Encountered error while querying %s, see %s", p, api, nb, p, err.Error())
	}

	if r.StatusCode != code {
		t.Errorf("[%s%s/%s] Status code should be %d", p, api, nb, code)
	}

	defer r.Body.Close()
	body, openErr := ioutil.ReadAll(r.Body)
	if openErr != nil {
		t.Errorf("[%s%s/%s] Response body should be readable by client", p, api, nb)
	}
	if !fail && string(body) == "" {
		t.Errorf("[%s%s/%s] Get request should return non empty string", p, api, nb)
	}
	fmt.Printf("[%s%s/%s] Get request answered with %s", p, api, nb, string(body))
}

func TestReturn(t *testing.T) {
	// Test Health handler return code
	handleReturn(t, "/health", prt, http.StatusOK)

	// Test Reload handler return code
	handleReturn(t, "/reload", prt, http.StatusTeapot)

	// Test Readiness probe return code
	handleReturn(t, "/ready", prt, http.StatusOK)

	// Test Error handler return code
	handleReturn(t, "/holla", prt, http.StatusForbidden)
}

func TestGet(t *testing.T) {
	// Test get handler for non existing api
	handleGet(t, "/get", "/v100", prt, http.StatusForbidden, true)

	// Test get handler for an existing api
	handleGet(t, "/get", "/v1", prt, http.StatusOK, false)

	// Test multiget handler for an non-existing api
	handleGetMul(t, "/get", "/v100", "10", prt, http.StatusForbidden, true)

	// Test multiget handler for an existing api
	handleGetMul(t, "/get", "/v1", "10", prt, http.StatusOK, false)

	// Test multiget handler for an existing api without negative nb value
	handleGetMul(t, "/get", "/v1", "0", prt, http.StatusInternalServerError, true)
}

func TestInsert(t *testing.T) {
	// Test empty API insert query
	req, err := http.NewRequest("GET", "http://localhost:8000/api", nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	// Test failed api server generator
	client := &http.Client{}
	r, reqErr := client.Do(req)
	if reqErr != nil {
		t.Errorf(reqErr.Error())
	}
	if r.StatusCode != http.StatusInternalServerError {
		t.Errorf("[/api] with empty param, server should return internal server error\n")
	} else {
		fmt.Printf("[/api] server returned %d\n", r.StatusCode)
	}

	// Test api server generator
	q := req.URL.Query()
	q.Add("name", "v2")
	q.Add("lock", "false")
	q.Add("separator", "-")
	q.Add("template", "{{.Name}}{{.Name}}")
	req.URL.RawQuery = q.Encode()
	r, reqErr = client.Do(req)
	if reqErr != nil {
		t.Errorf(reqErr.Error())
	}
	if r.StatusCode != http.StatusOK {
		t.Errorf("[/api?%v] with empty param, server should return internal server error\n", q)
	} else {
		fmt.Printf("[/api?%v] server returned %d\n", q, r.StatusCode)
	}

	// Test failed insert
	r, reqErr = client.Do(req)
	if reqErr != nil {
		t.Errorf(reqErr.Error())
	}
	if r.StatusCode != http.StatusInternalServerError {
		t.Errorf("[/api?%v] with duplicate name, server should return internal server error\n", q)
	} else {
		fmt.Printf("[/api?%v] server returned %d\n", q, r.StatusCode)
	}
}
