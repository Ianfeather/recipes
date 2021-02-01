# Recipes Go API

## Running the app

In the root of the app:
- run `./env`
- visit http://localhost:8080/health
- visit http://localhost:8080/recipes

This uses a regular server locally and a lambda outside of dev.

### local setup
- db user
```
CREATE USER 'admin'@'localhost' IDENTIFIED BY 'admin';
GRANT ALL PRIVILEGES ON bigshop.* TO 'admin'@'localhost';
```

## testing the app
```
go fmt ./...
go test ./... -v
```

## deploying
Master builds are automatically deployed to AWS via circle-ci ([dashboard](https://app.circleci.com/pipelines/github/Ianfeather))

Local builds can be deployed manually with the following:
```
1. ./build.sh
2. https://console.aws.amazon.com/lambda/home?region=us-east-1#/functions/recipes?newFunction=true&tab=configuration
3. Actions > Upload a zip file
```

## useful links
- [endpoint](https://pleeyu7yrd.execute-api.us-east-1.amazonaws.com/prod)
- [lambda](https://console.aws.amazon.com/lambda/home?region=us-east-1#/functions/recipes)
- [api gateway](https://console.aws.amazon.com/apigateway/home?region=us-east-1#/apis/pleeyu7yrd/stages/prod)
- [cloudwatch logs](https://console.aws.amazon.com/cloudwatch/home?region=us-east-1#logsV2:log-groups/log-group/$252Faws$252Flambda$252Frecipes)
- [RDS database](https://console.aws.amazon.com/rds/home?region=us-east-1#database:id=big-shop;is-cluster=false;tab=maintenance-and-backups)
