package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/0x-pankaj/ecom/cmd/api"
	"github.com/0x-pankaj/ecom/config"
	"github.com/0x-pankaj/ecom/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Printf(" %s\n", config.Envs.DBPassword)
	password := config.Envs.DBPassword
	println(password)

	db, err := db.NewSQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}) //return database string and err if

	// db, err := db.NewSQLStorage(mysql.Config{
	// 	User:                 "root",
	// 	Passwd:               "secret",
	// 	Addr:                 "http://localhost:3306",
	// 	DBName:               "ecom",
	// 	Net:                  "tcp",
	// 	AllowNativePasswords: true,
	// 	ParseTime:            true,
	// })
	if err != nil {
		log.Fatal("Error in db config: ", err)

	}

	initStorage(db) //check database connection

	server := api.NewAPIServer(fmt.Sprintf(":%s", config.Envs.Port), db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	} //run server on given port with given database string
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: successfully connected")
}
