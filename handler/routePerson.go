package handler

import (
	"net/http"

	"github.com/MatiasCeb/api-go/tree/main/model/middleware"
)

func RoutePerson(mux *http.ServeMux, storage Storage) {
	h := newPerson(storage)

	mux.HandleFunc("/v1/persons/create", middleware.Log(h.create))
	mux.HandleFunc("/v1/persons/update", h.update)
	mux.HandleFunc("/v1/persons/delete", middleware.Log(h.delete))
	mux.HandleFunc("/v1/persons/get-by-id", h.getByID)
	mux.HandleFunc("/v1/persons/get-all", middleware.Log(h.getAll))
}
