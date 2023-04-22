package jwtController

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"

	"github.com/any/companies/internal/api"
)

// @TODO implement sign up if i have more time
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type SigninRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (r SigninRequest) Validate() []error {
	var errs []error
	if len(r.Username) <= minUsername {
		errs = append(errs, errors.New("username must be longer then 3 symbols"))
	}
	if len(r.Password) <= minPassword {
		errs = append(errs, errors.New("password must be longer then 3 symbols"))
	}

	return errs
}

func (c Controller) SigninHandler(w http.ResponseWriter, r *http.Request) {
	signinRequest := SigninRequest{}
	if !api.ValidateRequest(&signinRequest, w, r) {

		return
	}
	expectedPassword, ok := users[signinRequest.Username]
	if !ok || expectedPassword != signinRequest.Password {
		w.WriteHeader(http.StatusUnauthorized)
		c.WriteErrorResponse(w, fmt.Errorf("wrong password"))

		return
	}

	expirationTime := time.Now().Add(tokenLifetimeMinutes * time.Minute)
	claims := &api.Claims{
		Username: signinRequest.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	api.WriteSuccessResponse(w, tokenString)
}
