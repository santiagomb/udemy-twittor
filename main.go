package main

import (
	"github.com/santiagomb/udemy-twittor/database"
	"github.com/santiagomb/udemy-twittor/handlers"
	"log"
)

func main() {
	success := database.Ping()
	if success {
		log.Println("Starting server ...")
		handlers.Handlers()
	}
}
