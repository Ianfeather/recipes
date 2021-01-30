package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"recipes/internal/pkg/app"
	"recipes/internal/pkg/common"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"

	_ "github.com/go-sql-driver/mysql"
)

var muxLambda *gorillamux.GorillaMuxAdapter

func init() {
	fmt.Println("init()")
	pass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	fmt.Println("init(): 1")
	fmt.Println(fmt.Sprintf("admin:%s@tcp(%s:3306)/bigshop", pass, dbHost))
	db, err := sql.Open("mysql", fmt.Sprintf("admin:%s@tcp(%s:3306)/bigshop", pass, dbHost))
	fmt.Println("init(): 2")

	if err != nil {
		fmt.Println("Failed to connect to database")
		panic(err.Error())
	}
	fmt.Println(db)

	env := &common.Env{DB: db}

	application, err := app.NewApp(env)

	if err != nil {
		fmt.Println("Failed to create application")
		fmt.Println(err)
	}

	r, err := application.GetRouter("")
	if err != nil {
		fmt.Println("Failed to get application router")
		fmt.Println(err)
	}

	muxLambda = gorillamux.New(r)
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("in handler")
	fmt.Println(req.Path)
	return muxLambda.ProxyWithContext(ctx, req)
}

func main() {
	fmt.Println("starting main()")
	lambda.Start(handler)
}
