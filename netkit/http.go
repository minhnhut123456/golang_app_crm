package netkit

import (
	"fmt"
	"net/http"
	"time"
)

type HTTPServer struct {
	server *http.Server
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewHTTPServer(handlers []*RouteHandler) *HTTPServer{
	mux := NewMux(handlers)

	srv := &http.Server{
		Addr:         ":8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler: mux,
}

	return &HTTPServer{server: srv}
}

func (h *HTTPServer) Start() error{
	fmt.Println("Server is running on port 8080")
	err := h.server.ListenAndServe()
	return err
}