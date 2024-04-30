package app

import (
	"banking/config"
	"banking/logger"
	//"banking/domain"
	db "banking/domain"
	"banking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start((config *config.AppConfig)  {
	router := mux.NewRouter()
	dbClient, err := db.Init(config.DB)
	if err != nil {
		logger.Fatal("Database error: " + err.Error())
	}
	//ch := CustomerHandlers{ service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{ service.NewCustomerService(dbClient, config.DB.Database, config.DB.UserCollection)}
	router.HandleFunc("/customers", ch.GetAllCustomer).Methods(http.MethodGet)
	//router.HandleFunc("/customers", CreateCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.GetCustomer).Methods(http.MethodGet)
	//starting server
	log.Fatal(http.ListenAndServe("localhost: 8000",router))

}

