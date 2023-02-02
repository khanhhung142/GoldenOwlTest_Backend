package main

import (
	"net/http"

	"github.com/gorilla/mux"
	shoeapi "github.com/khanhhung142/GoldenOwlTest_Backend/apis/shoeAPI"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/shoeapi/getall", shoeapi.GetAllShoes).Methods("GET")
	err := http.ListenAndServe(":5000", router)
	if err != nil {
		panic(err)
	}
}
