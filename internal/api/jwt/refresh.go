package jwtController

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"

	"github.com/any/companies/internal/api"
)

type RefreshRequest struct {
	Token string `json:"token"`
}

func (r RefreshRequest) Validate() []error {
	var errs []error
	if len(r.Token) == 0 {
		errs = append(errs, errors.New("empty token"))
	}

	return errs
}

func (c Controller) RefreshHandler(w http.ResponseWriter, r *http.Request) {
	refreshRequest := RefreshRequest{}
	if !api.ValidateRequest(&refreshRequest, w, r) {

		return
	}

	tknStr := refreshRequest.Token
	claims := &api.Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {

		return JwtKey, nil
	})
	if err != nil {

		if errors.Is(err, jwt.ErrSignatureInvalid) {
			w.WriteHeader(http.StatusUnauthorized)

			return
		}
		w.WriteHeader(http.StatusBadRequest)

		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// 30 seconds of expiry. Otherwise, return a bad request status
	if time.Until(time.Unix(claims.ExpiresAt, 0)) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expirationTime := time.Now().Add(tokenLifetimeMinutes * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	api.WriteSuccessResponse(w, tokenString)
}
