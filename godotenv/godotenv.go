package godotenv

import (
	"os"
	
	"github.com/joho/godotenv"
)

func New() Godotenv {
	return &Params{}
}

type Setup func() Godotenv

func (params *Params) Load() string {
	if err := godotenv.Load(".env"); err != nil {
		panic("failed to load .env: " + err.Error())
		
		return ""
	}
	
	appMode := os.Getenv("APP_MODE")
	
	if appMode == "" {
		panic("APP MODE environment variable is not set")
		
		return ""
	}
	
	return appMode
}
