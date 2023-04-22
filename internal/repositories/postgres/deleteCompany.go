package postgres

import "github.com/pkg/errors"

func (r Repository) DeleteCompany(uuid string) error {
	query := `
		delete from companies
		where uuid = $1;
	`
	_, err := r.DB.Exec(query, uuid)
	if err != nil {
		return errors.Wrap(err, "Exec")
	}

	return nil
}
