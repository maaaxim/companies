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
	Uuid string,
	Name string,
	Description string,
	EmployeesAmount int,
	Registered bool,
	Type CompanyType,
) (Company, error) {
	return Company{
		Uuid:            Uuid,
		Name:            Name,
		Description:     Description,
		EmployeesAmount: EmployeesAmount,
		Registered:      Registered,
		Type:            Type,
	}, nil
}

func (c *Company) Update(
	Name string,
	Description string,
	EmployeesAmount int,
	Registered bool,
	Type CompanyType,
) {
	c.Name = Name
	c.Description = Description
	c.EmployeesAmount = EmployeesAmount
	c.Registered = Registered
	c.Type = Type
}
