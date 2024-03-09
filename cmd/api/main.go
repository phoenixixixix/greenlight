package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0" // Current API version. TODO: generate version automatically at build time

// Will be used to hold config settings for the app.
type config struct {
	port int
	env  string
}

// This is defined to hold dependencies for HTTP handlers, helpers, and middleware
type application struct {
	config config      // copy
	logger *log.Logger // log object
}

func main() {
	var cfg config // new instance

	// Flags which can be specified on launch
	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse() // Must be called affter all fags are defined

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime) // *log.Logger

	// dependencies
	app := &application{
		config: cfg,
		logger: logger,
	}

	// HTTP server configuration
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Minute,
		WriteTimeout: 30 * time.Minute,
	}

	// Launch srv
	logger.Printf("Starting %s server on %s port", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
