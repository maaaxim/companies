package companyService

import (
	"github.com/google/uuid"

	"github.com/any/companies/internal/domain/models"
)

func (s Service) CreateCompany(company CompanyDto) (string, error) {
	companyUuid := uuid.New()
	companyModel, err := models.NewCompany(
		companyUuid.String(),
		company.Name,
		company.Description,
		company.EmployeesAmount,
		company.Registered,
		company.Type,
	)
	err = s.repository.CreateCompany(companyModel)
	if err != nil {
		return "", err
	}
	return companyUuid.String(), nil
}
