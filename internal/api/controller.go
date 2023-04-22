package api

import (
	"encoding/json"
	"net/http"

	"github.com/any/companies/internal/infr/logger"
)

type Controller struct {
	logger logger.Logger
}

func (c Controller) WriteErrorsResponse(writer http.ResponseWriter, errs []error) {
	c.logger.Errors(errs)
	writeResponse(writer, newErrorsResponse(errs))
}

func (c Controller) WriteErrorResponse(writer http.ResponseWriter, err error) {
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
