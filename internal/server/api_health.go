package server

import (
	"github.com/valyala/fasthttp"
	"net/http"
)

// HealthGet is the health handler (always return 200)
func HealthGet(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json; charset=UTF-8")
	ctx.SetStatusCode(http.StatusOK)
}
