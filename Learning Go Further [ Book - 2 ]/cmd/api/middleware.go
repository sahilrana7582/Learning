package main

import (
	"fmt"
	"net/http"
	"strings"
)

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {

			if err := recover(); err != nil {
				app.logger.Println("Panic recovered:", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				fmt.Println("Request failed due to panic:", err)
				return
			}

		}()
		next.ServeHTTP(w, r)
	})
}

func (app *application) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Vary", "Authorization")

		authorizationHeader := r.Header.Get("Authorization")

		if authorizationHeader == "" {
			app.badRequestResponse(w, r, fmt.Errorf("missing or invalid Authorization header"))
			return
		}

		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			app.badRequestResponse(w, r, fmt.Errorf("invalid Authorization header format"))
			return
		}

		token := headerParts[1]
		if token == "" {
			app.badRequestResponse(w, r, fmt.Errorf("missing token in Authorization header"))
			return
		}

		if token != "ABC" {
			app.invalidCredentialsResponse(w, r)
			return
		}
		fmt.Println("Token is valid, proceeding with request")
		next.ServeHTTP(w, r)
	})
}
