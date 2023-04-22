package jwtController

import (
	"os"

	"github.com/any/companies/internal/api"
	"github.com/any/companies/internal/infr/logger"
)

const minUsername = 3
const minPassword = 3
const tokenLifetimeMinutes = 5

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
