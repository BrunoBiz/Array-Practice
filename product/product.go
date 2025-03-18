package product

import (
	"errors"
)

type Product struct {
	Title string
	Id    int
	Price float32
}

func New(title string, id int, price float32) (Product, error) {

	if title == "" || id == 0 || price == 0 {
		return Product{}, errors.New("invalid input")
	}

	return Product{
		Title: title,
		Id:    id,
		Price: price,
	}, nil
}

// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
