package main

import (
	"fmt"
	"log"

	rule34 "github.com/Momgoloid69/rule34-go/client"
)

func main() {
	userID := "" // input your user ID here
	apiKey := "" // input your API key here
	rule34 := rule34.New(userID, apiKey)

	posts, err := rule34.Posts().Find()
	if err != nil {
		log.Fatalf("internal error: %v", err)
	}

	fmt.Println(posts)
}
