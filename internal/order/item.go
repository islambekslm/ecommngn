package order

import "sync"

type (
	Item struct {
		ID       string
		Quantity int
		Product  Product
	}
	ItemList struct {
		list map[string]*Item
		mu   sync.RWMutex
	}
)

func NewItemList() ItemList {
	return ItemList{
		list: make(map[string]*Item),
	}
}

func (i *Item) Price() float64 {
	return i.Product.Price * float64(i.Quantity)
}

type Product struct {
	ID    string
	Price float64
}
