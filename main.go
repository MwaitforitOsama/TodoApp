package main

import (
	"context"
	"fmt"

	"github.com/mwaitforitosama/TodoApp/routes"
)

func main() {
	app := routes.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start the server : ", err)
	}

}
