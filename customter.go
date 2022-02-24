package domain

import "context"

type Customer struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	MobileNumber string `json:"mobile_number"`
	Email     string `json:"email`
	Address      string `json:"address"`
	Education   string `json:"education"`
}


// CustomerRepository 
type CustomerRepository interface {
	Fetch(context.Context) ([]Customer, error)
	GetByID(context.Context, int64) (Customer, error)
	Update(context.Context, int64, map[string]interface{}) error
	Store(context.Context, *Customer) error
	Delete(context.Context, int64) error
}
