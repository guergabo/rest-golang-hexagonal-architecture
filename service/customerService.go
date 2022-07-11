// Implements Service Interface
// this is the business logic
package service

import (
	"banking-app/domain"
	"banking-app/dto"
	"banking-app/errs"
)

// primary port - service interface
// has depenedency on repository interface not concrete implementation
type CustomerService interface {
	GetAllCustomers(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

// implementation of the interface
type DefaultCustomerService struct {
	// interface not concrete implementation
	// loosely coupled anything that meets the interfaces standards can pass
	// good for testing
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {
	// depenedent on an interface that implements FindAll()
	switch status { // implicit break unlike C++ and Java
	case "active":
		status = "1"
	case "inactive":
		status = "0"
	default:
		status = ""
	}

	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	// can append to empty - A nil map is equivalent to an empty map except that elements can’t be added.
	response := make([]dto.CustomerResponse, 0)
	for _, c := range customers {
		response = append(response, c.ToDto())
	}

	return response, err
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	// transform domain object to data transfer object
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	response := c.ToDto()

	return &response, nil
}

// need a constructor
// takes depenedency of the repo so it can be injected in default customer service
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
