package cafe

import (
	"github.com/MGYOSBEL/design-patterns/iterator/iterator"
	"github.com/MGYOSBEL/design-patterns/iterator/menu"
)

type CafeMenu struct {
	items map[string]*menu.MenuItem
}

func NewCafeMenu() *CafeMenu {
	items := make(map[string]*menu.MenuItem)
	return &CafeMenu{
		items: items,
	}
}

func (cafe *CafeMenu) AddItem(name string, description string, vegetarian bool, price float64) {
	itm := menu.NewMenuItem(name, description, vegetarian, price)
	cafe.items[itm.GetName()] = itm
}

func (cafe *CafeMenu) GetMenuItems() map[string]*menu.MenuItem {
	return cafe.items
}

func (cafe *CafeMenu) GetNumberOfItems() int {
	return len(cafe.items)
}

func (cafe *CafeMenu) CreateIterator() iterator.Iterator {
	return NewCafeMenuIterator(cafe.items)
}

func InitCafeMenu() *CafeMenu {
	cafeMenu := NewCafeMenu()
	cafeMenu.AddItem("Veggie Burger & Air Fries", "Veggie burger on a whole wheat bun, lettuce, tomato and fries", true, 3.99)
	cafeMenu.AddItem("Soup of the day", "A cup of the soup of the day, with a side salad", false, 3.69)
	cafeMenu.AddItem("Burrito", "A large burrito, with a whole pinto beans, salsa, guacamole", true, 4.29)
	return cafeMenu
}
