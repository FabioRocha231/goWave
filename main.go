package main

import (
	fabinHttp "go_http/infra/http"
	"log"
	"net/http"
)

func main() {
	server := fabinHttp.NewApp()

	// server.Use(func(w http.ResponseWriter, r *http.Request) {
	// 	log.Println("OLHA O MIDDLEWARE")
	// })

	server.GET("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("Teste testando!")); err != nil {
			log.Println("Deu erro lek")
		}
	})

	server.GET("/teste", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("Opa fion")); err != nil {
			log.Println("Deu erro lek")
		}
	})

	log.Fatal(server.ListAndServe(":8080"))
}
