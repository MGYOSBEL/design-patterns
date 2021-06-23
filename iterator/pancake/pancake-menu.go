package pancake_house

import (
	item "github.com/MGYOSBEL/design-patterns/iterator/menu_item"
)

type PancakeHouseMenu struct {
	menuItems []*item.MenuItem
}

func NewPancakeHouseMenu() *PancakeHouseMenu {
	menuItems := make(item.MenuItem, 10)

	return &PancakeHouseMenu{
		menuItems: menuItems,
	}
}

func (menu *PancakeHouseMenu) AddItem(name string, description string, vegetarian bool, price float64) {
	itm := item.NewMenuItem(name, description, vegetarian, price)

	menu.menuItems = append(menu.menuItems, itm)
}

func (menu *PancakeHouseMenu) GetMenuItems() []*item.MenuItem {
	return menu.menuItems
}
