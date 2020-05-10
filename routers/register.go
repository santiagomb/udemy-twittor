package routers

import (
	"encoding/json"
	"github.com/santiagomb/udemy-twittor/database"
	"github.com/santiagomb/udemy-twittor/models"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user *models.User
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "Invalid email.", http.StatusBadRequest)
		return
	}

	if len(user.Password) < 6 {
		http.Error(w, "Password is too short. At least length 7 is required.", http.StatusBadRequest)
		return
	}

	object, err := database.GetUser(user.Email)
	if object != nil {
		http.Error(w, "User already registered. Please try another email.", http.StatusBadRequest)
		return
	}

	err = database.SaveUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	/*if !ok {
		http.Error(w, fmt.Sprintf("Error while saving user %s", user.Email), http.StatusInternalServerError)
		return
	}*/

	w.WriteHeader(http.StatusCreated)
}
