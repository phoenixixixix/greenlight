package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const version = "1.0.0" // Current API version. TODO: generate version automatically at build time

// Will be used to hold config settings for the app.
type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
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
	// Postgres options: User: greenlight, DB: greenlight, password: pa55word, ssl disabled
	// TODO: move to env variable
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://greenlight:pa55word@localhost/greenlight?sslmode=disable", "PostgreSQL DSN")
	// Flags to specify SQL connection pool options
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-mux-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-mux-idle-time", "15m", "PostgreSQL max connection idle time")
	flag.Parse() // Must be called affter all fags are defined

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime) // *log.Logger

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()
	logger.Printf("database connection pool established")

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
	err = srv.ListenAndServe()
	logger.Fatal(err)
}

// Returns connection pool
func openDB(cfg config) (*sql.DB, error) {
	// create empty connection pool
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	// Set custom options (from flags) to the connection pool
	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)

	// Here I check if connection is establiched. If connection couldn't be
	// establiched successfully within 5 second - return an error
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	// If connection establiched retrun pool
	return db, nil
}
