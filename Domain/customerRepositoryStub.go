package domain

//Defining Stub Adapter
type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindALL() ([]Customer, error) {
	return s.customers, nil
}

// helper function:This function is resposible for creating new customer repository stuff
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{

		{Id: "1001", Name: "Ashish", City: "New Delhi", Zipcode: "110011", DateofBirth: "2000-01-01", Status: "1"},

		{Id: "1003", Name: "Rob", City: "New Delhi", Zipcode: "110011", DateofBirth: "2000-01-01", Status: "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}
