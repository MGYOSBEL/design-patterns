package menu

type MenuItem struct {
	name        string
	description string
	vegetarian  bool
	price       float64
}

func NewMenuItem(name string, description string, vegetarian bool, price float64) *MenuItem {
	return &MenuItem{
		name:        name,
		description: description,
		vegetarian:  vegetarian,
		price:       price,
	}
}

func (item *MenuItem) GetName() string {
	return item.name
}
func (item *MenuItem) GetDescription() string {
	return item.description
}
func (item *MenuItem) GetPrice() float64 {
	return item.price
}
func (item *MenuItem) IsVegetarian() bool {
	return item.vegetarian
}
