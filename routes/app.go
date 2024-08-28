package routes

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
)

type App struct {
	router http.Handler
	Db     *sql.DB
}

func New() *App {
	router := loadRoutes()
	app := &App{
		router: router,
	}
	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	fmt.Println("Starting server on :3000")
	fmt.Println(server.Handler)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Server failed: %v", err)
		return fmt.Errorf("error: %v", err)
	}

	fmt.Println("Server is running")
	return nil
}

func (a *App) InitializeDB() {
	fmt.Println("i am here")
	connStr := buildConnStr()

	var err error
	a.Db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	err = a.Db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}
	fmt.Println("Successfully Connected")
}

func buildConnStr() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST"),
		os.Getenv("PORT"),
		os.Getenv("USER"),
		os.Getenv("PASSWORD"),
		os.Getenv("DB"),
	)
}
