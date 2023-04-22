package jwtController

import (
	"os"

	"github.com/any/companies/internal/api"
	"github.com/any/companies/internal/infr/logger"
)

var JwtKey = []byte(os.Getenv("JWT_KEY"))

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
