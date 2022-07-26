package controller

import (
	"encoding/json"
	"net/http"

	"github.com/eddwinpaz/customer-svc/customer/entity"
	"github.com/eddwinpaz/customer-svc/customer/service"
	"github.com/gorilla/mux"
)

type ServiceImpl struct {
	GetServiceCustomer service.ServiceInterface
}

func Response(status bool, desc string, data interface{}, w http.ResponseWriter, httpStatus int) {

	// Return Successfull Response
	var response = entity.Response{
		Status:      status,
		Description: desc,
		Data:        data,
	}

	// Return Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	_ = json.NewEncoder(w).Encode(response)
}

func (svc ServiceImpl) HealthCheck(w http.ResponseWriter, r *http.Request) {

	err := svc.GetServiceCustomer.HealthCheck()
	// Return Error Response
	if err != nil {
		Response(false, "Database unaccessible", nil, w, http.StatusNotFound)
		return
	}
	// Return Successfull
	Response(true, "UP", nil, w, http.StatusOK)
}

func (svc ServiceImpl) CustomerAuthentication(w http.ResponseWriter, r *http.Request) {

	var credentials entity.Credentials

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		Response(false, entity.ErrInvalidJSON.Error(), nil, w, http.StatusBadRequest)
		return
	}
	credentials.EncryptPassword()
	customer, err := svc.GetServiceCustomer.AuthenticateCustomer(credentials.Email, credentials.Password)

	if err != nil {
		Response(false, entity.ErrInvalidCredentials.Error(), nil, w, http.StatusBadRequest)
		return
	}

	resp, err := credentials.GenerateJwtToken(customer)

	if err != nil {
		Response(false, entity.ErrInternalServerError.Error(), nil, w, http.StatusInternalServerError)
		return
	}

	Response(true, "Customer logged in successfully", resp, w, http.StatusOK)
}

func (svc ServiceImpl) UpdateCustomerByID(w http.ResponseWriter, r *http.Request) {

	var customer entity.Customer

	// Get Customer ID
	params := mux.Vars(r)
	customer_id := params["id"]

	// validate param ID
	if customer_id == "" {
		Response(false, entity.ErrBadParamInput.Error(), nil, w, http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		Response(false, entity.ErrInvalidJSON.Error(), nil, w, http.StatusNotFound)
		return
	}

	// Get Customer
	isSaved := svc.GetServiceCustomer.UpdateCustomerByID(customer_id, customer)

	// Return Error Response
	if !isSaved {
		Response(false, entity.ErrNotUpdated.Error(), nil, w, http.StatusNotFound)
		return
	}

	// Return Successfull Response
	Response(true, "Customer updated", nil, w, http.StatusOK)
}

func (svc ServiceImpl) SaveCustomer(w http.ResponseWriter, r *http.Request) {

	var customer entity.Customer

	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		Response(false, entity.ErrInvalidJSON.Error(), nil, w, http.StatusNotFound)
		return
	}
	customer.GenerateUUID()
	customer.EncryptPassword()
	isSaved := svc.GetServiceCustomer.SaveCustomer(customer)

	// Return Error Response
	if !isSaved {
		Response(false, entity.ErrNotSaved.Error(), nil, w, http.StatusNotFound)
		return
	}

	// Return Successfull Response
	Response(true, "Customer created", nil, w, http.StatusOK)
}

func (svc ServiceImpl) GetCustomerByID(w http.ResponseWriter, r *http.Request) {

	// Get Customer ID
	params := mux.Vars(r)
	customer_id := params["id"]

	// validate param ID
	if customer_id == "" {
		Response(false, entity.ErrInvalidParameter.Error(), nil, w, http.StatusBadRequest)
		return
	}

	// customerCtx := r.Context().Value(entity.ContextCustomerKey).(*entity.Customer) // Get Customer from Context
	// fmt.Println(customerCtx)                                               // Send this context value to service and check in repository if customer record belongs to same customer

	// Get Customer
	customer, err := svc.GetServiceCustomer.GetCustomerByID(customer_id)

	// Return Error Response
	if err != nil {
		Response(false, entity.ErrNotFound.Error(), nil, w, http.StatusNotFound)
		return
	}

	// Return Successfull Response
	Response(true, "Customer found", customer.Public(), w, http.StatusOK)
}

func (svc ServiceImpl) DeleteCustomerByID(w http.ResponseWriter, r *http.Request) {

	// Get Customer ID
	params := mux.Vars(r)
	customer_id := params["id"]

	// validate param ID
	if customer_id == "" {
		Response(false, entity.ErrInvalidParameter.Error(), nil, w, http.StatusBadRequest)
		return
	}

	// Get Customer
	isDeleted := svc.GetServiceCustomer.DeleteCustomerByID(customer_id)

	// Return Error Response
	if !isDeleted {
		Response(false, entity.ErrCustomerNotDeleted.Error(), nil, w, http.StatusNotFound)
		return
	}

	// Return Successfull Response
	Response(true, "Customer deleted", nil, w, http.StatusOK)
}
