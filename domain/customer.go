package domain

import "banking-app/errs"

// business logic - what is a customer, domain object
type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

// primary port - service interface in the service package

// secondary port - repository interface
// help us find all the customers from the server side
// defining the port is like defining a protocol, any component
// which follows the guideliens of this protocol should be able
// to connect to this port
type CustomerRepository interface {
	// status "1", "0", "" (all customers)
	FindAll(status string) ([]Customer, *errs.AppError) // use method and pass stauts
	ById(string) (*Customer, *errs.AppError)            // returning a pointer in case there is no customer and we want to response with nil
}
