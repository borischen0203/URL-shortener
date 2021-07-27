package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env"
)

type envConfig struct {
	Version               string `env:"VERSION" envDefault:"0.0.1"`
	DBURL                 string `env:"DB_URL,required"`
	DBName                string `env:"DB_Name,required"`
	URL_HOST              string `env:"URL_HOST,required"`
	Port                  string `env:"PORT,required"`
	UrlInfoCollectionName string `env:"Url_Info_Collection_Name,required"`
}

var (
	// Env is the config
	Env = envConfig{}
)

// Setup setup config function
func Setup() {
	if err := env.Parse(&Env); err != nil {
		log.Fatalf("%+v\n", err)
	}

	fmt.Printf("%+v\n", Env)
}
