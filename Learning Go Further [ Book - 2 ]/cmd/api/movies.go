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

	// Step 1: Fetch existing movie
	movie, err := app.models.Movies.Get(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// Step 2: Parse request body with optional fields
	var input struct {
		Title    *string   `json:"title"`
		Year     *int      `json:"release_year"`
		Runtime  *int      `json:"runtime"`
		Genre    *[]string `json:"genre"`
		Director *string   `json:"director"`
		Actors   *[]string `json:"actors"`
		Plot     *string   `json:"plot"`
		Language *string   `json:"language"`
		Country  *string   `json:"country"`
		Awards   *string   `json:"awards"`
	}

	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.logger.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Step 3: Apply updates to existing movie
	if input.Title != nil {
		movie.Title = *input.Title
	}
	if input.Year != nil {
		movie.Year = *input.Year
	}
	if input.Runtime != nil {
		movie.Runtime = *input.Runtime
	}
	if input.Genre != nil {
		movie.Genre = *input.Genre
	}
	if input.Director != nil {
		movie.Director = *input.Director
	}
	if input.Actors != nil {
		movie.Actors = *input.Actors
	}
	if input.Plot != nil {
		movie.Plot = *input.Plot
	}
	if input.Language != nil {
		movie.Language = *input.Language
	}
	if input.Country != nil {
		movie.Country = *input.Country
	}
	if input.Awards != nil {
		movie.Awards = *input.Awards
	}

	// Step 4: Validate full movie
	v := validator.New()
	if valid := validator.ValidateMovieInput(v, *movie); !valid.Valid() {
		app.logger.Println("Validation error:", valid.Errors)
		app.errorResponse(w, r, http.StatusUnprocessableEntity, valid.Errors)
		return
	}

	// Step 5: Update movie
	err = app.models.Movies.Update(movie)
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

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Title    string
		Genres   []string
		Page     int
		PageSize int
		Sort     string
	}

	qs := r.URL.Query()

	v := validator.New()

	input.Title = app.readString(qs, "title", "")
	input.Genres, _ = app.readCSV(qs, "genres", nil)
	input.Page = app.readInt(qs, "page", 1, v)
	input.PageSize = app.readInt(qs, "page_size", 10, v)
	input.Sort = app.readString(qs, "sort", "id")

	if !v.Valid() {
		app.logger.Println("Validation error:", v.Errors)
		app.errorResponse(w, r, http.StatusUnprocessableEntity, v.Errors)
		return
	}

	movies, err := app.models.Movies.GetAllMovieWithQuery(input.Title, input.Genres, data.Filters{
		Page:     input.Page,
		PageSize: input.PageSize,
		Sort:     input.Sort,
	})
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJson(w, http.StatusOK, movies, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	app.logger.Println("Successfully fetched all movies with query parameters")
	app.logger.Printf("Title: %s, Genres: %v, Page: %d, PageSize: %d, Sort: %s", input.Title, input.Genres, input.Page, input.PageSize, input.Sort)
	app.logger.Println("Successfully fetched all movies with query parameters")

}
