package domain

import (
	"banking-app/dto"
	"banking-app/errs"
)

// domain object
type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{a.AccountId}
}

// depends on second port
type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
