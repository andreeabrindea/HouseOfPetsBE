package handlers

import (
	"HouseOfPets/db"
	"encoding/json"
	"net/http"
)

func enableCors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

func CorsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(w, r)
		next(w, r)
	}
}

func OptionsHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
}

func GetAllCats(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r)
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	cats, err := db.GetAllCats("postgres://nlgxpwcp:iamp04QJAP13wvmkXU3ngCaco0MrQeop@flora.db.elephantsql.com/nlgxpwcp")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	catsJS, _ := json.MarshalIndent(cats, "", "  ")
	_, err = w.Write(catsJS)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r)
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	catsJS, _ := json.MarshalIndent("hello there", "", "  ")
	_, err := w.Write(catsJS)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func InsertCat(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r)
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	var cat db.Cat
	err := json.NewDecoder(r.Body).Decode(&cat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err, _ = db.InsertCat(cat, "postgres://nlgxpwcp:iamp04QJAP13wvmkXU3ngCaco0MrQeop@flora.db.elephantsql.com/nlgxpwcp")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := "Cat added"
	jsonResp, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResp)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
}
