package middleware

import (
	"github.com/santiagomb/udemy-twittor/database"
	"net/http"
)

func CheckDataBase(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !database.Ping() {
			http.Error(w, "Lost connection with Mongo DB", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	}
}
