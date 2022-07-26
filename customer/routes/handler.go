package routes

import (
	"net/http"

	"github.com/eddwinpaz/customer-svc/customer/controller"
	"github.com/eddwinpaz/customer-svc/customer/middleware"
	"github.com/gorilla/mux"
)

// var mapping = "/api/customer/"

func Handlers(controllers controller.ServiceImpl) http.Handler {

	route := mux.NewRouter()

	// Unauthenticated endpoints
	unauthenticated := route.NewRoute().Subrouter()
	unauthenticated.HandleFunc("/api/customer/health", controllers.HealthCheck).Methods("GET")
	unauthenticated.HandleFunc("/api/customer/auth", controllers.CustomerAuthentication).Methods("POST")
	unauthenticated.HandleFunc("/api/customer", controllers.SaveCustomer).Methods("POST")

	// Authenticated endpoints
	authenticated := route.NewRoute().Subrouter()
	authenticated.HandleFunc("/api/customer/{id}", controllers.GetCustomerByID).Methods("GET")
	authenticated.HandleFunc("/api/customer/{id}", controllers.UpdateCustomerByID).Methods("PUT")
	// authenticated.HandleFunc("/api/customer", controllers.SaveCustomer).Methods("POST")
	authenticated.HandleFunc("/api/customer/{id}", controllers.DeleteCustomerByID).Methods("DELETE")
	authenticated.Use(middleware.JwtAuthentication)

	return route
}
