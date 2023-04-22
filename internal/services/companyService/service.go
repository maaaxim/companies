package companyService

import (
	"github.com/any/companies/internal/domain/models"
	"github.com/any/companies/internal/services/common/events"
)

type Service struct {
	repository      companyRepository
	eventsPublisher events.EventsPublisher
}

type companyRepository interface {
	GetCompany(uuid string) (models.Company, error)
	CreateCompany(company models.Company) error
	UpdateCompany(company models.Company) error
	DeleteCompany(uuid string) error
}

func NewService(
	repository companyRepository,
	eventsPublisher events.EventsPublisher,
) Service {
	return Service{
		repository:      repository,
		eventsPublisher: eventsPublisher,
	}
}

type CompanyDto struct {
	Name            string
	Description     string
	EmployeesAmount int
	Registered      bool
	Type            string
}
