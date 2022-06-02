package main

import (
	"Go_REST_API/internal/app/apiserver"
	"flag"
	"github.com/BurntSushi/toml"
	"log"
)

/*
netsh advfirewall firewall add rule name="Postgre Port" dir=in action=allow protocol=TCP localport=5432

createdb -U postgres testdb
*/

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")

}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
