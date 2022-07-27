package middleware

import (
	"context"
	"net/http"

	"github.com/eddwinpaz/customer-svc/customer/controller"
	"github.com/eddwinpaz/customer-svc/customer/entity"
	log "github.com/sirupsen/logrus"
)

func JwtAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var credentials entity.Credentials
		authorization_token := r.Header.Get("Authorization")

		if authorization_token == "" {
			controller.Response(false, entity.ErrMissingAuthorizationToken.Error(), nil, w, http.StatusBadRequest)
			return
		}

		claims, err := credentials.ValidateJwtToken(authorization_token)

		if err != nil {
			controller.Response(false, entity.ErrInvalidAuthorizationToken.Error(), err, w, http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), entity.ContextCustomerKey, claims.Customer)
		log.Info("JwtAuthentication executed...")
		log.Info(claims.Customer)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
