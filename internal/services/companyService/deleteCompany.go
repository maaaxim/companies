package companyService

import (
	"github.com/pkg/errors"

	"github.com/any/companies/internal/services/common/events"
)

func (s Service) DeleteCompany(uuid string) error {
	err := s.repository.DeleteCompany(uuid)
	if err != nil {

		return errors.Wrap(err, "DeleteCompany")
	}

	s.eventsPublisher.GoPublishEvent(
		events.NewCompanyDeletedEvent(uuid),
	)

	return nil
}
