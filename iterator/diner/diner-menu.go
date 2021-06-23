package diner

import (
	"fmt"

	item "github.com/MGYOSBEL/design-patterns/iterator/menu_item"
)

const MAX_ITEMS = 6

type DinerMenu struct {
	numberOfItems int
	menuItems     [MAX_ITEMS]*item.MenuItem
}

func NewDinerMenu() *DinerMenu {
	var menuItems [MAX_ITEMS]*item.MenuItem
	return &DinerMenu{
		numberOfItems: 0,
		menuItems:     menuItems,
	}
}

func (menu *DinerMenu) AddItem(name string, description string, vegetarian bool, price float64) {
	if menu.numberOfItems == MAX_ITEMS {
		fmt.Println("Max Number of items reached. Can't be added more items.")
	} else {
		itm := item.NewItem(name, description, vegetarian, price)
		menu.menuItems[menu.numberOfItems] = &itm
	}
}

func (menu *DinerMenu) GetMenuItems() [MAX_ITEMS]*item.MenuItem {
	return menu.menuItems
}
