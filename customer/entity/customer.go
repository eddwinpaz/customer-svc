package entity

import (
	"github.com/eddwinpaz/customer-svc/customer/utils"
	"github.com/google/uuid"
)

type ContextKey string

const ContextCustomerKey ContextKey = "customer"

type Customer struct {
	CustomerID    string `json:"customer_uuid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
	CreatedOn string `json:"created_on,omitempty"`
	LastLogin string `json:"last_login,omitempty"`
}

// encryptPassword encrypt string password to sha1 encode
func (customer *Customer) EncryptPassword() {
	customer.Password = utils.EncryptPassword(customer.Password)
}

func (customer *Customer) GenerateUUID() {
	uuid, _ := uuid.NewRandom()
	customer.CustomerID = uuid.String()
}

// isValidUUID to prevent toxic data entering.
// func (customer Customer) IsValidUUID() bool {
// 	_, err := uuid.Parse(customer.CustomerID)
// 	return err == nil
// }

func (customer *Customer) Public() *Customer {
	return &Customer{
		CustomerID:    customer.CustomerID,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
	}
}
