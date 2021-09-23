# go-marvel-api

Go based Marvel API

Go version 1.17

## Dependencies
 - [fiber](https://github.com/gofiber/fiber) 
 - [godotenv](https://github.com/joho/godotenv)  
 - [Marvel API](https://developer.marvel.com/)

## Documentation

OpenAPI documentation can be found in the [swagger.yaml](./swagger.yaml) file
`http://localhost:8080`  
Or can be accessed with UI on `http://localhost:8080/docs`  

Documentation is created using Library:
 - [go-swagger](https://www.gorillatoolkit.org/) 
 - [redoc](https://github.com/Redocly/redoc)  
 - [redoc-cli](https://www.npmjs.com/package/redoc-cli)

Test swagger.yml:  
`swagger validate ./swagger.yml`

Run swagger:  
`swagger serve ./swagger.yml`

Build `redoc-static.html`:  
`redoc-cli bundle ./swagger.yml`

## Prerequisite
Before running, make sure to create `.env` file to get environment variables.  
Use `.env.sample` as an example for the env file

## Running

The application can be run with `go run`

```
➜ go run main.go

curl localhost:8080/characters
```

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