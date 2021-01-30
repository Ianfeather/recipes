# Recipes Go API

## Running the app

In the root of the app:
- run `./env`
- visit http://localhost:8080/health
- visit http://localhost:8080/recipe


## deploying
yeah, it sucks for now - v manual

```
1. ./build.sh
2. https://console.aws.amazon.com/lambda/home?region=us-east-1#/functions/recipes?newFunction=true&tab=configuration
3. Actions > Upload a zip file
```
## Progress:
it can't find the function name. possible it never could
- logs: https://console.aws.amazon.com/cloudwatch/home?region=us-east-1#logsV2:log-groups/log-group/$252Faws$252Flambda$252Frecipes/log-events/2021$252F01$252F24$252F$255B$2524LATEST$255D5a2fbf5bdbdf411eb3a72623392ff42f
- endpoint: https://pleeyu7yrd.execute-api.us-east-1.amazonaws.com/test
