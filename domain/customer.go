package domain

import (
	"banking/dto"
	"banking/errs"
)

type Customer struct {
	Id          int
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.Status,
	}
}

// we introduced secondary port here------->1st step
type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ById(customerId int) (*Customer, *errs.AppError)
}
