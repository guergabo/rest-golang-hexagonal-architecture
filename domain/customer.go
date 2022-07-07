package domain

// business logic - what is a customer
type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

// primary port - service interface

// secondary port - repository interface
// help us find all the customers from the server side
// defining the port is like defining a protocol, any component
// which follows the guideliens of this protocol should be able
// to connect to this port
type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
