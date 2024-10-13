package config

import (
	"log"

	"github.com/Netflix/go-env"
)

type Environment struct {
	Env    string `env:"ENVIRONMENT"`
	Port   string `env:"PORT,default=1323"`
	DbPath string `env:"DB_PATH,default=c-ademy.db"`
}

func GetConfig() (*Environment, error) {
	var environment Environment
	_, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		log.Fatal(err)
	}
	// ...
	return &environment, nil
}
