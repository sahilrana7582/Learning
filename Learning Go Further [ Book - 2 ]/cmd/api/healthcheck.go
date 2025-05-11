package main

import (
	"fmt"
	"net/http"
	"time"
)

func (app *applicaton) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	app.logger.Println(`{"status": "available"}`)
	app.logger.Printf("environment: %s", app.config.env)
	app.logger.Printf("version: %s", version)
	app.logger.Printf("port: %d", app.config.port)
	app.logger.Printf("time: %s", time.Now().Format(time.RFC3339))

	fmt.Fprint(w, `{"status": "available"}`)
}
