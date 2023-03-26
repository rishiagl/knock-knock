package main

import (
	"fmt"
	"log"

	"github.com/R58235/knock-knock/internal/repositories"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("./pkg/env/.env")
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
}