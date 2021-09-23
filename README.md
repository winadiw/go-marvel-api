# go-marvel-api

Go based Marvel API

## Documentation

OpenAPI documentation can be found in the [swagger.yaml](./swagger.yaml) file
`http://localhost:8080`  
Or can be accessed with UI on `http://localhost:8080/docs`

## Running

The appliction can be run with `go run`

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