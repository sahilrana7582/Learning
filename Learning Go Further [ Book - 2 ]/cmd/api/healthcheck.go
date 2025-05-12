package main

import (
	"fmt"
	"net/http"
	"time"
)

func (app *applicaton) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	resp := `{
		"status": "success",
		"environment": %q,
		"version": %q,
		"timestamp": %q,
		}`

	resp = fmt.Sprintf(resp, app.config.env, version, time.Now().Format(time.RFC3339))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
	app.logger.Println("Healthcheck request received")
	app.logger.Println("Healthcheck response sent")

}
