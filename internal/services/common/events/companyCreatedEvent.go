package events

import "github.com/any/companies/internal/domain/models"

const (
	CompanyCreatedEventName = "company_created"
)

type CompanyCreatedEvent struct {
	blankEvent
}

func NewCompanyCreatedEvent(
	company models.Company,
) CompanyCreatedEvent {
	return CompanyCreatedEvent{
		blankEvent: newBlankEvent(
			CompanyCreatedEventName,
			company,
		),
	}
}
