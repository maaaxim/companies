package postgres

import (
	"github.com/any/companies/internal/domain/models"
	"github.com/pkg/errors"
)

func (r Repository) UpdateCompany(company models.Company) error {
	query := `
		update companies as c set
			name = :name
			, description = :description
			, employees_amount = :employees_amount
			, registered = :registered
			, type = :type
		where c.uuid = :uuid
	`
	dbCompany := convertCompanyModelToDb(company)
	res, err := r.DB.NamedExec(query, dbCompany)
	if err != nil {

		return errors.Wrap(err, "NamedExec")
	}
	count, err := res.RowsAffected()
	if err != nil {

		return errors.Wrap(err, "RowsAffected")
	}
	if count != 1 {

		return errors.New("mismatched affected rows")
	}

	return nil
}
