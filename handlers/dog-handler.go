package handlers

import (
	"HouseOfPets/db"
	"encoding/json"
	"net/http"
)

func GetAllDogs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	dogs, err := db.GetAllDogs("postgres://nlgxpwcp:iamp04QJAP13wvmkXU3ngCaco0MrQeop@flora.db.elephantsql.com/nlgxpwcp")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	dogsJs, _ := json.MarshalIndent(dogs, "", "  ")
	_, err = w.Write(dogsJs)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
}
