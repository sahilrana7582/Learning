package main

import (
	"fmt"
	"net/http"
)

func (app *applicaton) createNewMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, `{"status": "success", "message": "Movie created successfully"}`)
}

func (app *applicaton) getMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status": "success", "message": "Movie retrieved successfully"}`)
}

func (app *applicaton) getMovieById(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIdParamFromRequest(r)
	if err != nil {
		app.logger.Println("Invalid movie ID")
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status": "success", "message": "Movie with ID %d retrieved successfully"}`, id)
}
