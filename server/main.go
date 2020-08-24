package main

import (
	"log"
	"os"
	"strconv"

	"github.com/nchaloult/codenames/server"
)

const defaultPort = 6969

func main() {
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

	// Spin up the server and all of its goroutines.
	s := server.NewServer()
	err := s.Start(port)
	log.Fatalf("failed to spin up the server: %v\n", err)
}
