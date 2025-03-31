package service

import domain "github.com/harshsharma1211/Domain"

//Defining Primary Port
type CustomerService interface {
	GetALLCustomer() ([]domain.Customer, error)
}

//Service Implementation
type DefaultCustomerService struct {
	//dependency of the customerRepository from domain package
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetALLCustomer() ([]domain.Customer, error) {
	return s.repo.FindALL()
}

//Helper Function to initiate the defaultcustomer service
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
