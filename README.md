# rule34-go

[![Go Report Card](https://goreportcard.com/badge/github.com/Momgoloid69/rule34-go)](https://goreportcard.com/report/github.com/Momgoloid69/rule34-go)
[![GoDoc](https://godoc.org/github.com/Momgoloid69/rule34-go?status.svg)](https://godoc.org/github.com/Momgoloid69/rule34-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

`rule34-go` is an unofficial Go client library for interacting with the `rule34.xxx` API. It provides a simple and fluent interface for searching and retrieving posts.

## ⚠️ Disclaimer

This library is a client for an API that serves Not Safe For Work (NSFW) / adult content. By using this library, you acknowledge that you are of legal age to view such content in your jurisdiction and take full responsibility for its use. The author of this library is not responsible for how it is used.

## Features

-   Easy-to-use fluent [builder pattern](https://en.wikipedia.org/wiki/Builder_pattern) for post searching.
-   Strongly-typed helpers for API parameters like ratings, sorting, and filtering.
-   Search by tags, ID, score, rating, and more.
-   Blacklist tags from search results.
-   Sort results by various fields.
-   Built-in support for JSON response parsing.

## Installation

To install the library, use `go get`:

```bash
go get github.com/Momgoloid69/rule34-go
```

## Authentication

The rule34.xxx API requires a `userID` and `apiKey` for authenticated requests. You can obtain these from your account options page:

[https://rule34.xxx/index.php?page=account&s=options](https://rule34.xxx/index.php?page=account&s=options)

Then, create a new client instance:

```go
import rule34 "github.com/Momgoloid69/rule34-go/rule34"

func main() {
    userID := "YOUR_USER_ID"
    apiKey := "YOUR_API_KEY"

    client := rule34.New(userID, apiKey)

    // ... use the client to make requests
}
```

**Note:** It is strongly recommended not to hardcode your credentials in source code. Use environment variables or other secure methods to manage your keys.

## Usage

### Basic Example: Get Latest Posts

This is the simplest way to fetch posts, using the API's default values.

```go
package main

import (
	"fmt"
	"log"

	rule34 "github.com/Momgoloid69/rule34-go/rule34"
)

func main() {
	userID := "YOUR_USER_ID"
	apiKey := "YOUR_API_KEY"
	client := rule34.New(userID, apiKey)

	// Execute a simple request to find posts with default options
	posts, err := client.Posts().Find()
	if err != nil {
		log.Fatalf("Error finding posts: %v", err)
	}

    // Print the ID of the first post, if any
	if len(posts) > 0 {
		fmt.Printf("Found first post ID: %d\n", posts[0].ID)
        fmt.Printf("File URL: %s\n", posts[0].FileURL)
	} else {
		fmt.Println("No posts found.")
	}
}
```

### Advanced Example: Using the Query Builder

The query builder allows you to construct complex queries by chaining methods.

```go
package main

import (
	"fmt"
	"log"

	rule34 "github.com/Momgoloid69/rule34-go/rule34"
	"github.com/Momgoloid69/rule34-go/rule34/filtering"
	"github.com/Momgoloid69/rule34-go/rule34/operators"
	"github.com/Momgoloid69/rule34-go/rule34/rating"
	"github.com/Momgoloid69/rule34-go/rule34/sorting"
)

func main() {
	userID := "YOUR_USER_ID"
	apiKey := "YOUR_API_KEY"
	client := rule34.New(userID, apiKey)

	// Find 10 posts tagged with 'breasts' and 'looking_at_viewer',
	// excluding the 'solo' tag, with a 'questionable' rating,
	// a score of >= 50, sorted by score in descending order.
	posts, err := client.Posts().
		Limit(10).
		Tags("breasts", "looking_at_viewer").
		BlackList("solo").
		Rating(rating.Questionable).
		Where(filtering.Score, operators.GreaterEqual, 50).
		SortBy(sorting.Score).Desc().
		Find()

	if err != nil {
		log.Fatalf("Error finding posts: %v", err)
	}

	fmt.Printf("Found %d posts matching the criteria.\n", len(posts))
	for _, post := range posts {
		fmt.Printf("ID: %d, Score: %d, Tags: %s\n", post.ID, post.Score, post.Tags)
	}
}
```

## API Coverage

The following API endpoints are supported:

-   [X] **Posts** (`page=dapi&s=post&q=index`)
-   [ ] **Deleted Images** (`...&deleted=show`)
-   [ ] **Comments** (`page=dapi&s=comment&q=index`)
-   [ ] **Tags** (`page=dapi&s=tag&q=index`)
-   [ ] **Autocomplete** (`autocomplete.php`)

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for bug reports, feature requests, or questions.

## License

This project is licensed under the MIT License - see the `LICENSE` file for details.