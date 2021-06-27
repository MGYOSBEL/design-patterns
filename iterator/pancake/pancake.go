package pancake

import (
	"github.com/MGYOSBEL/design-patterns/iterator/menu"
)

type PancakeHouseMenu struct {
	menuItems []*menu.MenuItem
}

func NewPancakeHouseMenu() *PancakeHouseMenu {
	menuItems := make([]*menu.MenuItem, 0)

	return &PancakeHouseMenu{
		menuItems: menuItems,
	}
}

func (diner *PancakeHouseMenu) AddItem(name string, description string, vegetarian bool, price float64) {
	itm := menu.NewMenuItem(name, description, vegetarian, price)

	diner.menuItems = append(diner.menuItems, itm)
}

func (menu *PancakeHouseMenu) GetMenuItems() []*menu.MenuItem {
	return menu.menuItems
}

func InitPancakeHouseMenu() *PancakeHouseMenu {
	pancakeMenu := NewPancakeHouseMenu()
	pancakeMenu.AddItem("K&B's Pancake Breakfast", "Pancakes with scrambled eggs, and toast", true, 2.99)
	pancakeMenu.AddItem("Regular Pancake Breakfast", "Pancakes with fried eggs, sausage", false, 2.99)
	pancakeMenu.AddItem("Blueberry Pancakes", "Pancakes mades with fresh blueberries", true, 3.49)
	pancakeMenu.AddItem("Wafles", "Wafles, with your choice of blueberries or strawberries", true, 3.59)
	return pancakeMenu
}
