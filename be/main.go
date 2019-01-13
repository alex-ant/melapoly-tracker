package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/alex-ant/melapoly-tracker/be/api"
	"github.com/alex-ant/melapoly-tracker/be/params"
	"github.com/alex-ant/melapoly-tracker/be/players"
)

var (
	apiServer   *api.API
	playersProc *players.Players
)

func shutdown() {
	log.Println("Shutting down gracefully..")

	// Stop the API.
	apiServer.Stop()

	log.Println("terminating process")
	os.Exit(0)
}

func main() {
	// Initialize players processor.
	playersProc = players.New(*params.InitialAmount, *params.Salary)

	// Initialize API HTTP server.
	apiServer = api.New(*params.APIPort, playersProc)

	// Start API HTTP server.
	apiStartErr := apiServer.Start()
	if apiStartErr != nil {
		log.Fatal(apiStartErr)
	}

	// Shut down on SIGINT.
	go func() {
		intChan := make(chan os.Signal)
		signal.Notify(intChan, syscall.SIGINT, syscall.SIGTERM)
		<-intChan
		go shutdown()

		// Another signal will force process termination.
		signal.Notify(intChan, syscall.SIGINT, syscall.SIGTERM)
		<-intChan
		os.Exit(0)
	}()

	log.Println("Successfully started")

	// Keep the process running.
	select {}
}
