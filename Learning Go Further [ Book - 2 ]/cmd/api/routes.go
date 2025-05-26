package main

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {

	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.getMovieById)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createNewMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies", app.getMovieHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/movies/:id", app.updateMovieHandler)
	router.HandlerFunc(http.MethodGet, "/lala", app.getAllMovies)
	router.HandlerFunc(http.MethodDelete, "/v1/movies/:id", app.deleteMovieHandler)
	router.HandlerFunc(http.MethodGet, "/test", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Second)
		w.Write([]byte("Test endpoint reached"))
	},
	)

	// User registration route
	router.HandlerFunc(http.MethodPost, "/v1/users/register", app.registerUserHandler)
	router.HandlerFunc(http.MethodPost, "/v1/users", app.getUserByEmailHandler)
	// Authentication route
	router.HandlerFunc(http.MethodPost, "/v1/users/auth", app.createAuthTokenHandler)

	return router
}
