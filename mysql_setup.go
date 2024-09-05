package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func createDatabase() error {

	//load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//   s3Bucket := os.Getenv("S3_BUCKET")
	//   secretKey := os.Getenv("SECRET_KEY")
	MYSQL_USER := getEnv("DB_USER", "default_user")
	MYSQL_PASSWORD := getEnv("DB_PASSWORD", "")
	MYSQL_HOST := getEnv("DB_HOST", "localhost")
	MYSQL_PORT := getEnv("DB_PORT", "3306")
	DB_NAME := getEnv("DB_NAME", "default_db")

	// Data Source Name: <username>:<password>@tcp(<hostname>:<port>)
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/", MYSQL_USER, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_PORT)
	fmt.Println(dsn)
	// Open a connection to the MySQL server
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("could not open db connection: %v", err)
	}
	defer db.Close()

	// Create the database
	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", DB_NAME)
	_, err = db.Exec(query)
	if err != nil {
		return fmt.Errorf("could not open db connection: %v", err)
	}

	return nil
}

func main() {
	fmt.Println("Creating Database")
	err := createDatabase()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database created successfully")
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
