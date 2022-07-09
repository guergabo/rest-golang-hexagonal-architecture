// Implements Service Interface
package service

import (
	"banking-app/domain"
	"banking-app/errs"
)

// primary port - service interface
// has depenedency on repository interface not concrete implementation
type CustomerService interface {
	GetAllCustomers(string) ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
}

// implementation of the interface
type DefaultCustomerService struct {
	// interface not concrete implementation
	// loosely coupled anything that meets the interfaces standards can pass
	// good for testing
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, *errs.AppError) {
	// depenedent on an interface that implements FindAll()
	switch status { // implicit break unlike C++ and Java
	case "active":
		status = "1"
	case "inactive":
		status = "0"
	default:
		status = ""
	}

	return s.repo.FindAll(status)
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.ById(id)
}

// need a constructor
// takes depenedency of the repo so it can be injected in default customer service
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
