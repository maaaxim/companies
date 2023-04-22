package companiesController

import (
	"github.com/gorilla/mux"
	"net/http"

	"github.com/any/companies/internal/api"
)

type PatchRequest struct {
	Name            string `json:"name" validate:"optional"`
	Description     string `json:"description" validate:"optional"`
	EmployeesAmount int    `json:"employeesAmount" validate:"optional"`
	Registered      bool   `json:"registered" validate:"optional"`
	Type            string `json:"type" validate:"optional"`
}

// @TODO check validations best practices
func (r PatchRequest) Validate() []error {
	var errs []error
	// @TODO description check
	return errs
}

func (c Controller) PatchHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]

	patchRequest := PatchRequest{}
	if !api.ValidateRequest(&patchRequest, w, r) {

		return
	}
	dto := c.makeCreateCompanyDto(
		patchRequest.Name,
		patchRequest.Description,
		patchRequest.EmployeesAmount,
		patchRequest.Registered,
		patchRequest.Type,
	)
	err := c.companiesService.UpdateCompany(uuid, dto)
	if err != nil {
		c.WriteErrorResponse(w, err)

		return
	}

	api.WriteSuccessResponse(w, "ok")
}
