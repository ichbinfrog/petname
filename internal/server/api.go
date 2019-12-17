package server

import (
	"fmt"
	"github.com/ichbinfrog/petname/pkg/response"
	"github.com/valyala/fasthttp"
	"net/http"
)

// AddAPI adds an API endpoint
func (i *Instance) AddAPI(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json; charset=UTF-8")
	param := ctx.QueryArgs()
	err := ""
	if param.Len() < 1 {
		err = err + response.QueryEmptyParam
	}

	if !param.Has("lock") {
		err = err + response.QueryEmptyLock
	}

	if !param.Has("name") {
		err = err + response.QueryEmptyName
	}

	if !param.Has("template") {
		err = err + response.QueryEmptyTemplate
	}

	if len(err) != 0 {
		ctx.Error(err, http.StatusBadRequest)
		return
	}

	if i.SetupAPI(string(param.Peek("name")), param.GetBool("lock"), string(param.Peek("template"))) {
		ctx.SetStatusCode(http.StatusOK)
	} else {
		ctx.Error(response.APIAddDuplicateError, http.StatusBadRequest)
	}
}

// GetInfoAPI returns informations about a specific API
func (i *Instance) GetInfoAPI(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json; charset=UTF-8")
	for _, a := range i.API {
		if a.Name == ctx.UserValue("api") {
			ctx.Write([]byte(fmt.Sprintf("%+v\n", a)))
			return
		}
	}

	ctx.Error(http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

// ReloadAPI cleans the Used binary tree for a specific API
func (i *Instance) ReloadAPI(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json; charset=UTF-8")

	if a, ok := i.API[ctx.UserValue("api").(string)]; ok {
		a.Generator.Used.Flush()
		return
	}

	ctx.Error(http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

const (
	paramAdj  = "adj"
	paramAdv  = "adv"
	paramName = "name"
)

// AddSeed adds a seed to a specific api endpoint
// note that duplicate seed is explicitly allowed in order to allow
// for increasing odds as well as to allow some names to pop up twice
func (i *Instance) AddSeed(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json; charset=UTF-8")
	param := ctx.QueryArgs()
	if param.Len() < 2 {
		ctx.Error(response.SeedAddParamRequired, http.StatusBadRequest)
		return
	}
	seedType := string(param.Peek("type"))
	if len(seedType) != 0 {
		value := param.PeekMulti("value")
		if len(value) == 0 || (len(value) == 1 && string(value[0]) == "") {
			ctx.Error(response.SeedAddValueRequired, http.StatusBadRequest)
			return
		}

		if a, ok := i.API[ctx.UserValue("api").(string)]; ok {
			if seedType == paramAdj {
				a.Generator.Adjectives = appendSlice(a.Generator.Adjectives, value)
				a.Generator.AvailableAdj++
			} else if seedType == paramAdv {
				a.Generator.Adverbs = appendSlice(a.Generator.Adverbs, value)
				a.Generator.AvailableAdv++
			} else if seedType == paramName {
				a.Generator.Names = appendSlice(a.Generator.Names, value)
				a.Generator.AvailableName++
			} else {
				ctx.Error(response.SeedAddTypeRequired, http.StatusBadRequest)
			}
		} else {
			ctx.Error(http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	}
}

// RemoveSeed removes a seed to a specific api endpoint
func (i *Instance) RemoveSeed(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json; charset=UTF-8")
	param := ctx.QueryArgs()
	if param.Len() < 2 {
		ctx.Error(response.SeedRmParamRequired, http.StatusBadRequest)
		return
	}

	seedType := string(param.Peek("type"))
	if len(seedType) != 0 {
		value := param.PeekMulti("value")
		if len(value) == 0 || (len(value) == 1 && string(value[0]) == "") {
			ctx.Error(response.SeedRmValueRequired, http.StatusBadRequest)
			return
		}

		if a, ok := i.API[ctx.UserValue("api").(string)]; ok {
			if seedType == paramAdj {
				a.Generator.Adjectives = removeSlice(a.Generator.Adjectives, value)
				a.Generator.AvailableAdj--
			} else if seedType == paramAdv {
				a.Generator.Adverbs = removeSlice(a.Generator.Adverbs, value)
				a.Generator.AvailableAdv--
			} else if seedType == paramName {
				a.Generator.Names = removeSlice(a.Generator.Names, value)
				a.Generator.AvailableName--
			} else {
				ctx.Error(response.SeedRmTypeRequired, http.StatusBadRequest)
			}
		} else {
			ctx.Error(http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	}
}

func appendSlice(a []string, k [][]byte) []string {
	for _, v := range k {
		a = append(a, string(v))
	}
	return a
}

func removeSlice(a []string, k [][]byte) []string {
	for _, v := range k {
		a = removeValue(a, string(v))
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
