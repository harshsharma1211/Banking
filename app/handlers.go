package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/ashishjuyal/banking/service"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}
type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	//customers := []Customer{
	//	{Name: "Anukul", City: "Greater Noida", Zipcode: "110075"},
	//	{Name: "Raj", City: "Gorakhpur", Zipcode: "110075"},
	//	{Name: "Sandeep", City: "Noida", Zipcode: "110075"},
	//}
	customers, _ := ch.service.GetAllCustomer()

	//w.Header().Add("Content-Type", "application/json")
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)

	} else {

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)

	}
}
