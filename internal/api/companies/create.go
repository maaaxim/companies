package companiesController

import (
	"github.com/any/companies/internal/api"
	"net/http"
)

/*
• ID (uuid) required
• Name (15 characters) required - unique
• Description (3000 characters) optional
• Amount of Employees (int) required
• Registered (boolean) required
  @TODO enum
• Type (Corporations | NonProfit | Cooperative | Sole Proprietorship) required
*/

type CreateRequest struct {
	Name            string `json:"name" validate:"required"`
	Description     string `json:"description" validate:"optional"`
	EmployeesAmount int    `json:"employeesAmount" validate:"required"`
	Registered      bool   `json:"registered" validate:"required"`
	Type            string `json:"type" validate:"required"`
}

// @TODO check validations best practices
func (r CreateRequest) Validate() []error {
	var errs []error
	// validate description @TODO
	return errs
}

func (c Controller) CreateHandler(w http.ResponseWriter, r *http.Request) {
	createRequest := CreateRequest{}
	if !api.ValidateRequest(&createRequest, w, r) {

		return
	}

	api.WriteSuccessResponse(w, "Uuid code new one")
}
