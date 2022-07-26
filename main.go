package main

import (
	"net"
	"net/http"
	"os"

	"github.com/eddwinpaz/customer-svc/customer/controller"
	"github.com/eddwinpaz/customer-svc/customer/entity"
	"github.com/eddwinpaz/customer-svc/customer/repository"
	"github.com/eddwinpaz/customer-svc/customer/routes"
	"github.com/eddwinpaz/customer-svc/customer/service"
	"github.com/eddwinpaz/customer-svc/logging"
	"github.com/gorilla/handlers"
	log "github.com/sirupsen/logrus"
)

func setupLogging() {
	logging.InitializeLogging("vendor.log")
}

func main() {

	setupLogging()

	port := ":9000" //os.Getenv("PORT") // ":9000"
	server, err := net.Listen("tcp", port)

	if err != nil {
		log.Infof("Port %s is in use... ", port)
		return
	}
	_ = server.Close()

	// Open Dataase Connection
	db, err := repository.OpenConnection()

	if err != nil {
		log.Errorf("%s %s", entity.ErrDatabaseConnection.Error(), err)
		os.Exit(1)
	}

	// Postgres repository init
	repo := repository.NewPostgresCustomerRespository(db)

	defer repo.CloseConnection()

	// Dependencies

	// Service init
	services := service.RepositoryImpl{
		GetRespositoryCustomer: repo,
	}

	// Controller init
	controllers := &controller.ServiceImpl{
		GetServiceCustomer: services,
	}

	route := routes.Handlers(*controllers)

	log.Infof("Server running on port %s", port)

	loggedRouter := handlers.LoggingHandler(os.Stdout, route)

	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	err = http.ListenAndServe(port, handlers.CORS(headers, methods, origins)(loggedRouter))

	if err != nil {
		log.Errorf("%s %s\n", entity.ErrStartingServer.Error(), err)
		os.Exit(1)
	}

}
