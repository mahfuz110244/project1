package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

const DEFAULT_MAX_OPEN_CONNS = 5

type (

	// Config -.
	Config struct {
		HTTP  `mapstructure:",squash"`
		Debug bool `env:"DEBUG" env-default:"false"`
	}
	// App -.
	App struct {
		Name    string `mapstructure:"NAME" json:"name"`
		Version string `mapstructure:"VERSION" json:"version"`
	}

	// HTTP -.
	HTTP struct {
		HTTPAddress string `env:"HTTP_ADDRESS"`
	}
)

// Read properties from config.env file
// Command line enviroment variable will overwrite config.env properties
func NewConfig(configFile string) *Config {
	config := Config{}
	err := cleanenv.ReadConfig(configFile, &config)
	if err != nil {
		log.Fatalln(err)
	}
	cleanenv.ReadEnv(&config)
	//fmt.Printf("%#v", config)
	//viper.SetDefault("SetMaxOpenConns", DEFAULT_MAX_OPEN_CONNS) // Set Maximum default connection to 5 when not provided
	return &config
}
