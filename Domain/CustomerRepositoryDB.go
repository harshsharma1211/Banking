package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client 

// receiver function to implement the FindAll method in CustomerRepositoryDb struct
func (d CustomerRepositoryDb) FindALL ([]Customer, error) {

	//Defining the sql
	FindALLSql := "select customer_id,name,city,zipcode,date_of_birth,status from customers"

	//Quering the Database
	rows, err := d.client.Query(FindALLSql)
	if err != nil {
		log.Println("Error while quering customer table " + err.Error())
		return nil, err
	}

	//looping over all the rows we received from the Database
	for rows.Next() {
		var c Customer
		rows.Scan(&c.Id, &c.name, &c.City, &c.Zipcode, &c.DateofBirth, &c.DateofBirth, &c.Status)
		if err != nil {
			log.Println("Error while quering customer table " + err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}

	return customers, nil

}

func (d CustomerRepositoryDb) ById(id string) (*Customer,error){
	customerSql := "select customer_id,name,city,zipcode,date_of_birth,status from customers where customer_id =?"
    row := d.client.QueryRow(customerSql,id)
	var c customer
	row.Scan()

}

func NewCustomerRepositoryDb() CustomerRepositoryDb {

	client, err := sql.Open("mysql", "root:Harsh1211(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{}

}

}