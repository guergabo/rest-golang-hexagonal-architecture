package app

import (
	"banking-app/domain"
	"banking-app/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// describes the application endpoints and starts app
func Start() {

	// define our own multiplexer
	router := mux.NewRouter()

	// wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// registers handler and pattern with default multiplexer
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// starts the server, relying on default router
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
