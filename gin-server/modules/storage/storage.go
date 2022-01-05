package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	DB, err := sql.Open("mysql", "root:toor@tcp(localhost:4123)/CSE")
	// key
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
