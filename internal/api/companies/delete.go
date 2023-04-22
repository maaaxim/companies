package companiesController

import (
	"net/http"

	"github.com/any/companies/internal/api"
)

func (c Controller) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//uuid := vars["uuid"]

	api.WriteSuccessResponse(w, "ok")
}
