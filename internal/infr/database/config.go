package database

import "time"

type PostgresConfig struct {
	Server          string        `envconfig:"SERVER"`
	Port            uint16        `envconfig:"PORT"`
	User            string        `envconfig:"USER"`
	Password        string        `envconfig:"PASSWORD"`
	DatabaseName    string        `envconfig:"DB_NAME"`
	DialTimeout     string        `envconfig:"DIAL_TIMEOUT" default:"2s"`
	MaxOpenConns    int           `envconfig:"MAXOPENCONNS"`
	MaxIdleConns    int           `envconfig:"MAXIDLECONNS" default:"2"`
	ConnMaxLifetime time.Duration `envconfig:"CONNMAXLIFETIME" default:"10m"`
}
