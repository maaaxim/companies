package companyService

import (
	"github.com/pkg/errors"

	"github.com/any/companies/internal/domain/models"
)

func (s Service) GetCompany(uuid string) (models.Company, error) {
	modelCompany, err := s.repository.GetCompany(uuid)
	if err != nil {
		return modelCompany, errors.Wrap(err, "GetCompany")
	}

	return modelCompany, nil
}
