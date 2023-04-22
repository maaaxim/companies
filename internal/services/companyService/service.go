package companyService

import (
	"github.com/any/companies/internal/domain/models"
)

type Service struct {
	repository companyRepository
	//eventsPublisher eventsPublisher
}

type companyRepository interface {
	GetCompany(uuid string) (models.Company, error)
	CreateCompany(company models.Company) error
	UpdateCompany(company models.Company) error
	DeleteCompany(uuid string) error
}

//type Event interface {
//	GetName() string
//	Marshal() ([]byte, error)
//}
//
//type eventsPublisher interface {
//	GoPublishEvent(event Event)
//}

func NewService(
	repository companyRepository,
	// eventsPublisher eventsPublisher,
) Service {
	return Service{
		repository: repository,
		//eventsPublisher: eventsPublisher,
	}
}

type CompanyDto struct {
	Name            string
	Description     string
	EmployeesAmount int
	Registered      bool
	Type            string
}
