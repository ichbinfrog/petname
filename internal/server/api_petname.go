package server

import (
	"encoding/json"
	"github.com/ichbinfrog/petname/pkg/response"
	"github.com/valyala/fasthttp"
	"net/http"
	"sync"
)

// GetPetname returns petname when queried
func (i *Instance) GetPetname(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json; charset=UTF-8")
	query := ctx.QueryArgs()
	var err error

	nb, err := query.GetUint("amount")
	if !query.Has("amount") {
		nb = 1
	} else {
		if err != nil {
			ctx.Error(response.QueryAmountInvalid, http.StatusBadRequest)
		}
	}

	if a, ok := i.API[ctx.UserValue("api").(string)]; ok {
		var g []string
		var wg sync.WaitGroup
		q := make(chan string, 1)

		for i := 0; i < nb; i++ {
			wg.Add(1)
			go func() {
				s, err := a.Generator.Get()
				if err != nil {
					ctx.Error(err.Error(), http.StatusInternalServerError)
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
		json.NewEncoder(ctx).Encode(g)
		return
	}

	ctx.Error(http.StatusText(http.StatusNotFound), http.StatusNotFound)
}
