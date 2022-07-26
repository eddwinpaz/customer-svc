package entity

import (
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/eddwinpaz/customer-svc/customer/utils"
)

// jwtKey is the key used to sign the JWT
var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	Customer *Customer `json:"customer"`
	jwt.StandardClaims
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *Credentials) EncryptPassword() {
	c.Password = utils.EncryptPassword(c.Password)
}

func (c *Credentials) GenerateJwtToken(customer *Customer) (interface{}, error) {

	expirationTime := time.Now().Add(time.Minute * 60)

	claims := &Claims{
		Customer: customer.Public(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return nil, err
	}
	return interface{}(map[string]interface{}{
		"token": tokenString,
		"customer":  customer.Public(),
	}), nil
}

func (c *Credentials) ValidateJwtToken(authorization_header string) (Claims, error) {
	claims := &Claims{}

	authorization := strings.Split(authorization_header, " ")

	if authorization[0] != "Bearer" {
		return *claims, jwt.ErrInvalidKey
	}

	tkn, err := jwt.ParseWithClaims(authorization[1], claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return Claims{}, jwt.ErrSignatureInvalid
		}
	}
	if !tkn.Valid {
		return Claims{}, jwt.ErrSignatureInvalid
	}

	return *claims, nil
}
