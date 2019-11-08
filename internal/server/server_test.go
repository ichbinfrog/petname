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
		t.Errorf("[%s] Encountered error while querying , see %s\n", p, err.Error())
	}
	if r.StatusCode != code {
		t.Errorf("[%s] Status code should be %d\n", p, code)
	}
}

func handleGet(t *testing.T, path string, port string, code int, fail bool) {
	r, err := http.Get("http://localhost:" + port + path)
	if err != nil {
		t.Errorf("[%s] Encountered error while performing GET, see %s\n", path, err.Error())
	}

	if r.StatusCode != code {
		t.Errorf("[%s] Status code should be %d, got %d\n", path, code, r.StatusCode)
	}

	defer r.Body.Close()
	body, openErr := ioutil.ReadAll(r.Body)
	if openErr != nil {
		t.Errorf("[%s] Response body should be readable by client\n", path)
	}
	if !fail && string(body) == "" {
		t.Errorf("[%s] Get request should return non empty string\n", path)
	}
	fmt.Printf("[%s] Get request answered with %s\n", path, string(body))
}

func handleGetParam(t *testing.T, path string, prt string, code int, fail bool, param map[string][]string) {
	req, err := http.NewRequest("GET", "http://localhost:"+prt+path, nil)
	if err != nil {
		t.Errorf("[%s?%+v] Encountered error while building query, see %s\n", path, param, err.Error())
	}

	q := req.URL.Query()
	for k, v := range param {
		if len(v) > 1 {
			for _, subV := range v {
				q.Add(k, subV)
			}
		} else if len(v) == 1 {
			q.Add(k, v[0])
		} else {
			q.Add(k, "")
		}
	}
	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	r, reqErr := client.Do(req)
	if reqErr != nil {
		t.Errorf("[%s?%+v] Encountered error while executing query, see %s\n", path, param, err.Error())
	}
	if r.StatusCode != code {
		t.Errorf("[%s?%+v] Status code should be %d, got %d\n", path, param, code, r.StatusCode)
	}

	defer r.Body.Close()
	body, openErr := ioutil.ReadAll(r.Body)
	if openErr != nil {
		t.Errorf("[%s?%+v] Response body should be readable by client\n", path, param)
	}
	if !fail && string(body) == "" {
		t.Errorf("[%s?%+v] Get request should return non empty string\n", path, param)
	}
	fmt.Printf("[%s?%+v] Get request answered with %s", path, param, string(body))
}

func TestReturn(t *testing.T) {
	// Test Health handler return code
	handleReturn(t, "/health", prt, http.StatusOK)

	// Test Error handler return code
	handleReturn(t, "/holla", prt, http.StatusNotFound)
}

func TestGet(t *testing.T) {
	// Test get handler for non existing api
	handleGet(t, "/get/0000000000", prt, http.StatusNotFound, true)
	// Test get handler for the default api
	handleGet(t, "/get/default", prt, http.StatusOK, false)

	// Test get handler with positive param with the default api
	handleGetParam(t, "/get/default", prt, http.StatusOK, false, map[string][]string{
		"amount": []string{"100"},
	})

	// Test get handler with negative param with the default api
	handleGetParam(t, "/get/default", prt, http.StatusBadRequest, false, map[string][]string{
		"amount": []string{"-100"},
	})

	// Test get handler with null param with the default api
	handleGetParam(t, "/get/default", prt, http.StatusOK, false, map[string][]string{
		"amount": []string{"0"},
	})

	// Test reload used map for default API
	api := i.API["default"]
	if api.Generator.Used.Depth <= 0 {
		t.Errorf("[/api/default/] Should have a tree with depth >= 1 by this point\n")
	}

	handleGet(t, "/api/default/reload", prt, http.StatusOK, true)
	if api.Generator.Used.Depth != 0 {
		t.Errorf("[/api/default/reload] Depth should be nil because tree is cleared\n")
	}

	// Test non existant API reload
	handleGet(t, "/api/0000000000/reload", prt, http.StatusNotFound, false)
}

