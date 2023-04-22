package companyService

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/any/companies/internal/domain/models"
	"github.com/any/companies/internal/services/common/events"
)

func (s Service) CreateCompany(company CompanyDto) (string, error) {
	companyUuid := uuid.New()
	companyType, err := models.NewCompanyTypeFromString(company.Type)
	if err != nil {
		return "", errors.Wrap(err, "NewCompanyTypeFromString")
	}
	companyModel, err := models.NewCompany(
		companyUuid.String(),
		company.Name,
		company.Description,
		company.EmployeesAmount,
		company.Registered,
		companyType,
	)
	err = s.repository.CreateCompany(companyModel)
	if err != nil {
		return "", errors.Wrap(err, "CreateCompany")
	}

	s.eventsPublisher.GoPublishEvent(
		events.NewCompanyUpdatedEvent(companyModel),
	)

	return companyUuid.String(), nil
}
