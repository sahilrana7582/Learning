package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sahilrana7582/Learning/internal/data"
	"github.com/sahilrana7582/Learning/internal/validator"
)

type registerUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var input registerUserInput
	var userData data.User

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	if v := validator.ValidateUserInput(v, input.Name, &input.Email, &input.Password); !v.Valid() {
		app.errorResponse(w, r, http.StatusUnprocessableEntity, v.Errors)
		return
	}

	userData = data.User{
		Name:      input.Name,
		Email:     input.Email,
		Password:  input.Password,
		Activated: true,
	}

	err = app.models.User.Insert(&userData)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	resp := map[string]interface{}{
		"status":  "success",
		"message": "User registered successfully",
		"user_id": userData.ID,
		"email":   userData.Email,
	}

	go func() {

		defer func() {
			if r := recover(); r != nil {
				app.logger.Println("Panic recovered in welcome email goroutine:", r)
			}
		}()

		err = app.mailer.Send(
			userData.Email,
			"user_welcome.tmpl",
			userData,
		)
		if err != nil {
			app.logger.Println("Error sending welcome email:", err)
			app.errorResponse(w, r, http.StatusInternalServerError, "Failed to send welcome email")
			return
		}

	}()

	err = app.writeJson(w, http.StatusCreated, resp, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

type userByEmail struct {
	Email string `json:"email"`
}

func (app *application) getUserByEmailHandler(w http.ResponseWriter, r *http.Request) {
	var input userByEmail
	fmt.Println("getUserByEmailHandler called")

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, "Invalid request payload")
		return
	}

	v := validator.New()
	validator.ValidateEmail(&input.Email)
	if !v.Valid() {
		app.errorResponse(w, r, http.StatusUnprocessableEntity, v.Errors)
		return
	}

	userData, err := app.models.User.GetByEmail(input.Email)
	if err != nil {
		app.errorResponse(w, r, http.StatusNotFound, "User not found")
		return
	}

	err = app.writeJson(w, http.StatusOK, userData, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

}
