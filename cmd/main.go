package main

import (
	"fmt"
	"net/http"
	"recipes/internal/pkg/app"
	"time"
)

func main() {
	application, err := app.NewApp()

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
}
