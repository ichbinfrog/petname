package petname

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/buaazp/fasthttprouter"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
)

type Router struct {
	Router *fasthttprouter.Router
	APIs   map[string]*Generator
}

var (
	strApplicationJSON = []byte("application/json")
	strContentType     = []byte("Content-Type")
)

// AddAPI adds an API endpoint
func (r *Router) AddAPI(ctx *fasthttp.RequestCtx) {
	ctx.SetContentTypeBytes(strApplicationJSON)
	ctx.SetStatusCode(fasthttp.StatusOK)

	param := ctx.QueryArgs()
	if param.Len() < 1 {
		ctx.Error(
			ErrQueryEmptyParam.Error(),
			http.StatusBadRequest,
		)
		return
	}
	name := string(param.Peek("name"))
	template := string(param.Peek("template"))
	if _, ok := r.APIs[name]; ok {
		ctx.Error(ErrAPIAddDuplicate.Error(), http.StatusBadRequest)
		return
	}

	var err error
	r.APIs[name], err = NewGenerator("", template, 5)
	if err != nil {
		ctx.Error(err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

// GetInfoAPI returns informations about a specific API
func (r *Router) GetInfoAPI(ctx *fasthttp.RequestCtx) {
	ctx.SetContentTypeBytes(strApplicationJSON)
	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)

	if api, ok := r.APIs[ctx.UserValue("api").(string)]; ok {
		body, err := json.Marshal(api)
		if err != nil {
			ctx.Error(ErrFailedJSONUnmarshal.Error(), http.StatusInternalServerError)
			return
		}
		ctx.Write(body)
		return
	}
	ctx.Error(http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

// ReloadAPI cleans the Cache for a specific API
func (r *Router) ReloadAPI(ctx *fasthttp.RequestCtx) {
	ctx.SetContentTypeBytes(strApplicationJSON)
	ctx.SetStatusCode(fasthttp.StatusOK)

	if api, ok := r.APIs[ctx.UserValue("api").(string)]; ok {
		api.Cache.Clear()
		return
	}

	ctx.Error(http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

// GetPetname returns petname when queried
func (r *Router) GetPetname(ctx *fasthttp.RequestCtx) {
	ctx.SetContentTypeBytes(strApplicationJSON)
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Response.Header.SetBytesKV(strContentType, strApplicationJSON)

	query := ctx.QueryArgs()

	var err error
	nb := 1
	if query.Has("amount") {
		nb, err = query.GetUint("amount")
		if err != nil {
			ctx.Error(
				ErrQueryAmountInvalid.Error(),
				http.StatusBadRequest,
			)
		}
	}

	if api, ok := r.APIs[ctx.UserValue("api").(string)]; ok {
		nameChan := make(chan *string)
		go api.Generate(nb, nameChan)
		res := []string{}
		for name := range nameChan {
			res = append(res, *name)
		}
		if err := json.NewEncoder(ctx).Encode(res); err != nil {
			ctx.Error(ErrFailedJSONUnmarshal.Error(), http.StatusInternalServerError)
		}
		return
	}

	ctx.Error(http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

// SetupRouter generates the initial router configurations
// for the petname API server
func (r *Router) setupRouter() {
	r.Router = fasthttprouter.New()
	r.APIs = make(map[string]*Generator)

	var err error
	r.APIs["default"], err = NewGenerator("", "{{ Adverb }}-{{ Adjective }}-{{ Name }}", 5)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to build default API")
	}

	r.Router.GET("/get/:api", r.GetPetname)
	r.Router.GET("/api/:api", r.GetInfoAPI)
	r.Router.GET("/api", r.AddAPI)
}

// Start loads router configuration and starts the http listening
func (r *Router) Start(port int) {
	r.setupRouter()
	log.Info().Int("port", port).Msg("Serving on")

	log.Info().
		Str("path", "localhost/get/:api?amount={nb}").
		Msg("	Get {nb} petname")
	log.Info().
		Str("path", "localhost/api?name={name},template={template}").
		Msg("	Create a new API")
	log.Info().
		Str("path", "localhost/api/:api").
		Msg("	Fetch information on a given API")

	fasthttp.ListenAndServe(fmt.Sprintf(":%d", port), r.Router.Handler)
}
