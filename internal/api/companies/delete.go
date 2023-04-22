package companiesController

import (
	"github.com/gorilla/mux"
	"net/http"

	"github.com/any/companies/internal/api"
)

func (c Controller) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	err := c.companiesService.DeleteCompany(uuid)
	if err != nil {
		c.WriteErrorResponse(w, err)

		return
	}

	api.WriteSuccessResponse(w, "ok")
}
