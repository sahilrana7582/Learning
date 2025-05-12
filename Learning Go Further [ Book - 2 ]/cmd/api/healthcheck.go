package main

import (
	"fmt"
	"net/http"
	"time"
)

func (app *applicaton) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "status: %s\n", http.StatusText(http.StatusOK))
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "date: %s\n", time.Now().Format(time.RFC3339))
	fmt.Fprintf(w, "version: %s\n", version)
	fmt.Fprintf(w, "port: %d\n", app.config.port)
	fmt.Fprintf(w, "uptime: %s\n", time.Since(time.Now().Add(-time.Hour)))
	w.Header().Set("Content-Type", "application/json")

}
