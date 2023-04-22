package companiesController

import (
	"github.com/any/companies/internal/api"
	"github.com/any/companies/internal/infr/logger"
	"github.com/any/companies/internal/services/companyService"
)

type Controller struct {
	api.Controller
	companiesService companyService.Service
}

func NewController(
	logger logger.Logger,
	companiesService companyService.Service,
) Controller {
	return Controller{
		api.NewController(logger),
		companiesService,
	}
}
