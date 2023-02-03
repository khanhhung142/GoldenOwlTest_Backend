package main

import (
	"net/http"

	"github.com/gorilla/mux"
	shoeapi "github.com/khanhhung142/GoldenOwlTest_Backend/apis/shoeAPI"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})
	router.HandleFunc("/shoe/getall", shoeapi.GetAllShoes).Methods("GET")

	handler := cors.Handler(router)
	err := http.ListenAndServe(":5000", handler)
	if err != nil {
		panic(err)
	}
}
