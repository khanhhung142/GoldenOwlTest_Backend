package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Printf("ye")
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/user/getall", shoeapi.GetAllShoes).Methods("GET")
	err := http.ListenAndServe(":5000", router)
	if err != nil {
		panic(err)
	}
}
