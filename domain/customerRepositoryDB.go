package domain

// note the underscore
import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
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
			log.Println("Error whil scanning customres " + err.Error())
		}
		customers = append(customers, c)
	}
	return customers, nil
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
