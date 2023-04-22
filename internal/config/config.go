package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"

	"github.com/any/companies/internal/infr/database"
	"github.com/any/companies/internal/infr/kafka"
	"github.com/any/companies/internal/infr/logger"
	"github.com/any/companies/internal/infr/server"
)

type Env string

const (
	ProdEnv = Env("prod")
	TestEnv = Env("test")
	DevEnv  = Env("dev")
)

func (e Env) validate() error {
	if e != ProdEnv && e != TestEnv && e != DevEnv {
		return errors.Errorf("unknown env value - %v", e)
	}

	return nil
}

type Config struct {
	Env         Env                     `envconfig:"ENV" default:"test"`
	AppId       string                  `envconfig:"ID" default:"companies"`
	Server      server.Config           `envconfig:"SERVER" required:"true"`
	Logger      logger.Config           `envconfig:"LOGGER" required:"true"`
	Postgres    database.PostgresConfig `envconfig:"COMPANIES_DATABASE" required:"true" vault:"postgres/"`
	KafkaWriter kafka.WriterConfig      `envconfig:"COMPANIES_KAFKA_WRITER" required:"true"`
}

func InitConfig(prefix string) (
	Config,
	error,
) {
	_ = godotenv.Load(".env")

	var cfg Config
	if err := envconfig.Process(prefix, &cfg); err != nil {
		return Config{}, errors.Wrap(err, "get config from env error")
	}

	if err := cfg.Env.validate(); err != nil {
		return Config{}, errors.Wrap(err, "validate")
	}

	return cfg, nil
}
