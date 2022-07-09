package app

import (
	"banking-app/domain"
	"banking-app/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// acts like an adapter accessing the service port, wiring together
// describes the application endpoints and starts app
func Start() {

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
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
