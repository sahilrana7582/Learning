package main

import (
	"net/http"
	"time"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"status":  "available",
		"env":     app.config.env,
		"version": version,
		"time":    time.Now().Format("2006-01-02 15:04:05"),
	}

	err := app.writeJson(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Println("Error writing JSON response:", err)
		app.serverErrorResponse(w, r, err)
	}
}
