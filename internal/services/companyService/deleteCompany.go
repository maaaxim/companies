package companyService

import "github.com/pkg/errors"

func (s Service) DeleteCompany(uuid string) error {
	err := s.repository.DeleteCompany(uuid)
	if err != nil {

		return errors.Wrap(err, "DeleteCompany")
	}

	return nil
}
