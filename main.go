package main

import (
	"fmt"
	"log"
	"main/db"
	"main/routes"
	"main/services"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Microservice beta")

	if err := godotenv.Load(); err != nil {
		log.Printf("no .env file loaded: %v", err)
	}

	client, err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	services.SetClient(client)

	routes.Servers(client)
	services.FeedBacktaker()
}

//strech