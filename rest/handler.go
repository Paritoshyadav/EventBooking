package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/eventbooking/lib/persistence"
	"github.com/gorilla/mux"
)

type EventServiceHandle struct {
	dbhandler persistence.DatabaseHandler
}

func newEventHandler(databasehandler persistence.DatabaseHandler) *EventServiceHandle {
	return &EventServiceHandle{
		dbhandler: databasehandler,
	}
}

func (eh *EventServiceHandle) findEventHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	SearchCriteria, ok := vars["SearchCriteria"]
	if !ok {
		rw.WriteHeader(400)
		fmt.Fprintf(rw, `{error: No search criteria found, you can either search by id via /id/4
			to search by name via /name/coldplayconcert}`)
		return
	}
	SearchKey, ok := vars["search"]
	if !ok {
		rw.WriteHeader(400)
		fmt.Fprintf(rw, `{error: No search key found, you can either search by id via /id/4
			to search by name via /name/coldplayconcert}`)
		return
	}
	var event persistence.Event
	var err error

	switch strings.ToLower(SearchCriteria) {
	case "name":
		event, err = eh.dbhandler.FindEventByName(SearchKey)

	case "id":

		event, err = eh.dbhandler.FindEvent(SearchKey)

	}

	if err != nil {
		fmt.Fprintf(rw, "{error %s}", err)
		return
	}
	rw.Header().Set("Content-Type", "application/json;charset=utf8")
	err = json.NewEncoder(rw).Encode(&event)
	if err != nil {
		rw.WriteHeader(400)
		fmt.Fprintf(rw, "{error: Error occured while trying encode events to JSON %s}", err)
	}
}
func (eh *EventServiceHandle) allEventHandler(rw http.ResponseWriter, r *http.Request) {
	events, err := eh.dbhandler.FindAllAvailableEvents()
	fmt.Printf("%v", events)
	if err != nil {
		rw.WriteHeader(400)
		fmt.Fprintf(rw, "error occur while searching %s", err)
	}
	rw.Header().Set("Content-Type", "application/json;charset=utf8")
	err = json.NewEncoder(rw).Encode(&events)
	if err != nil {
		rw.WriteHeader(400)
		fmt.Fprintf(rw, "{error: Error occured while trying encode events to JSON %s}", err)
	}
}

func (eh *EventServiceHandle) newEventHandler(rw http.ResponseWriter, r *http.Request) {
	var event persistence.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	fmt.Println(event)
	if err != nil {
		rw.WriteHeader(400)
		fmt.Fprintf(rw, "{error: Error occured while trying decode events from JSON request %s}", err)
	}
	newEventId, err := eh.dbhandler.AddEvent(event)
	if err != nil {
		rw.WriteHeader(500)
		fmt.Fprintf(rw, "{error: Error occured while adding event %s}", err)
	}
	err = json.NewEncoder(rw).Encode(newEventId)
	if err != nil {
		rw.WriteHeader(500)
		fmt.Fprintf(rw, "{error: Error occured while trying encode events to JSON %s}", err)
	}

}
