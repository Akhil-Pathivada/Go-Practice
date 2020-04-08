package ems

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func openDatabaseConnection() *sql.DB {

	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "codec@rgukt.in"
	dbName := "EMS"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	if err != nil {
		panic(err.Error())
	}
	return db
}
