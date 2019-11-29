package main

import (
	"flag"
	"github.com/trzhensimekh/cursesGo/task2Rest/internal/app/apiserver"
	"log"
	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init(){
	flag.StringVar(&configPath,"config-path","configs/apiserver.toml","config file path")
}

func main() {
	flag.Parse()
	config :=apiserver.NewConfig()
	_ , err := toml.DecodeFile(configPath,config)
	if err !=nil {
		log.Fatal(err)
	}


	s:=apiserver.New(config)
	if err := s.Start(); err!=nil {
		log.Fatal(err)
	}
	println("hello, apiserver")
}
