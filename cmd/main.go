package main

import (
	//inbuilt
	"log"

	//3rd party
	"github.com/joho/godotenv/cmd/godotenv"


	//local
	"github.com/arhantbararia/ecom_api/cmd/api"
	"github.com/arhantbararia/ecom_api/db"

)


func main() {
	fmt.Println("Starting Server")
	fmt.Println("Loading ENV file")
	
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %v" , err )

	}


	



	//initialize database engine
	db , err := db.MySqlConnection(mysql.Config{
		User:
	})

	server := api.NewAPIServer(":8000")

	if err := server.Run(); err != nil {
		log.Fatal(err)

	}
}