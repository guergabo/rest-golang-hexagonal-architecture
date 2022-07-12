package app

import (
	"banking-app/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// part of the wiring
type CustomerHandlers struct {
	// to interface not concrete
	service service.CustomerService
}

// http.ResponseWriter goes back to the client
// handler should have a dependency of the service
// 500 or 200, can't be a 404 cause its all
func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status") // returns empty string if no status key
	customers, err := ch.service.GetAllCustomers(status)

	if err != nil { // internal server error
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	// success
	writeResponse(w, http.StatusOK, customers)
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		// 404 not found, not dynamic, but don't want to match content and send status code because of future change,
		// it will break existing code, instead of that we should work on the ids
		// now not match any if conditions, separate the error logic outside of here
		// must be in this order to have content-type applied
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	writeResponse(w, http.StatusOK, customer)
}

// responsible for writing response
func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	// encode returns an error
	err := json.NewEncoder(w).Encode(data)
	if err != nil { // something went really wrong encoding
		panic(err)
	}
}
