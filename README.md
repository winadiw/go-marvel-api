# go-marvel-api

Go based Marvel API

Go version 1.17

## Dependencies
 - [fiber](https://github.com/gofiber/fiber) 
 - [godotenv](https://github.com/joho/godotenv)  
 - [Marvel API](https://developer.marvel.com/)

## Documentation

OpenAPI documentation can be found in the [swagger.yaml](./swagger.yaml) file  
Or run the server first, then get from:
 - YML: `http://localhost:8080/swagger.yml`
 - UI: `http://localhost:8080/docs`

Documentation is created using Library:
 - [go-swagger](https://www.gorillatoolkit.org/) 
 - [redoc](https://github.com/Redocly/redoc)  
 - [redoc-cli](https://redoc.ly/docs/redoc/quickstart/cli/)

Test swagger.yml:  
`swagger validate ./swagger.yml`

Run swagger:  
`swagger serve ./swagger.yml`

Build `redoc-static.html`:  
`redoc-cli bundle ./swagger.yml`

Using `make`:  
`make swagger-build`

## Prerequisite
Before running, make sure to create and populate the content of `.env` file to get environment variables.  
Use `.env.sample` as an example for the env file
```
HOST=127.0.0.1:8080
ENV=development
REDIS_ADDR=localhost:6379
REDIS_PASS=
REDIS_DB=0
MARVEL_BASE_URL=https://gateway.marvel.com:443
MARVEL_PRIVATE_KEY=
MARVEL_PUBLIC_KEY=
```

## Running

The application can be run with `go run`

```
➜ go run main.go

curl localhost:8080/characters
```

Or if you prefer docker, use Docker compose  
running: `docker-compose up`
daemon mode: `docker-compose up -d` or `make docker-up`
stop docker: `docker-compose down` or `make docker-down` 

## Testing
Unit Test
```
go test -v -short ./...

# to check coverage
go test -v -short ./... -coverprofile=coverage.out && go tool cover -html=coverage.out
```

Using `make`:  
`make unit-test` or `make integration-test`

## VSCode Enable Debugging
### launch.json for Mac/Linux/Windows
```
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch",
            "type": "go",
            "request": "launch",
            "port": 8080,
            "host": "127.0.0.1",
            "mode": "debug",
            "program": "${workspaceRoot}",
            "showLog": true,
            "debugAdapter": "legacy"
        }
    ]
}
```