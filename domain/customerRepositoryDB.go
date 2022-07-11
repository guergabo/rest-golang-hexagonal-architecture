// Implements Repository Interface (protocol)
package domain

// note the underscore
import (
	"banking-app/errs"
	"banking-app/logger"
	"database/sql" // must be used in conjuction with a database driver
	"time"

	_ "github.com/go-sql-driver/mysql" // the actual driver that implements the interface
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

// golang database driver used to enable interaction with mysql server, go-sql-driver
func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {

	var findAllSql string
	// var rows *sql.Rows
	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSql = "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		// rows, err = d.client.Query(findAllSql)
		err = d.client.Select(&customers, findAllSql)
	} else {
		findAllSql = "select customer_id, name, city, zipcode, date_of_birth, status from customers where status=?"
		// rows, err = d.client.Query(findAllSql, status)
		// querying and marshalling into 1 single function
		err = d.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	// if no error loop, for sql rows, sqlx helps us marshall rows into structs
	// to help us get rid of the boiler plate code below where we have to loop
	// over the rows every time to get into struct format and append to
	// slice of Customer
	// picky matches name by name can fix with db tags, doesn't care about capital though
	// err = sqlx.StructScan(rows, &customers)
	// if err != nil {
	// 	logger.Error("Error while scanning customers " + err.Error())
	// 	return nil, errs.NewUnexpectedError("Unexpected database error")
	// }
	/*
		for rows.Next() {
			// giving destination for scan to write to
			var c Customer
			err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
			if err != nil {
				logger.Error("Error while scanning customers " + err.Error())
			}
			customers = append(customers, c)
		}
	*/
	return customers, nil
}

func (d CustomerRepositoryDB) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id=?"

	// at most one customer with id
	// row := d.client.QueryRow(customerSql, id)
	var c Customer

	err := d.client.Get(&c, customerSql, id)

	// err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		// customer does not exist - 404 not found, user error no need to log, developer does not care, proper functionality
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}
		// database is down - 500 internal server error, we want to log
		logger.Error("Error while scanning customer " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return &c, nil
}

// ask to maintain the connection
func NewCustomerRepositoryDB() CustomerRepositoryDB {
	// connection to mysql database server
	// db = client of mysql
	// dbClient, err := sql.Open("mysql", "root:Popeye101!@/banking")
	dbClient, err := sqlx.Open("mysql", "root:Popeye101!@/banking")
	if err != nil {
		panic(err)
	}

	dbClient.SetConnMaxLifetime(time.Minute * 3)
	dbClient.SetMaxOpenConns(10)
	dbClient.SetMaxIdleConns(10)

	return CustomerRepositoryDB{dbClient}
}
