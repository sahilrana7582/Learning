package main

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *applicaton) readIdParamFromRequest(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil {
		app.logger.Println("Invalid movie ID")
		return 0, err
	}
	return id, nil
}
