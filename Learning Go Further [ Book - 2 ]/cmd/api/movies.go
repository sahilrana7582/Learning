package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sahilrana7582/Learning/internal/data"
	"github.com/sahilrana7582/Learning/internal/validator"
)

func (app *applicaton) createNewMovieHandler(w http.ResponseWriter, r *http.Request) {
	var movieData data.Movie
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&movieData)
	if err != nil {
		app.logger.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// ✅ Validate before sending any response
	if v := validator.ValidateMovieInput(movieData); !v.Valid() {
		app.logger.Println("Validation error:", v.Errors)
		app.errorResponse(w, r, http.StatusUnprocessableEntity, v.Errors)
		return
	}

	movieData.ID = 1

	err = app.writeJson(w, http.StatusCreated, movieData, nil)
	if err != nil {
		app.logError(r, err)
		app.serverErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, `{"status": "success", "message": "Movie created successfully", "movie_id": %d}`, movieData.ID)
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

	movieData := data.Movie{
		ID:       id,
		Title:    "Inception",
		Year:     2010,
		Runtime:  148,
		Genre:    []string{"Action", "Sci-Fi"},
		Director: "Christopher Nolan",
		Actors:   []string{"Leonardo DiCaprio", "Joseph Gordon-Levitt"},
		Plot:     "A thief who steals corporate secrets through the use of dream-sharing technology is given the inverse task of planting an idea into the mind of a CEO.",
		Language: "English",
		Country:  "USA",
		Awards:   "Oscar, BAFTA",
	}

	err = app.writeJson(w, http.StatusOK, movieData, nil)
	if err != nil {
		app.logError(r, err)
		app.serverErrorResponse(w, r, err)
	}

}
