package app

import (
	"banking/service"
	"encoding/json"
	"encoding/xml"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers)GetAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		err := xml.NewEncoder(w).Encode(customers)
		if err != nil {
			return
		}
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func (ch *CustomerHandlers) GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	customerId, _ := strconv.Atoi(id)
	customer, err := ch.service.GetCustomer(customerId)
	if err != nil{
		WriteResponse(w, err.Code, err.AsMessage())
	} else {
		WriteResponse(w, http.StatusOK, customer)
	}
}

func WriteResponse(w http.ResponseWriter, code int, data interface{})  {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil{
		panic(err)
	}
}