package store

import (
	"backendexam-project/product"
	"errors"
	"fmt"
)

type Store interface {
	AddItem(item product.Product)
}

type store struct {
	name      string
	Products  []product.Product //list of products
	SoldItems []product.Product
}

func NewStore(name string) *store {
	return &store{
		name:      name,
		Products:  []product.Product{},
		SoldItems: []product.Product{},
	}
}

// AddItem - add item to the store
func (s *store) AddItem(item product.Product) {
	fmt.Printf("adding product to the store,id: %+v \n\n", item.Name())
	s.Products = append(s.Products, item)
}

// total number of cars available to be sold
func (s *store) ListItemsCount() int {
	fmt.Printf("list items count: %+v \n\n", s.Products)
	var count int
	for _, v := range s.Products {
		count += v.Quantity()
	}
	return count
}

//ListItems - list all product items in the store
func (s *store) ListItems() {
	fmt.Printf("list items: %+v \n\n", s.Products)	
}

// total number of cars available to be sold
func (s *store) ItemsTotalPrice() float32 {
	fmt.Printf("items total price:%+v \n\n", s.Products)
	var price float32
	for _, v := range s.Products {
		price += float32(v.Quantity()) * v.Price()
	}
	return price
}

func (s *store) Sellitem(productName string, quantity int) error {
	var product product.Product
	for _, p := range s.Products {
		if p.Name() == productName {
			product = p
		}
	}
	if product == nil {
		return errors.New("product not found")
	}

	err := product.Sell(quantity)
	if product == nil {
		return err
	}

	s.SoldItems = append(s.SoldItems, product)
	return nil
}

func (s *store) ShowSoldItemsCount() {
	var totalNumberSold int
	for _, v := range s.SoldItems {
		totalNumberSold += v.Quantity()
	}
	fmt.Printf("total number of items sold: %v \n\n", totalNumberSold)
}

//ShowSoldItems - list all product items in the store that have been sold
func (s *store) ShowSoldItems() {
	fmt.Printf("sold items: %+v \n\n", s.SoldItems)	
}

func (s *store) ShowSoldItemsPrice() {
	var totalAmountSold float32
	for _, v := range s.SoldItems {
		totalAmountSold += v.Price() * float32(v.Quantity())
	}
	fmt.Printf("total sold items price: %v \n\n", totalAmountSold)
}
