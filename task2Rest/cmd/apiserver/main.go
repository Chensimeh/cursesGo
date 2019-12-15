package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/trzhensimekh/cursesGo/task2Rest/internal/app"
	"log"
)

var (
	configPath string
)

func init(){
	flag.StringVar(&configPath,"config-path","configs/apiserver.toml","config file path")
}

func main() {
	flag.Parse()
	config := app.NewConfig()
	_ , err := toml.DecodeFile(configPath,config)
	if err !=nil {
		log.Fatal(err)
	}

	s:= app.New(config)
	if err := s.Start(); err!=nil {
		log.Fatal(err)
	}

}
