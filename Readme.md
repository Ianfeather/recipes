# Recipes Go API

## Running the app

In the root of the app:
- run `./env`
- visit http://localhost:8080/health
- visit http://localhost:8080/recipe


# Deploying
```
go get -v all

GOOS=linux go build -o build/main cmd/main.go

zip -jrm build/main.zip build/main
```
