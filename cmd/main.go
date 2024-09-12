package main

import (
	//inbuilt
	"fmt"
	"log"
	"os"

	//3rd party

	"github.com/joho/godotenv"

	//local
	"github.com/arhantbararia/ecom_api/cmd/api"
	"github.com/arhantbararia/ecom_api/db"
)

func main() {
	fmt.Println("Starting Server")
	fmt.Println("Loading ENV file")

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)

	}

	//initialize database engine
	mysql_db := db.MySQlConnection{
		User:    getEnv("DB_USER", "root"),
		Passwd:  getEnv("DB_PASSWORD", ""),
		HOST:    getEnv("DB_HOST", "127.0.0.1"),
		PORT:    getEnv("DB_PORT", "3306"), //its a mysql connection afterall
		DB_NAME: getEnv("DB_NAME", "ecom_go"),
	}

	mysql_db_cxn, err := mysql_db.Connect()

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	db.CheckDB(mysql_db_cxn)

	server := api.NewAPIServer(":8000", mysql_db_cxn)

	if err := server.Run(); err != nil {
		log.Fatal(err)

	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
