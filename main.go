package main

import (
	"fmt"
	"log"
	"os"

	rule34 "github.com/Momgoloid/rule34-go/client"
	"github.com/joho/godotenv"
)

func main() {
	userID, apiKey := MustLoad()
	rule34 := rule34.New(userID, apiKey)

	post, err := rule34.Posts().Find()
	if err != nil {
		log.Fatalf("internal error: %v", err)
	}

	fmt.Println(post.Posts)
}

func MustLoad() (string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env file: %v\n", err)
	}

	userID := os.Getenv("USER_ID")
	if userID == "" {
		log.Fatal("Variable USER_ID is not set")
	}
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("Variable API_KEY is not set")
	}

	return userID, apiKey
}
