package main

import (
	"errors"
	"fmt"
)

type Product struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}

type Inventory struct {
	prodmap map[int]Product
}

func NewInventory() *Inventory {
	return &Inventory{
		prodmap: make(map[int]Product),
	}
}

func (i *Inventory) AddProduct(p Product) {
	i.prodmap[p.ID] = p
}

func (i *Inventory) SellProduct(id, qnt int) error {
	product, ok := i.prodmap[id]
	if !ok {
		return errors.New("НЕТУ")
	}
	if product.Quantity < qnt {
		return errors.New("ИДИ НАХУЙ")
	}
	product.Quantity = -qnt
	i.prodmap[id] = product
	return nil

}

func main() {
	inv := NewInventory()
	inv.AddProduct(Product{ID: 1, Name: "Телефон", Price: 30000, Quantity: 10})
	fmt.Println(inv.SellProduct(1, 3))
	fmt.Println(inv.SellProduct(1, 8))
	fmt.Println(inv.SellProduct(2, 6))
}
