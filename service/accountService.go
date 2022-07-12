package service

import (
	"banking-app/domain"
	"banking-app/dto"
	"banking-app/errs"
	"time"
)

// primary port
type AccountService interface {
	// towards domain
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	// reference to the secondary port
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	// validate request before doing any kind of processing
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	a := domain.Account{
		AccountId:   "", // don't have this will be given to use after database insert operation
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"), // format to be accepted by database
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	// pass dto to the repository
	// service layer does transformation of dto to domain and back
	newAccount, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}

	// success - need to transform domain object into new account response (dto)
	response := newAccount.ToNewAccountResponseDto()
	return &response, nil
}

// need to inject account repository
func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}
