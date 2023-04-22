package companyService

import (
	"github.com/pkg/errors"

	"github.com/any/companies/internal/domain/models"
	"github.com/any/companies/internal/services/common/events"
)

func (s Service) UpdateCompany(uuid string, dto CompanyDto) error {
	modelCompany, err := s.repository.GetCompany(uuid)
	if err != nil {
		return errors.Wrap(err, "GetCompany")
	}
	companyType, err := models.NewCompanyTypeFromString(dto.Type)
	if err != nil {
		return errors.Wrap(err, "NewCompanyTypeFromString")
	}
	if err != nil {
		return errors.Wrap(err, "GetCompany")
	}
	modelCompany.Update(
		dto.Name,
		dto.Description,
		dto.EmployeesAmount,
		dto.Registered,
		companyType,
	)
	err = s.repository.UpdateCompany(modelCompany)
	if err != nil {

		return errors.Wrap(err, "UpdateCompany")
	}

	s.eventsPublisher.GoPublishEvent(
		events.NewCompanyUpdatedEvent(modelCompany),
	)

	return nil
}
