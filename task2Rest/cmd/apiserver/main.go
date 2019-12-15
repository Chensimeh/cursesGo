package main

import (
	"github.com/BurntSushi/toml"
	"github.com/trzhensimekh/cursesGo/task2Rest/internal/app"
	"log"
)

func main() {
	config := app.NewConfig()
	_, err := toml.DecodeFile("configs/apiserver.toml", config)
	if err != nil {
		log.Fatal(err)
	}
	s := app.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
