package domain

import (
	"banking-app/errs"
	"banking-app/logger"
	"strconv"

	"github.com/jmoiron/sqlx"
)

// will implement interface and hold sqlx.DB in order
// to make connection
type AccountRepositoryDB struct {
	client *sqlx.DB
}

func (d AccountRepositoryDB) Save(a Account) (*Account, *errs.AppError) {
	// pass by value cause don't want to change what's passed
	// to me but could, want to return account struct with update id
	// account_id is the primary key
	sqlInsert := "INSERT INTO accounts(customer_id, opening_date, account_type, amount, status) VALUES(?, ?, ?, ?, ?)"

	// takes query and arguments
	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	// success
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	// updating copy and sending back
	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepositoryDB(dbClient *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{dbClient}
}
