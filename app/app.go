package app

import (
	"log"
	"net/http"

	"github.com/ashishjuyal/banking/domain"
	"github.com/ashishjuyal/banking/service"
	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter() // mux is gorilla Mutlipexer which is a router which simplifies route definition and also popular for its routing capabilities

	//wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.CustomerRepositoryDb)}

	// defining routes( multiplexer)
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))

	/*router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)
	*/

	// starting server
}
