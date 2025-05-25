package main

import (
	"fmt"
	"net/http"
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
