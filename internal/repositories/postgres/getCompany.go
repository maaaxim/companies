package postgres

import (
	"github.com/pkg/errors"

	"github.com/any/companies/internal/domain/models"
)

func (r Repository) GetCompany(uuid string) (models.Company, error) {
	var company models.Company
	var err error
	var dbCompany DbCompany

	queryString := `
		SELECT
		    c.uuid
			, c.name	  
			, c.description	  
			, c.employees_amount	  
			, c.registered	  
			, c.type  
		FROM
			companies as c
		WHERE
		    c.uuid = $1::varchar
	`

	err = r.DB.QueryRowx(queryString, uuid).StructScan(&dbCompany)
	if err != nil {

		return company, errors.Wrap(err, "StructScan")
	}

	company, err = convertCompanyFromDb(dbCompany)
	if err != nil {
		return company, errors.Wrap(err, "convertCompanyFromDb")
	}

	return company, nil
}
