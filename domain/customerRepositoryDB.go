// CustomerRepository implementation
package domain

// note the underscore
import (
	"banking-app/errs"
	"database/sql" // must be used in conjuction with a database driver
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" // the actual driver that implements the interface
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

// golang database driver used to enable interaction with mysql server, go-sql-driver
func (d CustomerRepositoryDB) FindAll() ([]Customer, error) {
	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
	rows, err := d.client.Query(findAllSql)
	if err != nil {
		log.Println("Error while querying customer table " + err.Error())
		return nil, err
	}

	// if no error loop, for sql rows
	customers := make([]Customer, 0)
	for rows.Next() {
		// giving destination for scan to write to
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			log.Println("Error while scanning customers " + err.Error())
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRepositoryDB) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id=?"

	// at most one customer with id
	row := d.client.QueryRow(customerSql, id)

	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		// customer does not exist - 404 not found
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}
		// database is down - 500 internal server error
		log.Println("Error while scanning customer " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return &c, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	// connection to mysql database server
	// db = client of mysql
	dbClient, err := sql.Open("mysql", "root:Popeye101!@/banking")
	if err != nil {
		panic(err)
	}

	dbClient.SetConnMaxLifetime(time.Minute * 3)
	dbClient.SetMaxOpenConns(10)
	dbClient.SetMaxIdleConns(10)

	return CustomerRepositoryDB{dbClient}
}
