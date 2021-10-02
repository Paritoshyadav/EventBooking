package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

func server(endpoint string) error {

	r := mux.NewRouter()
	handler := &EventServiceHandle{}
	eventRouter := r.PathPrefix("/events").Subrouter()
	eventRouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.findEventHandler)
	eventRouter.Methods("GET").Path("").HandlerFunc(handler.allEventHandler)
	eventRouter.Methods("POST").Path("").HandlerFunc(handler.newEventHandler)

	return http.ListenAndServe(endpoint, r)
}
