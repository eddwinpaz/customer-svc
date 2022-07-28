package service

import "github.com/eddwinpaz/customer-svc/customer/entity"

type ServiceInterface interface {
	HealthCheck() error
	GetCustomerByID(customer_uuid string) (*entity.Customer, error)
	UpdateCustomerByID(customer entity.Customer) bool
	SaveCustomer(customer entity.Customer) bool
	DeleteCustomerByID(customer_uuid string) bool
	AuthenticateCustomer(email string, password string) (*entity.Customer, error)
}
