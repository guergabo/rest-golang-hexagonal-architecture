package domain

// stub adapter, mocking the database of customers
// should implement FindAll() function
type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

// responsible for creating our dummy customers, constructor
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "000", Name: "Gabriel", City: "Margate", Zipcode: "33063", DateOfBirth: "07/29/1997", Status: "1"},
		{Id: "001", Name: "Vikitha", City: "Hyderabad", Zipcode: "07311", DateOfBirth: "02/04/1998", Status: "1"},
	}

	return CustomerRepositoryStub{customers}
}
