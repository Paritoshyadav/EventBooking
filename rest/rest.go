package rest

import (
	"net/http"

	"github.com/eventbooking/lib/persistence"
	"github.com/gorilla/mux"
)

func Server(endpoint string, dbhandler persistence.DatabaseHandler) error {

	r := mux.NewRouter()
	handler := newEventHandler(dbhandler)
	eventRouter := r.PathPrefix("/events").Subrouter()
	eventRouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.findEventHandler)
	eventRouter.Methods("GET").Path("").HandlerFunc(handler.allEventHandler)
	eventRouter.Methods("POST").Path("").HandlerFunc(handler.newEventHandler)

	return http.ListenAndServe(endpoint, r)
}
