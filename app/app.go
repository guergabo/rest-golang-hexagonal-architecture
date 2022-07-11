package app

import (
	"banking-app/domain"
	"banking-app/service"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// if environment variable is missing we should exit, check all the variables, print out which one etc.
func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variable not defined...")
	}
}

// acts like an adapter accessing the service port, wiring together
// describes the application endpoints and starts app
func Start() {

	// sanity check
	sanityCheck()

	// define our own multiplexer
	router := mux.NewRouter()

	// wiring
	// stub adapter - conforms to port (protocol, interface)
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	// database adapter - conforms to port (protocol, interface)
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	// registers handler and pattern with default multiplexer
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// starts the server, relying on default router
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
