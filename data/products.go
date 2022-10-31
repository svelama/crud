package data

import (
	"encoding/json"
	"io"
)

type Product struct{
	Id string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
}

var products = []*Product{
	{
		Id: "1",
		Name: "Pen",
		Description: "Worlds best pen, it can only write on earth",
		Price: 22.75,
	},
	{
		Id: "2",
		Name: "Pencil",
		Description: "Worlds best pen, it can write anywhere",
		Price: 2.75,
	},
}

func GetProducts() Products {
	return products
}

type Products []*Product

func (p *Products) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}