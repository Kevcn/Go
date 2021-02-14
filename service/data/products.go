package data

import (
	"encoding/json"
	"io"
)

type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	SKU         string  `json:"sku" validate:"sku"`
}

func GetProducts() Products {
	return productList
}

// --- For post endpoint ---
func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	// throws error if fail to decode the input
	return e.Decode(p)
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	// not very nice way of adding an element...hmmm
	productList = append(productList, p)
}

func getNextID() int64 {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

// --- For get endpoint ---
// Custom type so we can have a method on it
type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	// Newencoder writes to whatever the io is passed in!
	e := json.NewEncoder(w)
	return e.Encode(p)
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
	},
	&Product{
		ID:          2,
		Name:        "Esspresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
	},
}
