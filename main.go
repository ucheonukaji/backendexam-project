package main

import (
	"backendexam-project/product"
	"backendexam-project/store"
)
 
func main() {
	//create a new product

	product1 :=product.NewProduct("toyota","red",100, 500000, "avalon")
	product2 :=product.NewProduct("honda","blue",50, 900000, "iris")

	store :=store.NewStore("uche's store")
	store.AddItem(product1)
	store.AddItem(product2)

	store.ListItems()

	store.Sellitem(product1.Name(), 2)

	store.ShowSoldItems()

	product1.Status()

	store.AddItem(product2)

	//number of cars left to sell(quantities&product)
	store.ListItemsCount()

//sum of prices of cars left
	store.ItemsTotalPrice()

	//numbers of cars sold
	store.ShowSoldItemsCount()

	//sum total of cars prices sold
	store.ShowSoldItemsPrice()

	//list of orders of sales he made
	store.ShowSoldItems()



}

