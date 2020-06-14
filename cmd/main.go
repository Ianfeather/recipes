package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"recipes/internal/pkg/app"
	"recipes/internal/pkg/common"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	pass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	db, err := sql.Open("mysql", fmt.Sprintf("%s@tcp(%s:3306)/shoppinglist", pass, dbHost))

	if err != nil {
		fmt.Println("Failed to connect to database")
		panic(err.Error())
	}

	env := &common.Env{DB: db}

	application, err := app.NewApp(env)

	if err != nil {
		fmt.Println("Failed to create application")
		fmt.Println(err)
	}

	router, err := application.GetRouter("")
	if err != nil {
		fmt.Println("Failed to get application router")
		fmt.Println(err)
	}

	server := http.Server{
		Addr:         ":8080",
		ReadTimeout:  3000 * time.Millisecond,
		WriteTimeout: 3000 * time.Millisecond,
		Handler:      router,
	}

	server.ListenAndServe()
	defer db.Close()
}
