package data

import (
	"encoding/json"
	"errors"
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

func RemoveProducts(id string) bool {

	var index int
	var found bool
	for i, p := range products {
		if p.Id == id {
			index = i
			found = true
			break
		}
	}

	if !found {
		return false
	}
	
	products = append(products[:index], products[index+1:]...)
	return found
}

func AddProducts(p *Product){
	products = append(products, p)
}

func UpdateProducts(p *Product, id string) error {

	for _, product := range products{
		if product.Id == id{
			product.Name = p.Name
			product.Description = p.Description
			product.Price = p.Price
			return nil
		}
	}

	return errors.New("Product doesnt exist")
}

type Products []*Product

func (p *Products) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}
