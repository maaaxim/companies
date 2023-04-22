package companiesController

import (
	"github.com/gorilla/mux"
	"net/http"

	"github.com/any/companies/internal/api"
)

type GetRequest struct {
	Id string `json:"id" validate:"required"`
}

type GetResponse struct {
	Uuid            string `json:"uuid"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	EmployeesAmount int    `json:"employeesAmount"`
	Registered      bool   `json:"registered"`
	Type            string `json:"type"`
}

func (c Controller) GetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]

	m, err := c.companiesService.GetCompany(uuid)
	if err != nil {
		c.WriteErrorResponse(w, err)

		return
	}

	getResponse := GetResponse{
		Uuid:            uuid,
		Name:            m.Name,
		Description:     m.Description,
		EmployeesAmount: m.EmployeesAmount,
		Registered:      m.Registered,
		Type:            m.Type,
	}

	api.WriteSuccessResponse(w, getResponse)
}
