package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/any/companies/internal/domain/models"
	"github.com/any/companies/internal/infr/database"
)

type Repository struct {
	DB *sqlx.DB
}

func NewPostgresRepository(cfg database.PostgresConfig) (Repository, error) {
	db, err := sqlx.Open("postgres", GetPostgresDsn(cfg))
	repo := Repository{
		DB: db,
	}
	if err != nil {
		return repo, err //nolint:wrapcheck
	}

	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)

	if !repo.Ping() {
		return repo, fmt.Errorf("ping connection error")
	}

	return repo, nil
}

func GetPostgresDsn(cfg database.PostgresConfig) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.Server,
		cfg.User,
		cfg.Password,
		cfg.DatabaseName,
		cfg.Port,
	)
}

type DbCompany struct {
	Id              int    `db:"id"`
	Uuid            string `db:"uuid"`
	Name            string `db:"name"`
	Description     string `db:"description"`
	EmployeesAmount int    `db:"employees_amount"`
	Registered      bool   `db:"registered"`
	Type            string `db:"type"`
}

func convertCompanyModelToDb(company models.Company) DbCompany {
	return DbCompany{
		Uuid:            company.Uuid,
		Name:            company.Name,
		Description:     company.Description,
		EmployeesAmount: company.EmployeesAmount,
		Registered:      company.Registered,
		Type:            company.Type,
	}
}

func convertCompanyFromDb(company DbCompany) models.Company {
	return models.Company{
		Uuid:            company.Uuid,
		Name:            company.Name,
		Description:     company.Description,
		EmployeesAmount: company.EmployeesAmount,
		Registered:      company.Registered,
		Type:            company.Type,
	}
}
