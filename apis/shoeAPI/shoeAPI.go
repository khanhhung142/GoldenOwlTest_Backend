package shoeapi

import (
	"encoding/json"
	"net/http"

	"github.com/khanhhung142/GoldenOwlTest_Backend/models"
)

func GetAllShoes(response http.ResponseWriter, request *http.Request) {
	shoes := models.GetAllShoes()
	responseWithJSON(response, http.StatusOK, shoes)
}

func responseWithJSON(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}
