package handlers

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/santiagomb/udemy-twittor/middleware"
	"github.com/santiagomb/udemy-twittor/routers"
	"log"
	"net/http"
	"os"
)

func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middleware.CheckDataBase(routers.Register)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
