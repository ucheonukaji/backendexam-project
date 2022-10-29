package product

import (
	"errors"
	"fmt"
)

var _ Product = &product{}

type Product interface {
	Quantity() int
	Price() float32
	Name() string

	Sell(quantity int) error
	Display()
	Status() bool
}

type car struct {
	name  string
	color string
}

type product struct {
	car
	productName string
	quantity    int
	price       float32
}

func NewProduct(carName string, carColor string, quantity int, price float32, productName string) *product {
	return &product{
		car: car{
			name:  carName,
			color: carColor,
		},
		productName: productName,
		quantity:    quantity,
		price:       price,
	}
}

func (p *product) Name() string {
	fmt.Printf("product name: %v", p.name)
	return p.name
}
func (p *product) Display() {
	fmt.Printf("product: %v", p)
}
func (p *product) Quantity() int {
	fmt.Printf("quantity: %v", p.quantity)
	return p.quantity
}
func (p *product) Status() bool {
	if p.quantity == 0 {
		fmt.Println("This product is not available")
		return false
	}

	return true
}
func (p *product) Sell(quantity int) error {
	if p.quantity < quantity {
		return errors.New("not enough products to sell")
	}

	p.quantity = p.quantity - quantity

	return nil
}
func (p *product) Price() float32 {
	return p.price
}
