package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MySQlConnection struct {
	User    string
	Passwd  string
	HOST    string
	PORT    string
	DB_NAME string
}

func (msql *MySQlConnection) Connect() (*sql.DB, error) {

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/", msql.User, msql.Passwd, msql.HOST, msql.PORT)
	//dsn string with DB_NAME
	dsn = fmt.Sprintf("%v%v?parseTime=true", dsn, msql.DB_NAME)
	// Open a connection to the MySQL server
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("could not open db connection: %v", err)
	}
	return db, err
}

func CheckDB(db *sql.DB) {
	log.Println("Checking Database Connection")
	err := db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	log.Println("Connection to DB Successfull")

}
