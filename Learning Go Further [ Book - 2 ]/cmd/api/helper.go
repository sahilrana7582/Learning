package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/sahilrana7582/Learning/internal/validator"
)

func (app *application) readIdParamFromRequest(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil {
		app.logger.Println("Invalid movie ID")
		return 0, err
	}
	return id, nil
}

func (app *application) writeJson(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		app.logger.Println("Error marshalling JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return err
	}

	for key, value := range headers {
		w.Header()[key] = value
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

func (app *application) readString(qs url.Values, key string, defaultValue string) string {

	value := qs.Get(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func (app *application) readInt(qs url.Values, key string, defaultValue int, v *validator.Validator) int {
	value := qs.Get(key)
	if value == "" {
		return defaultValue
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		v.AddError(key, "must be a valid integer")
		app.logger.Println("Error converting string to int:", err)
		return defaultValue
	}
	return intValue
}

func (app *application) readCSV(qs url.Values, key string, defaultValue []string) ([]string, error) {
	value := qs.Get(key)
	if value == "" {
		return defaultValue, nil
	}

	csvValues := strings.Split(value, ",")
	return csvValues, nil
}
