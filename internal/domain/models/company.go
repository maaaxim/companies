package models

type Company struct {
	Uuid            string
	Name            string
	Description     string
	EmployeesAmount int
	Registered      bool
	Type            CompanyType
}

func NewCompany(
	uuid string,
	name string,
	description string,
	employeesAmount int,
	registered bool,
	companyType CompanyType,
) (Company, error) {
	return Company{
		Uuid:            uuid,
		Name:            name,
		Description:     description,
		EmployeesAmount: employeesAmount,
		Registered:      registered,
		Type:            companyType,
	}, nil
}

func (c *Company) Update(
	name string,
	description string,
	employeesAmount int,
	registered bool,
	companyType CompanyType,
) {
	c.Name = name
	c.Description = description
	c.EmployeesAmount = employeesAmount
	c.Registered = registered
	c.Type = companyType
}
