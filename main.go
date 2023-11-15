package main

import (
	"HouseOfPets/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc(
		"/cats",
		handlers.GetAllCats,
	)

	http.HandleFunc(
		"/dogs",
		handlers.GetAllDogs,
	)

	http.HandleFunc(
		"/add-cat",
		handlers.InsertCat,
	)

	http.HandleFunc(
		"/hello",
		handlers.Hello,
	)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
