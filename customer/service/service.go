package service

import (
	"github.com/eddwinpaz/customer-svc/customer/entity"
	"github.com/eddwinpaz/customer-svc/customer/repository"
)

type RepositoryImpl struct {
	GetRespositoryCustomer repository.RepositoryInterface
}

func (rep RepositoryImpl) HealthCheck() error {
	return rep.GetRespositoryCustomer.HealthCheck()
}

func (rep RepositoryImpl) AuthenticateCustomer(email string, password string) (*entity.Customer, error) {
	return rep.GetRespositoryCustomer.AuthenticateCustomer(email, password)
}

func (rep RepositoryImpl) UpdateCustomerByID(id string, customer entity.Customer) bool {
	return rep.GetRespositoryCustomer.UpdateCustomerByID(id, customer)
}

func (rep RepositoryImpl) SaveCustomer(customer entity.Customer) bool {
	return rep.GetRespositoryCustomer.SaveCustomer(customer)
}

func (rep RepositoryImpl) GetCustomerByID(id string) (*entity.Customer, error) {
	return rep.GetRespositoryCustomer.GetCustomerByID(id)
}

func (rep RepositoryImpl) DeleteCustomerByID(id string) bool {
	return rep.GetRespositoryCustomer.DeleteCustomerByID(id)
}
