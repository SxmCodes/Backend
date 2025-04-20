package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"syscall"

	"os/Signal"
)

func main() {
	// loading the config packages.
	// This will call the load() function from the config package. cfg will hold info like port number and other shit.
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load the configuration. %v", err)
	}

	// creating a proxy handler.
	proxyHandler = proxy.NewHandler(cfg.TargetURL) // this will pass the url from our config.

	// creating a new server.
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: proxyHandler,
	}

	// starting http server at specified port. Started a goroutine.
	log.Printf("Starting go sever on port %d forwarding to %s", cfg.Port, cfg.TargetURL)

	if err := server.ListenandServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server can't be loaded! %v", err)
	}

	// waiting for interupt signal to shut down the server.
	quit := make(chan os.Signal, 1) // this will make a channel(like a mailbox) that will recieve signals.

	signal.Notify()
}
