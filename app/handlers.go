package app

import (
	"banking-app/service"
	"encoding/json"
	"net/http"
)

type Customer struct {
	// json tags
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
}

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
