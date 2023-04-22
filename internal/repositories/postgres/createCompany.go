package postgres

import (
	"github.com/pkg/errors"

	"github.com/any/companies/internal/domain/models"
)

func (r Repository) CreateCompany(company models.Company) error {
	query := `
		insert into companies (
	   		uuid
	   		, name
	   		, description
	   		, employees_amount
	   		, registered
	   		, type
		) values (
			 :uuid,
			 :name,
			 :description,
		     :employees_amount,
		     :registered,
		     :type
		)
	`
	dbCompany := convertCompanyModelToDb(company)
	_, err := r.DB.NamedQuery(query, dbCompany)
	if err != nil {
		return errors.Wrap(err, "NamedQuery")
	}

	return nil
}
