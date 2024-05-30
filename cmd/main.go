package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/kevalsabhani/go-ecom/cmd/api"
	"github.com/kevalsabhani/go-ecom/config"
	"github.com/kevalsabhani/go-ecom/db"
)

func main() {
	db, err := db.NewMySqlStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Net:                  "tcp",
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}
	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
