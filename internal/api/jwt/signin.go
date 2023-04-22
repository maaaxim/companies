package jwtController

import (
	"errors"
	"fmt"
	"github.com/any/companies/internal/api"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

// @TODO implement sign up if i have more time
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type SigninRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r SigninRequest) Validate() []error {
	var errs []error
	if len(r.Username) <= 0 {
		errs = append(errs, errors.New("empty username"))
	}
	if len(r.Password) <= 0 {
		errs = append(errs, errors.New("empty password"))
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

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
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
