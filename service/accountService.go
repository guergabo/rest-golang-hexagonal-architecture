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
	MakeTransaction(dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
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

func (s DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	// incoming request validation - 1
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	// server side validation for checking the available balance in the account - 2
	if req.IsTransactionTypeWithdrawal() {
		account, err := s.repo.FindBy(req.AccountId)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdraw(req.Amount) {
			return nil, errs.NewValidationError(("Insufficiet balance in the amount"))
		}
	}

	// success
	t := domain.Transaction{
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}
	transaction, appError := s.repo.SaveTransaction(t)
	if appError != nil {
		return nil, appError
	}
	response := transaction.ToDto()
	return &response, nil
}
