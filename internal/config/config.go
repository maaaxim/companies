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

// @TODO актуализировать
const (
	ProdEnv  = Env("prod")  // продовое окружение
	TestEnv  = Env("test")  // тестовое окружение
	TestsEnv = Env("tests") // тесты
	DevEnv   = Env("dev")   // разработка
)

func (e Env) validate() error {
	if e != ProdEnv && e != TestEnv && e != DevEnv && e != TestsEnv {
		return errors.Errorf("unknown env value - %v", e)
	}

	return nil
}

// func (e Env) toString() string {
//	return string(e)
// }

type Config struct {
	Env         Env                     `envconfig:"ENV" default:"test"`
	AppId       string                  `envconfig:"ID" default:"collections-service"`
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
