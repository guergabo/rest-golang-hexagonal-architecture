package dto

import (
	"banking-app/errs"
	"strings"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

// kept here even though using it in the service because it is
// a logical unit with NewAccounRequest since it is validating
// NewAccountRequest
// capital to export to have it used by service layer
func (r NewAccountRequest) Validate() *errs.AppError {
	// validate the ammount - want a minimum of 5,000
	// to open account
	if r.Amount < 5000 {
		return errs.NewValidationError("To open a new account you need to deposit atleast $5000.00")
	}

	// check account type
	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		return errs.NewValidationError("Account type should be checking or saving")
	}

	return nil
}
