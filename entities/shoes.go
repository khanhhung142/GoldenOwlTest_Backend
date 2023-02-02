package entities

type Shoe struct {
	id          int
	image       string
	name        string
	description string
	price       float32
	color       string
}

type Shoes struct {
	Shoes []Shoe `json:"shoe"`
}
