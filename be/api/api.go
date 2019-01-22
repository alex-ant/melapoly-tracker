package api

import (
	"net"
	"net/http"
	"strconv"

	"github.com/alex-ant/color-id"
	"github.com/alex-ant/melapoly-tracker/be/players"
	"github.com/go-zoo/bone"
	"github.com/jcuga/golongpoll"
)

// API contains HTTP server's settings.
type API struct {
	port        int
	listener    net.Listener
	mux         *bone.Mux
	playersProc *players.Players
	lpManager   *golongpoll.LongpollManager
	lpUpdatedTS int64
	colorSet    *color.Set
}

// New returns new API.
func New(port int, playersProc *players.Players) *API {
	return &API{
		port:        port,
		playersProc: playersProc,
		colorSet:    color.NewSet(),
	}
}

func (a *API) defineMux() error {
	a.mux = bone.New()

	var lpManagerErr error
	a.lpManager, lpManagerErr = golongpoll.StartLongpoll(golongpoll.Options{})
	if lpManagerErr != nil {
		return lpManagerErr
	}

	a.mux.HandleFunc("/lp", func(w http.ResponseWriter, r *http.Request) {
		setCORSHeaders(w)
		a.lpManager.SubscriptionHandler(w, r)
	})

	a.mux.Post("/player", http.HandlerFunc(a.addPlayerHandler))

	a.mux.Delete("/player", http.HandlerFunc(a.removePlayerHandler))
	a.mux.Options("/player", http.HandlerFunc(a.corsRequestHandler))

	a.mux.Get("/player/:token", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setCORSHeaders(w)
		a.authHandler(w, r)
	}))

	a.mux.Get("/players", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setCORSHeaders(w)
		a.getAllPlayersHandler(w, r)
	}))
	a.mux.Options("/players", http.HandlerFunc(a.corsRequestHandler))

	a.mux.Get("/game/info", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setCORSHeaders(w)
		a.getGameInfoHandler(w, r)
	}))

	a.mux.Post("/cash/add", http.HandlerFunc(a.addCashHandler))
	a.mux.Options("/cash/add", http.HandlerFunc(a.corsRequestHandler))

	a.mux.Post("/salary/add", http.HandlerFunc(a.addSalaryHandler))
	a.mux.Options("/salary/add", http.HandlerFunc(a.corsRequestHandler))

	a.mux.Post("/cash/deduct", http.HandlerFunc(a.deductCashHandler))
	a.mux.Options("/cash/deduct", http.HandlerFunc(a.corsRequestHandler))

	a.mux.Post("/cash/transfer", http.HandlerFunc(a.transferCashHandler))
	a.mux.Options("/cash/transfer", http.HandlerFunc(a.corsRequestHandler))

	a.mux.Get("/transactions", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setCORSHeaders(w)
		a.getTransactionsHandler(w, r)
	}))

	return nil
}

// Start starts the HTTP server.
func (a *API) Start() (err error) {
	err = a.defineMux()
	if err != nil {
		return err
	}

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
