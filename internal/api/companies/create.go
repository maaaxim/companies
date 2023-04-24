package companiesController

import (
	"errors"
	"net/http"

	"github.com/any/companies/internal/api"
	"github.com/any/companies/internal/services/companyService"
)

type CreateRequest struct {
	Name            string `json:"name" validate:"required"`
	Description     string `json:"description" validate:"optional"`
	EmployeesAmount int    `json:"employeesAmount" validate:"required"`
	Registered      bool   `json:"registered" validate:"required"`
	Type            string `json:"type" validate:"required"`
}

func (r CreateRequest) Validate() []error {
	var errs []error
	if len(r.Description) > maxDescription {
		errs = append(errs, errors.New("description length must be less then 3000 symbols"))
	}

	if len(r.Name) > maxName {
		errs = append(errs, errors.New("name must be less then 15 symbols"))
	}

	return errs
}

func (c Controller) CreateHandler(w http.ResponseWriter, r *http.Request) {
	createRequest := CreateRequest{}
	if !api.ValidateRequest(&createRequest, w, r) {

		return
	}
	dto := c.makeCreateCompanyDto(
		createRequest.Name,
		createRequest.Description,
		createRequest.EmployeesAmount,
		createRequest.Registered,
		createRequest.Type,
	)
	uuid, err := c.companiesService.CreateCompany(dto)
	if err != nil {
		c.WriteErrorResponse(w, err)

		return
	}

	api.WriteSuccessResponse(w, uuid)
}

func (c Controller) makeCreateCompanyDto(
	name string,
	description string,
	employeesAmount int,
	registered bool,
	theType string,
) companyService.CompanyDto {
	dto := companyService.CompanyDto{
		Name:            name,
		Description:     description,
		EmployeesAmount: employeesAmount,
		Registered:      registered,
		Type:            theType,
	}

	return dto
}
