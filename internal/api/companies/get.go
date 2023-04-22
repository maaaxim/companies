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

	// @TODO service
	getResponse := GetResponse{
		Uuid:            uuid,
		Name:            "2",
		Description:     "3",
		EmployeesAmount: 4,
		Registered:      false,
		Type:            "5555",
	}

	api.WriteSuccessResponse(w, getResponse)
}
