package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/khanhhung142/GoldenOwlTest/Backend/entities"
)

var shoes = make([]*entities.Shoe, 0)

func MakeAListShoes() {
	jsonFile, err := os.Open("../data/shoes.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &shoes)
}

func GetAllShoes() []*entities.Shoe {
	MakeAListShoes()
	return shoes
}
