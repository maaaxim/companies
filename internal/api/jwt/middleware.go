package jwtController

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"

	"github.com/any/companies/internal/api"
)

func (c Controller) JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var tknStr = r.Header.Get("token")
		if tknStr == "" {
			w.WriteHeader(http.StatusForbidden)
			c.WriteErrorResponse(w, fmt.Errorf("empty token"))

			return
		}

		claims := &api.Claims{}
		_, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {

			return JwtKey, nil
		})

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			c.WriteErrorResponse(w, err)

			return
		}

		next.ServeHTTP(w, r)
	})
}
