package main

import (
	"log"
	"os"
	"strconv"
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

	server, err := NewServer(port)
	if err != nil {
		log.Fatalf("Failed to create new Server object: %v\n", err)
	}
	server.Start()
}
