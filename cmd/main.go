package main

import (
	"command-line-argumentsD:\\Codes\\Go\\Microservicw\\todo_by_hand\\server.go"
	"log"
	"todo"
	"todo_app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error uncured while running http server: %s", err.Error())
	}
}
