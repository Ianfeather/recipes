package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"recipes/internal/pkg/app"
	"recipes/internal/pkg/common"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "recipe_app@tcp(127.0.0.1:3306)/shoppinglist")

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

	router, err := application.GetRouter()
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
