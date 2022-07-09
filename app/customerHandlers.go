package app

import (
	"banking-app/service"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type CustomerHandlers struct {
	// to interface not concrete
	service service.CustomerService
}

// http.ResponseWriter goes back to the client
// handler should have a dependency of the service
func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomers()
	// Response header
	w.Header().Add("Content-Type", "application/json")
	// Marshshaling data structures to JSON representation
	json.NewEncoder(w).Encode(customers)
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	customer, err := ch.service.GetCustomer(id)
	fmt.Println(customer)
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
	if err != nil { // something went really wrong
		panic(err)
	}
}
