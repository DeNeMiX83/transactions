package main

import (
	"api/app"
	"fmt"
    "github.com/joho/godotenv"
    "log"
)

func init() {
    if err := godotenv.Load("deploy/.env"); err != nil {
	    log.Print("No .env file found")
    }
}


func main() {
    fmt.Println("Starting")
    app := app.Application{}
    app.Inicialize()
    fmt.Println("Initialized")
}