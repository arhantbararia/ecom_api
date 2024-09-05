package db

import (
	"database/sql"
	"fmt"

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
	// Open a connection to the MySQL server
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("could not open db connection: %v", err)
	}
	return db, err
}
