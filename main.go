package main

import (
	"github.com/santiagomb/udemy-twittor/bd"
	"github.com/santiagomb/udemy-twittor/handlers"
	"log"
)

func main() {
	success := bd.Ping()
	if success {
		log.Println("Starting server ...")
		handlers.Handlers()
	}
}
