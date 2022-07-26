package repository

import "github.com/eddwinpaz/customer-svc/customer/entity"

// Repository interface to repository
type RepositoryInterface interface {
	HealthCheck() error
	GetCustomerByID(id string) (*entity.Customer, error)
	UpdateCustomerByID(customer_id string, customer entity.Customer) bool
	SaveCustomer(customer entity.Customer) bool
	DeleteCustomerByID(customer_id string) bool
	AuthenticateCustomer(email string, password string) (*entity.Customer, error)
}
