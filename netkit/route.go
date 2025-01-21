package netkit

import "net/http"

type Handler func(http.ResponseWriter, *http.Request)

type Route struct {
	Name string
	Method string
	Path string
}

type RouteHandler struct {
	Route *Route
	Handler Handler
}
