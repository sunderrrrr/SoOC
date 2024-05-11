package main

import (
	"SoCC/pkg/api"
	"fmt"
	"log"
	"net/http"
)

func runServer() {
	api := api.New("localhost:8080", http.NewServeMux())
	api.FillEndpoints()
	fmt.Println("Сервер запущен на порту 8080")
	log.Fatal(api.ListenAndServe())
}

func main() {
	runServer()
}
