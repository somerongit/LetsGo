package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {

	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Printf("Could not load .env file")
		os.Exit(1)
	}
	envMap, mapErr := godotenv.Read(".env")
	if mapErr != nil {
		fmt.Printf("Error loading .env into map[string]string\n")
		os.Exit(1)
	}
	fmt.Println(envMap["PORT"])
	fmt.Printf(os.Getenv("PORT"))
}
