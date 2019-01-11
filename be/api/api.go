package api

import (
	"net"
	"net/http"
	"strconv"

	"github.com/alex-ant/melapoly-tracker/be/players"
	"github.com/go-zoo/bone"
)

// API contains HTTP server's settings.
type API struct {
	port        int
	listener    net.Listener
	mux         *bone.Mux
	playersProc *players.Players
}

// New returns new API.
func New(port int, playersProc *players.Players) *API {
	return &API{
		port:        port,
		playersProc: playersProc,
	}
}

func (a *API) defineMux() {
	a.mux = bone.New()

	a.mux.Post("/auth", http.HandlerFunc(a.authHandler))
	a.mux.Options("/auth", http.HandlerFunc(a.corsRequestHandler))
}

// Start starts the HTTP server.
func (a *API) Start() (err error) {
	a.defineMux()

	a.listener, err = net.Listen("tcp", ":"+strconv.Itoa(a.port))
	if err != nil {
		return
	}

	go http.Serve(a.listener, a.mux)

	return
}

// Stop stops the server.
func (a *API) Stop() {
	a.listener.Close()
}
