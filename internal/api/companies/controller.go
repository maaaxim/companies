package companiesController

import (
	"github.com/any/companies/internal/api"
	"github.com/any/companies/internal/infr/logger"
)

type Controller struct {
	api.Controller
}

func NewController(
	logger logger.Logger,
) Controller {
	return Controller{
		api.NewController(logger),
	}
}
