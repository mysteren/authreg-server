package main

import (
	"log"

	"github.com/joho/godotenv"
	"gitlab.devkeeper.com/authreg/server/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	// flag.StringVar(&configPath, "config", "configs/apiserver.json", "path to config")
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	c := apiserver.NewConfig()

	s := apiserver.New(c)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
