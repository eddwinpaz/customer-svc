package repository

import (
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/eddwinpaz/customer-svc/customer/entity"
	_ "github.com/lib/pq"
)

type PostgresCustomerRespository struct {
	db *sql.DB
}

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "root"
	dbname   = "realstate"
)

const DATABASE_DRIVER = "postgres"

func OpenConnection() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open(DATABASE_DRIVER, psqlconn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (repository *PostgresCustomerRespository) CloseConnection() {
	_ = repository.db.Close()
}

func NewPostgresCustomerRespository(db *sql.DB) *PostgresCustomerRespository {
	return &PostgresCustomerRespository{db}
}

func (repository *PostgresCustomerRespository) HealthCheck() error {

	log.Infof("Connecting...")
	log.Infof("Database: %s ", dbname)
	log.Infof("User: %s", user)

	if err := repository.db.Ping(); err != nil {
		log.Error(err)
		defer repository.db.Close() //nolint
		return err
	}
	return nil
}

func (repository *PostgresCustomerRespository) AuthenticateCustomer(email string, password string) (*entity.Customer, error) {

	customer := &entity.Customer{}

	query := fmt.Sprintf(`SELECT customer_uuid, first_name, last_name, email, "password", created_on, last_login 
			  			  FROM customer 
			  			  WHERE email = '%s' AND password = '%s'`, email, password)

	row := repository.db.QueryRow(query)
	err := row.Scan(&customer.CustomerID, &customer.FirstName, &customer.LastName, &customer.Email,
		&customer.Password, &customer.CreatedOn, &customer.LastLogin)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, entity.ErrNotFound
		}
		return nil, entity.ErrSQLError
	}
	return customer, nil
}

func (repository *PostgresCustomerRespository) UpdateCustomerByID(customer entity.Customer) bool {

	query := fmt.Sprintf(`UPDATE customer SET first_name='%s', last_name='%s' WHERE customer_uuid = '%s'`,
		customer.FirstName, customer.LastName, customer.CustomerID)
	result, err := repository.db.Exec(query)

	if err != nil {
		return false
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		return false
	}
	return true
}

func (repository *PostgresCustomerRespository) SaveCustomer(customer entity.Customer) bool {

	query := fmt.Sprintf(`INSERT INTO customer (customer_uuid,first_name, last_name, email, "password", created_on, last_login) 
						  VALUES('%s','%s', '%s', '%s', '%s', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);`,
		customer.CustomerID, customer.FirstName, customer.LastName, customer.Email, customer.Password)

	result, err := repository.db.Exec(query)
	log.Infof("Executing Query: SaveCustomer")
	log.Infof(query)

	if err != nil {
		log.Error(err)
		return false
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		return false
	}
	return true
}

func (repository *PostgresCustomerRespository) GetCustomerByID(customer_id string) (*entity.Customer, error) {

	customer := &entity.Customer{}

	query := fmt.Sprintf(`SELECT customer_uuid, first_name, last_name, email, "password", created_on, last_login 
			  			  FROM customer 
			  			  WHERE customer_uuid = '%s'`, customer_id)

	row := repository.db.QueryRow(query)

	err := row.Scan(&customer.CustomerID, &customer.FirstName, &customer.LastName, &customer.Email, &customer.Password, &customer.CreatedOn, &customer.LastLogin)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, entity.ErrNotFound
		}
		return nil, entity.ErrSQLError
	}

	return customer, nil
}

func (repository *PostgresCustomerRespository) DeleteCustomerByID(customer_id string) bool {

	query := fmt.Sprintf(`DELETE FROM customer WHERE customer_uuid = '%s'`, customer_id)
	_, err := repository.db.Exec(query)

	if err != nil {
		fmt.Print(err)
		if err == sql.ErrNoRows {
			return false
		}
		return false
	}

	return true
}
