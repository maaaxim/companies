package api

import (
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"net/http"

	"github.com/any/companies/internal/infr/logger"
)

type Controller struct {
	logger logger.Logger
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (c Controller) WriteErrorResponse(writer http.ResponseWriter, err error) {
	c.logger.Error(err.Error())
	writeResponse(writer, newErrorsResponse([]error{err}))
}

func NewController(logger logger.Logger) Controller {
	return Controller{
		logger: logger,
	}
}

type Validatable interface {
	Validate() []error
}

func ValidateRequest(req Validatable, writer http.ResponseWriter, request *http.Request) bool {
	if err := json.NewDecoder(request.Body).Decode(req); err != nil {
		writeResponse(writer, newErrorsResponse([]error{err}))

		return false
	}
	if errs := req.Validate(); len(errs) > 0 {
		writeResponse(writer, newErrorsResponse(errs))

		return false
	}

	return true
}

func WriteSuccessResponse(writer http.ResponseWriter, jsonableResponse any) {
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(newSuccessResponse(jsonableResponse)) //nolint:errcheck
}

func writeResponse(writer http.ResponseWriter, response any) {
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response) //nolint:errcheck
}
