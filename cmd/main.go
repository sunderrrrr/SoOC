package main

import (
	"log"
	sooc "orderalready"
	"orderalready/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(sooc.Server)
	if err := srv.Run("8090", handlers.InitRouts()); err != nil {
		log.Fatalf("error occurupted while running http server: " + err.Error())
	}
}
