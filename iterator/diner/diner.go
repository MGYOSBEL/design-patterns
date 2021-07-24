package diner

import (
	"fmt"

	"github.com/MGYOSBEL/design-patterns/iterator/iterator"
	"github.com/MGYOSBEL/design-patterns/iterator/menu"
)

const MAX_ITEMS = 6

type DinerMenu struct {
	numberOfItems int
	menuItems     [MAX_ITEMS]*menu.MenuItem
}

func NewDinerMenu() *DinerMenu {
	var menuItems [MAX_ITEMS]*menu.MenuItem
	return &DinerMenu{
		numberOfItems: 0,
		menuItems:     menuItems,
	}
}

func (diner *DinerMenu) AddItem(name string, description string, vegetarian bool, price float64) {
	if diner.numberOfItems == MAX_ITEMS {
		fmt.Println("Max Number of items reached. Can't be added more items.")
	} else {
		itm := menu.NewMenuItem(name, description, vegetarian, price)
		diner.menuItems[diner.numberOfItems] = itm
		diner.numberOfItems++
	}
}

func (menu *DinerMenu) GetMenuItems() [MAX_ITEMS]*menu.MenuItem {
	return menu.menuItems
}

func (menu *DinerMenu) GetNumberOfItems() int {
	return menu.numberOfItems
}

func (menu *DinerMenu) CreateIterator() iterator.Iterator {
	return NewDinerMenuIterator(menu.menuItems)
}

func InitDinerMenu() *DinerMenu {
	dinerMenu := NewDinerMenu()
	dinerMenu.AddItem("Vegetarian BLT", "(Fakin') Bacon with lettuce & tomato on whole wheat", true, 2.99)
	dinerMenu.AddItem("BLT", "Bacon with lettuce & tomato on whole wheat", false, 2.99)
	dinerMenu.AddItem("Soup of the day", "Soup of the day, with a side of potato salad", false, 3.29)
	dinerMenu.AddItem("Hotdog", "A hot dog, with saurkraut, relish, onions, topped with cheese", false, 3.05)

	return dinerMenu
}
