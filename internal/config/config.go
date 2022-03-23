package config

import (
	"flag"
	"github.com/caarlos0/env"
	"log"
	"sync"
)

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS" envDefault:"localhost:8080"`
	FilePath      string `env:"FILE_PATH" envDefault:"images/"`
}

var Cfg Config

func New() {
	var serverAddress string
	var filePath string
	var once sync.Once

	once.Do(func() {
		if err := env.Parse(&Cfg); err != nil {
			log.Fatal(err)
		}

		flag.StringVar(&serverAddress, "a", Cfg.ServerAddress, "An address and a port of service.")
		flag.StringVar(&filePath, "p", Cfg.FilePath, "A directory for images")

		flag.Parse()

		Cfg.ServerAddress = serverAddress
		Cfg.FilePath = filePath
	})
}
