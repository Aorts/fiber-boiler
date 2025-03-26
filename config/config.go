package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Server server
	Mongo  mongo
	JWT    jwt
}

type jwt struct {
	Secret string `envconfig:"JWT_SECRET" default:"JWT_SECRET"`
}

type server struct {
	Port string `envconfig:"PORT" default:"8080"`
}

type mongo struct {
	URI      string `envconfig:"MONGO_URI" default:"mongodb://localhost:27017"`
	Database string `envconfig:"DB_NAME" default:"user-api"`
}

var cfg config

func New() {
	_ = godotenv.Load()
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("read env error : %s", err.Error())
	}
}

func Get() config {
	return cfg
}
