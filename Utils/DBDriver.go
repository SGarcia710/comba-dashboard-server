package Utils

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func OpenDB() *sql.DB {
	fmt.Println("Trying to connect to DB...")
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME))
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected to DB.")
	}
	return db
}
