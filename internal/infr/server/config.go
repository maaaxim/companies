package server

import "time"

type Config struct {
	Http HttpConfig `envconfig:"HTTP"`
}

type HttpConfig struct {
	Port         string        `envconfig:"PORT" required:"true"`
	ReadTimeout  time.Duration `envconfig:"READ_TIMEOUT" default:"10s"`
	WriteTimeout time.Duration `envconfig:"WRITE_TIMEOUT" default:"10s"`
}

func (c Config) getHttpAddr() string {
	return ":" + c.Http.Port
}
