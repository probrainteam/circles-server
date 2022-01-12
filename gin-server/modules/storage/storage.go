package storage

import (
	. "circlesServer/modules/reader"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	path := GetConfig(`db.ID`) + ":" + GetConfig(`db.PW`) + "@tcp(localhost:" + GetConfig(`db.PORT`) + ")/" + GetConfig(`db.DB`)
	DB, err := sql.Open("mysql", path)
	if err != nil {
		panic(err)
	}
	db = DB
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)
}
func DB() *sql.DB {
	return db
}
