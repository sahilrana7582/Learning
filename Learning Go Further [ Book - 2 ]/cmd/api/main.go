package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sahilrana7582/Learning/internal/data"
	"github.com/sahilrana7582/Learning/internal/db"
	"github.com/sahilrana7582/Learning/internal/mailer"
)

const version = "1.0.0"

type smtpConfig struct {
	host     string
	port     int
	username string
	password string
	sender   string
}

type config struct {
	port int
	env  string
	smtp smtpConfig
}

type application struct {
	config config
	logger *log.Logger
	models *data.Models
	mailer mailer.Mailer
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	logger := log.New(os.Stdout, "[info] ", log.Ldate|log.Ltime|log.Lshortfile)

	cfg2 := db.DBConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "1234",
		DBName:   "go_movies_db",
		SSLMode:  "disable",
	}

	database, err := db.NewDB(cfg2)
	if err != nil {
		logger.Fatalf("DB connection failed: %v", err)
	}
	defer database.Close()

	flag.StringVar(&cfg.smtp.host, "smtp-host", "sandbox.smtp.mailtrap.io", "SMTP host")
	flag.IntVar(&cfg.smtp.port, "smtp-port", 2525, "SMTP port")
	flag.StringVar(&cfg.smtp.username, "smtp-username", "c25991a5c81003", "SMTP username")
	flag.StringVar(&cfg.smtp.password, "smtp-password", "77287bdf4c9c6b", "SMTP password")
	flag.StringVar(&cfg.smtp.sender, "smtp-sender", "Greenlight <noreply@greenlight.local>", "SMTP sender")
	flag.Parse()

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModal(database),
		mailer: mailer.New(
			cfg.smtp.host,
			cfg.smtp.port,
			cfg.smtp.username,
			cfg.smtp.password,
			cfg.smtp.sender,
		),
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.recoverPanic(app.routes()),
		IdleTimeout:  5 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Create a channel to wait for shutdown to complete
	shutdownComplete := make(chan struct{})

	// Start shutdown listener in goroutine
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

		s := <-sig
		logger.Printf("Received shutdown signal: %s", s)

		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		logger.Println("Attempting graceful shutdown")
		if err := srv.Shutdown(ctx); err != nil {
			logger.Fatalf("Could not gracefully shutdown: %v", err)
		}

		logger.Println("Shutdown complete")
		close(shutdownComplete)
	}()

	logger.Printf("Starting %s server on port %d", app.config.env, app.config.port)
	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Server error: %v", err)
	}

	// Wait for shutdown to finish
	<-shutdownComplete
	logger.Println("Server has exited cleanly")
}
