# go-marvel-api

# VSCode Enable Debugging
## launch.json for Mac/Linux/Windows
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