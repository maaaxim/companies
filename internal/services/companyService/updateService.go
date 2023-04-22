package companyService

import "github.com/pkg/errors"

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
	return nil
}
