package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/sahilrana7582/Learning/internal/data"
	"github.com/sahilrana7582/Learning/internal/db"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
	models *data.Models
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	logger := log.New(os.Stdout, "[info]", log.Ldate|log.Ltime|log.Lshortfile)

	cfg2 := db.DBConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "1234",
		DBName:   "go_movies_db",
		SSLMode:  "disable",
	}

	// Initialize the database connection
	database, err := db.NewDB(cfg2)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}
	defer database.Close()
	logger.Printf("database connection pool established")

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModal(database),
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.recoverPanic(app.routes()),
		IdleTimeout:  5 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Printf("Starting %s server on port %d", app.config.env, app.config.port)
	err = srv.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}
