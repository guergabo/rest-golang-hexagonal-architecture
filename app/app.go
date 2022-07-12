package app

import (
	"banking-app/domain"
	"banking-app/service"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// if environment variable is missing we should exit, check all the variables, print out which one etc.
func sanityCheck() {
	// server
	env := []string{"SERVER_ADDRESS", "SERVER_PORT", "DB_USER", "DB_PASSWD", "DB_ADDR", "DB_PORT", "DB_NAME"}

	for _, v := range env {
		if os.Getenv(v) == "" {
			log.Fatalf("Environment variable %s not defined", v)
		}
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
	dbClient := getDBClient()
	customerRepositoryDb := domain.NewCustomerRepositoryDB(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDB(dbClient)

	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDb)}

	// registers handler and pattern with default multiplexer
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}/account", ah.newAccount).Methods(http.MethodPost) // creat new account for existing customer, real life would have to go right before
	// transaction is related to the account so should use account handler
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).Methods(http.MethodPost) // creat new account for existing customer, real life would have to go right before

	// starts the server, relying on default router
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

func getDBClient() *sqlx.DB {
	// connection to mysql database server
	// db = client of mysql
	// dbClient, err := sql.Open("mysql", "root:Popeye101!@/banking")
	// get rid of hardcoded stuff for production environment
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbAddr, dbPort, dbName)

	dbClient, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	dbClient.SetConnMaxLifetime(time.Minute * 3)
	dbClient.SetMaxOpenConns(10)
	dbClient.SetMaxIdleConns(10)
	return dbClient
}
