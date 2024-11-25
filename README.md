# GoWave

GoWave is a lightweight HTTP framework built in Go, inspired by the simplicity and structure of ExpressJS. It aims to provide a minimalistic approach for handling HTTP requests and building web servers, using the standard Go HTTP library.

## Features

- Lightweight and minimal framework.
- Simple routing with HTTP methods (GET, POST, etc.).
- Ability to define custom middleware (currently under development).
- Easy-to-use handler functions.
- Built on top of Go's standard `net/http` package.

## Installation

You can include the `goWave` framework in your Go project by using the following command:

```bash
go get github.com/yourusername/goWave
```

## Usage
# Example Application
```go
package main

import (
	fabinHttp "goWave"
	"log"
	"net/http"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Teste testando!")); err != nil {
		log.Println("Error occurred:", err)
	}
}

func opaHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Opa fion")); err != nil {
		log.Println("Error occurred:", err)
	}
}

func main() {
	// Initialize the server
	server := fabinHttp.NewApp()

	// Define routes and handlers
	server.GET("/", testHandler)
	server.GET("/teste", opaHandler)

	// Start the server on port 8080
	log.Fatal(server.ListAndServe(":8080"))
}
```

## Explanation

- Initialization:
- goWave.NewApp() initializes the application, creating a new instance of the server.
# Routing:
- You define routes using methods like server.GET("/path", handler). Each route is linked to a specific handler function that processes the HTTP request.
# Handlers:
- Handlers are simple Go functions that accept an http.ResponseWriter and http.Request as arguments. In this example, testHandler and opaHandler are defined to handle requests at / and /teste, respectively.
# Starting the server:
- server.ListAndServe(":8080") starts the HTTP server on the specified port (8080).

# Middleware (Under Development)
The framework allows for adding middleware to modify the request-response cycle. Although middleware functionality is in progress, you can envision it being implemented like this:
```go
server.Use(func(w http.ResponseWriter, r *http.Request) {
	log.Println("This is a middleware")
})
```

## Roadmap
- Implement custom middleware support.
- Add support for more HTTP methods (POST, PUT, DELETE, etc.).
- Improve error handling and logging utilities.
- Enhance routing capabilities (dynamic routes, query parameters, etc.).

Steps to contribute:
1 - Fork this repository.
2 - Create a new branch.
3 - Make your changes.
4 - Run tests (once added).
5 - Submit a pull request.

## License
GoWave is licensed under the MIT License. See the LICENSE file for more details.

## Contact
For any questions, open an issue or contact the repository owner.
