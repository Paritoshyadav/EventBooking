package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type EventServiceHandle struct {
}

func (eh *EventServiceHandle) findEventHandler(w http.ResponseWriter, r *http.Request) {

}
func (eh *EventServiceHandle) allEventHandler(w http.ResponseWriter, r *http.Request) {

}
func (eh *EventServiceHandle) newEventHandler(w http.ResponseWriter, r *http.Request) {

}

func server(endpoint string) error {

	r := mux.NewRouter()
	handler := &EventServiceHandle{}
	eventRouter := r.PathPrefix("/events").Subrouter()
	eventRouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.findEventHandler)
	eventRouter.Methods("GET").Path("").HandlerFunc(handler.allEventHandler)
	eventRouter.Methods("POST").Path("").HandlerFunc(handler.newEventHandler)

	return http.ListenAndServe(endpoint, r)
}

func main() {

}
