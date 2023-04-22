package companyService

import (
	"github.com/pkg/errors"

	"github.com/any/companies/internal/services/common/events"
)

func (s Service) UpdateCompany(uuid string, dto CompanyDto) error {
	modelCompany, err := s.repository.GetCompany(uuid)
	if err != nil {
		return errors.Wrap(err, "GetCompany")
	}
	modelCompany.Update(
		dto.Name,
		dto.Description,
		dto.EmployeesAmount,
		dto.Registered,
		dto.Type,
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
