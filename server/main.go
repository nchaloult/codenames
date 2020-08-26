package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/nchaloult/codenames/api"
	"github.com/nchaloult/codenames/realtime"
)

const defaultPort = 6969

func main() {
	// Listen for OS interrupts (like if someone presses Ctrl+C or something).
	// Spin down gracefully if this process is interrupted.
	sigintChan := make(chan os.Signal, 1)
	signal.Notify(sigintChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		os.Interrupt,
		os.Kill,
	)
	go func() {
		<-sigintChan

		log.Println("Spinning down....")
		os.Exit(0)
	}()

	// Get the port that the web server listens on.
	portAsStr := os.Getenv("PORT")
	var port int
	if portAsStr != "" {
		// os.Getenv() returns an environment variable's value as a string.
		// Convert this to an int.
		port, _ = strconv.Atoi(portAsStr)
	} else {
		port = defaultPort
	}

	// Create a new Manager to manage active games.
	manager := realtime.NewManager()

	// Set up HTTP endpoint handlers.
	router := mux.NewRouter()
	a, err := api.NewAPI(router, port)
	if err != nil {
		log.Fatalf("failed to stand up HTTP endpoints: %v", err)
	}
	a.ListenOnEndpoints([]api.RouteHandler{
		api.NewHealthHandler(manager),
		api.NewWSHandler(manager),
	})
}
