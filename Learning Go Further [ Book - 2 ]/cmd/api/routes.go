package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *applicaton) routes() *httprouter.Router {

	// New Router Instance
	router := httprouter.New()

	// Add All The Routes Here and Its Handlers

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.getMovieById)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createNewMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies", app.getMovieHandler)

	return router
}
