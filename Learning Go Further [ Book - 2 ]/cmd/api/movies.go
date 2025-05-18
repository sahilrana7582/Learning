package main

import (
	"encoding/json"
	"net/http"

	"github.com/sahilrana7582/Learning/internal/data"
	"github.com/sahilrana7582/Learning/internal/validator"
)

func (app *application) createNewMovieHandler(w http.ResponseWriter, r *http.Request) {
	var movieData data.Movie

	err := json.NewDecoder(r.Body).Decode(&movieData)
	if err != nil {
		app.logger.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	v := validator.New()
	if v := validator.ValidateMovieInput(v, movieData); !v.Valid() {
		app.logger.Println("Validation error:", v.Errors)
		app.errorResponse(w, r, http.StatusUnprocessableEntity, v.Errors)
		return
	}

	err = app.models.Movies.Insert(&movieData)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	resp := map[string]interface{}{
		"status":   "success",
		"message":  "Movie created successfully",
		"movie_id": movieData.ID,
	}

	err = app.writeJson(w, http.StatusCreated, resp, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getMovieHandler(w http.ResponseWriter, r *http.Request) {
	movies, err := app.models.Movies.GetAll()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJson(w, http.StatusOK, movies, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getMovieById(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdParamFromRequest(r)
	if err != nil {
		app.logger.Println("Invalid movie ID:", err)
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	movie, err := app.models.Movies.Get(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJson(w, http.StatusOK, movie, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdParamFromRequest(r)
	if err != nil {
		app.logger.Println("Invalid movie ID:", err)
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	var movieData data.Movie
	err = json.NewDecoder(r.Body).Decode(&movieData)
	if err != nil {
		app.logger.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Validate input
	v := validator.New()
	if v := validator.ValidateMovieInput(v, movieData); !v.Valid() {
		app.logger.Println("Validation error:", v.Errors)
		app.errorResponse(w, r, http.StatusUnprocessableEntity, v.Errors)
		return
	}

	movieData.ID = id // ensure path ID and payload ID match

	err = app.models.Movies.Update(&movieData)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	resp := map[string]interface{}{
		"status":  "success",
		"message": "Movie updated successfully",
	}

	err = app.writeJson(w, http.StatusOK, resp, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdParamFromRequest(r)
	if err != nil {
		app.logger.Println("Invalid movie ID:", err)
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	err = app.models.Movies.Delete(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	resp := map[string]interface{}{
		"status":  "success",
		"message": "Movie deleted successfully",
	}

	err = app.writeJson(w, http.StatusOK, resp, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
