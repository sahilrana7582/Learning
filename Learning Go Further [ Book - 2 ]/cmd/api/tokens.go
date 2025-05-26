package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sahilrana7582/Learning/internal/validator"
)

func (app *application) createAuthTokenHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	if v := validator.ValidateCredentials(v, input.Email, input.Password); !v.Valid() {
		app.errorResponse(w, r, http.StatusUnprocessableEntity, v.Errors)
		return
	}

	user, err := app.models.User.GetByEmail(input.Email)
	if err != nil {
		fmt.Println("Error fetching user by email:", err)
		app.serverErrorResponse(w, r, err)
		return
	}

	if user == nil {
		app.notFoundResponse(w, r)
		return
	}

	match, err := user.Password.Compare(input.Password)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if !match {
		app.errorResponse(w, r, http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
		return
	}

	err = app.writeJson(w, http.StatusOK, map[string]string{
		"status":  "success",
		"message": "Authentication successful",
		"token":   "Token Info",
	}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

}
