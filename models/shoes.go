package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/khanhhung142/GoldenOwlTest_Backend/entities"
)

var shoes entities.Shoes

func MakeAListShoes() {
	absPathJson, _ := filepath.Abs("./data/shoes.json")
	jsonFile, err := os.Open(absPathJson)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &shoes)
}

func GetAllShoes() entities.Shoes {
	MakeAListShoes()

	return shoes
}
