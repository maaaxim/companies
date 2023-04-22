package events

const (
	CompanyDeletedEventName = "company_deleted"
)

type CompanyDeletedEvent struct {
	blankEvent
}

func NewCompanyDeletedEvent(
	uuid string,
) CompanyUpdatedEvent {
	return CompanyUpdatedEvent{
		blankEvent: newBlankEvent(
			CompanyDeletedEventName,
			uuid,
		),
	}
}
