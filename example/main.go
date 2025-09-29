// Package main provides a simple example of how to use the rule34-go client library.
package main

import (
	"fmt"
	"log"

	"github.com/Momgoloid69/rule34-go/rule34"
)

// main is the entry point for the example application.
// It demonstrates how to initialize the client and make a basic API call
// to fetch the latest posts.
func main() {
	// IMPORTANT: Replace with your actual user ID and API key.
	// You can obtain these from your account options page on the rule34 website.
	userID := "" // input your user ID here
	apiKey := "" // input your API key here
	rule34 := rule34.New(userID, apiKey)

	// Make a simple request to find posts with default options.
	posts, err := rule34.Posts().Find()
	if err != nil {
		log.Fatalf("internal error: %v", err)
	}

	// Print the retrieved posts.
	fmt.Println(posts)
}
