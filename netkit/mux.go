package netkit

import (
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

// wrapper and handle http multiplexer
type Mux struct {
	*mux.Router
}


func NewMux(handlers []*RouteHandler) *Mux {
	mux := Mux{Router: &mux.Router{}}

	for _,h := range handlers {
		mux.HandleFunc(h.Route.Path, h.Handler).Methods(h.Route.Method)
	}

	return &mux
}


func Vars(r *http.Request) map[string]string {
	return mux.Vars(r)
}

func Queries(r *http.Request) url.Values {
	return r.URL.Query()
}