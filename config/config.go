package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

var (
	// Get current file full path from runtime
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	ProjectRootPath = filepath.Join(filepath.Dir(b), "../")
)

// Config func to get env value
func Config(key string) string {
	// load .env file
	err := godotenv.Load(ProjectRootPath + "/.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	return os.Getenv(key)
}