func TestAPI(t *testing.T) {
	// Test get information from existing API
	handleGet(t, "/api/default", prt, http.StatusOK, false)

	// Test get information from existing API
	handleGet(t, "/api/0000000000", prt, http.StatusNotFound, false)

	// Test add API with no parameters
	handleGetParam(t, "/api", prt, http.StatusBadRequest, false, map[string][]string{})

	// Test add API with no parameters
	handleGetParam(t, "/api", prt, http.StatusBadRequest, false, map[string][]string{})

	// Test add API with failed lock conversion
	handleGetParam(t, "/api", prt, http.StatusBadRequest, false, map[string][]string{
		"lock":      []string{"holla"},
		"name":      []string{"holla"},
		"template":  []string{"holla"},
		"separator": []string{"holla"},
	})

	// Test valid API add
	handleGetParam(t, "/api", prt, http.StatusOK, false, map[string][]string{
		"lock":      []string{"false"},
		"name":      []string{"v1"},
		"template":  []string{"{{ .Name }}{{ .Adverb }}{{ .Adverb }}"},
		"separator": []string{"~~~~"},
	})

	// Test duplicate API add
	handleGetParam(t, "/api", prt, http.StatusBadRequest, false, map[string][]string{
		"lock":      []string{"false"},
		"name":      []string{"v1"},
		"template":  []string{"{{ .Name }}{{ .Adverb }}{{ .Adverb }}"},
		"separator": []string{"~~~~"},
	})
}

func TestSeed(t *testing.T) {
	// Test empty seed add
	handleGetParam(t, "/api/default/add", prt, http.StatusBadRequest, false, map[string][]string{})

	// Test seed add with non existant values
	handleGetParam(t, "/api/default/add", prt, http.StatusBadRequest, false, map[string][]string{
		"type": []string{"holla"},
	})

	// Test seed add with non existant API
	handleGetParam(t, "/api/0000000000/add", prt, http.StatusBadRequest, false, map[string][]string{
		"type": []string{"holla"},
	})

	// Test empty value type with existing api
	handleGetParam(t, "/api/default/add", prt, http.StatusBadRequest, true, map[string][]string{
		"type":  []string{"holla"},
		"value": []string{},
	})

	// Test non existant type seed add with default API
	handleGetParam(t, "/api/default/add", prt, http.StatusBadRequest, true, map[string][]string{
		"type":  []string{"holla"},
		"value": []string{"holla", "como"},
	})

	// Test adj seed add with default API
	oldAdj := len(i.API["default"].Generator.Adjectives)
	handleGetParam(t, "/api/default/add", prt, http.StatusOK, true, map[string][]string{
		"type":  []string{"adj"},
		"value": []string{"holla", "como"},
	})
	if len(i.API["default"].Generator.Adjectives) != oldAdj+2 {
		t.Errorf("[/api/default/add?type=adj&value=holla&value=como] should have %d entries, but has %d\n", oldAdj+2, len(i.API["default"].Generator.Adjectives))
	}

	// Test adv seed add with default API
	oldAdv := len(i.API["default"].Generator.Adverbs)
	handleGetParam(t, "/api/default/add", prt, http.StatusOK, true, map[string][]string{
		"type":  []string{"adv"},
		"value": []string{"holla", "como"},
	})
	if len(i.API["default"].Generator.Adverbs) != oldAdv+2 {
		t.Errorf("[/api/default/add?type=adv&value=holla&value=como] should have %d entries, but has %d\n", oldAdv+2, len(i.API["default"].Generator.Adverbs))
	}

	// Test name seed add with default API
	oldName := len(i.API["default"].Generator.Names)
	handleGetParam(t, "/api/default/add", prt, http.StatusOK, true, map[string][]string{
		"type":  []string{"name"},
		"value": []string{"holla", "como"},
	})
	if len(i.API["default"].Generator.Names) != oldName+2 {
		t.Errorf("[/api/default/add?type=name&value=holla&value=como] should have %d entries, but has %d\n", oldName+2, len(i.API["default"].Generator.Names))
	}

}
