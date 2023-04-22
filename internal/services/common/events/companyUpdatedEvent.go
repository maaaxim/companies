package events

import "github.com/any/companies/internal/domain/models"

const (
	CompanyUpdatedEventName = "company_updated"
)

type CompanyUpdatedEvent struct {
	blankEvent
}

func NewCompanyUpdatedEvent(
	company models.Company,
) CompanyUpdatedEvent {
	return CompanyUpdatedEvent{
		blankEvent: newBlankEvent(
			CompanyUpdatedEventName,
			company,
		),
	}
}
