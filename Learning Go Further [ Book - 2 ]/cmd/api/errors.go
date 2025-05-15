package main

import "net/http"

func (app *applicaton) logError(r *http.Request, err error) {
	app.logger.Println(err)
}

func (app *applicaton) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := map[string]interface{}{
		"error": message,
	}
	err := app.writeJson(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

func (app *applicaton) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotFound, "Resource not found")
}

func (app *applicaton) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (app *applicaton) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
}

func (app *applicaton) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusMethodNotAllowed, "Method not allowed")
}
func (app *applicaton) invalidCredentialsResponse(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusUnauthorized, "Invalid credentials")
}
