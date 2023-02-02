package entities

type Shoe struct {
	Id          int     `json:"id"`
	Image       string  `json:"image"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Color       string  `json:"color"`
}

type Shoes struct {
	Shoes []Shoe `json:"shoes"`
}
