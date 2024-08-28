package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mwaitforitosama/TodoApp/routes"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	app := &routes.App{}

	app.InitializeDB()
	defer app.Db.Close()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("Failed to start the server: ", err)
	}
}
