package service

import "banking-app/domain"

// has depenedency on repository interface not concrete implementation
type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
}

// implementation of the interface
type DefaultCustomerService struct {
	// interface not concrete implementation
	// loosely coupled anything that meets the interfaces standards can pass
	// good for testing
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	// depenedent on an interface that implements FindAll()
	return s.repo.FindAll()
}

// need a constructor
// takes depenedency of the repo so it can be injected in default customer service
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
