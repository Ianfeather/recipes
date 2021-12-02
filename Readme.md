# Recipes Go API

Todo list: https://trello.com/b/LnaGkQyG/bigshop

## Running the app

In the root of the app:
- run `./env`
- visit http://localhost:8080/health
- visit http://localhost:8080/recipes

This uses a regular server locally and a lambda outside of dev.

### local setup

To enter the mysql workspace:
```
mysql -u root
use bigshop;
```

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

## runnning db migrations
I haven't created a decent workflow for this yet :(

What I've been doing is switching the RDS instance to be publicly accessible then accessing it via mysql workbench and running the migrations ad-hoc.

I've also been using workbench for dumping the db from prod to local.

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
- [Auth0 (for managing user)](https://manage.auth0.com/dashboard/eu/dev-x-n37k6b/applications/HxkTOH3ZYxjbsgrVI4ii1CV2TQx7hk9G/settings)
