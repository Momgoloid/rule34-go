package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	rule34 "github.com/Momgoloid/rule34-go/client"
	"github.com/joho/godotenv"
)

func main() {
	userID, apiKey := MustLoad()

	rule34 := rule34.New(userID, apiKey)

	post, err := rule34.GetPost(1)
	if err != nil {
		log.Fatalf("internal error: %v", err)
	}

	// _ = post

	fmt.Println(post)
}

func MustLoad() (int, string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env file: %v\n", err)
	}

	userIDStr := os.Getenv("USER_ID")
	if userIDStr == "" {
		log.Fatal("Variable USER_ID is not set")
	}
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("Variable API_KEY is not set")
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		log.Fatalf("Failed to convert string to int: %v", err)
	}

	return userID, apiKey
}
