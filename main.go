package main

import (
	"fmt"
	fabinHttp "go_http/infra/http"
	"html"
	"log"
	"net/http"
)

func main() {
	server := fabinHttp.NewApp()

	server.GET("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(server.ListAndServe(":8080"))
}
