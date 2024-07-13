package order

type Order struct {
	ID     string
	UserID string
	Items  ItemList
}

func NewOrder(id, userID string) *Order {
	return &Order{
		ID:     id,
		UserID: userID,
		Items:  NewItemList(),
	}
}

func (o *Order) Price() float64 {
	var sum float64
	o.Items.mu.RLock()
	for _, item := range o.Items.list {
		sum += item.Price()
	}
	o.Items.mu.RUnlock()
	return sum
}

func (o *Order) AddItem(item Item) {
	o.Items.mu.Lock()
	o.Items.list[item.ID] = &item
	o.Items.mu.Unlock()
}

func (o *Order) RemoveItem(id string) {
	o.Items.mu.Lock()
	delete(o.Items.list, id)
	o.Items.mu.Unlock()
}

func (o *Order) UpdateItem(id string, quantity int) {
	o.Items.mu.Lock()
	if item, ok := o.Items.list[id]; ok {
		item.Quantity = quantity
	}
	o.Items.mu.Unlock()
}
