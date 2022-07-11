package domain

import (
	"banking-app/dto"
	"banking-app/errs"
)

// domain object - mapped with server side layer (database)
// business logic - what is a customer, domain object
// maps to database model
type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

func (c Customer) statusAsText() string {
	// 1 and 0 is internal statement, change to active or inactive
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}

	return statusAsText
}

// transferred responsibility of creating dto to the domain
func (c Customer) ToDto() dto.CustomerResponse {
	// 1 and 0 is internal statement, change to active or inactive
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(),
	}
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
