package companiesController

import (
	"errors"
	"net/http"

	"github.com/any/companies/internal/api"
)

type PatchRequest struct {
	Id              string `json:"id" validate:"required"`
	Name            string `json:"name" validate:"optional"`
	Description     string `json:"description" validate:"optional"`
	EmployeesAmount int    `json:"employeesAmount" validate:"optional"`
	Registered      bool   `json:"registered" validate:"optional"`
	Type            string `json:"type" validate:"optional"`
}

// @TODO check validations best practices
func (r PatchRequest) Validate() []error {
	var errs []error
	if len(r.Id) <= 0 {
		errs = append(errs, errors.New("wrong id"))
	}

	return errs
}

func (c Controller) PatchHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//uuid := vars["uuid"]

	patchRequest := PatchRequest{}
	if !api.ValidateRequest(&patchRequest, w, r) {

		return
	}

	api.WriteSuccessResponse(w, "ok")
}
