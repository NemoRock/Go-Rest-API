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


migrate -path migrations -database "postgresql://localhost/restapi_dev?sslmode=disable" up

password=11111

scram-sha-256

migrate create -ext sql -dir migrations create_users

migrate -path migrations -database "postgres://localhost/restapi_dev?user=postgres password=11111 sslmode=disable" up

database_url = "host=localhost user=postgres dbname=restapi_dev password=11111 sslmode=disable"
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
