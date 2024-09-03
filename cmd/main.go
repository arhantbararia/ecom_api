package main

import (
	//inbuilt
	"log"

	//3rd party


	//local
	"github.com/arhantbararia/ecom_api/cmd/api"

)


func main() {

	
	server := api.NewAPIServer(":8000")

	if err := server.Run(); err != nil {
		log.Fatal(err)

	}
}