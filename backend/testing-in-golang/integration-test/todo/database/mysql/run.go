package mysql

import (
	"database/sql"
	"fmt"
	"log"
)

func Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) *sql.DB {
	var err error
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	db, err := sql.Open(Dbdriver, DBURL)
	if err != nil {
		log.Fatal("This is the error connecting to the database:", err)
	}
	fmt.Printf("We are connected to the %s database", Dbdriver)

	return db
}
