package main

import (
	"log"

	"github.com/0x-pankaj/ecom/cmd/api"
	"github.com/0x-pankaj/ecom/config"
	"github.com/0x-pankaj/ecom/db"
	"github.com/go-sql-driver/mysql"
)

func main() {

	db, err := db.NewSQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal("Error in db config: ", err)
	}

	server := api.NewAPIServer(":9090", db)

	if err := server.Run(); err != nil {
		log.Fatal("Error while running server: ", err)
	}
}
