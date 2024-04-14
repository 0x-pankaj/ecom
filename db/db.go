package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

// var connStr = "mysql://root:secret@localhost:3306/mysql-database"

// db, err := sql.Open("mysql", connStr)
// if err != nil {
// 	log.Fatal("Error while connection to database: ", err)
// }

// defer db.Close()

func NewSQLStorage(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal("error from sql connection: ", err)
	}

	return db, nil
}
