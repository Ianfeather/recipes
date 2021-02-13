package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"recipes/internal/pkg/app"
	"recipes/internal/pkg/common"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/codegangsta/negroni"
	_ "github.com/go-sql-driver/mysql"
)

var muxLambda *gorillamux.GorillaMuxAdapter
var router *negroni.Negroni

func init() {
	pass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	db, err := sql.Open("mysql", fmt.Sprintf("admin:%s@tcp(%s:3306)/bigshop", pass, dbHost))

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

	router, err = application.GetRouter("")
	if err != nil {
		fmt.Println("Failed to get application router")
		fmt.Println(err)
	}

	// muxLambda = gorillamux.New(router)
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return muxLambda.ProxyWithContext(ctx, req)
}

func main() {
	args := os.Args

	if len(args) > 1 && args[1] == "dev" {
		server := http.Server{
			Addr:         ":8080",
			ReadTimeout:  3000 * time.Millisecond,
			WriteTimeout: 3000 * time.Millisecond,
			Handler:      router,
		}
		server.ListenAndServe()
	} else {
		lambda.Start(handler)
	}
}
