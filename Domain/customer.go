package domain

import "github.com/ashishjuyal/banking-lib/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

// Secondary Port or Repository Interface
type CustomerRepository interface {
	FindALL() ([]Customer, error)
	ById(string) (*Customer, *errs.AppError)
}
