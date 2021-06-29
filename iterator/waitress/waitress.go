package waitress

import (
	"fmt"

	"github.com/MGYOSBEL/design-patterns/iterator/cafe"
	"github.com/MGYOSBEL/design-patterns/iterator/diner"
	"github.com/MGYOSBEL/design-patterns/iterator/iterator"
	"github.com/MGYOSBEL/design-patterns/iterator/pancake"
)

type Waitress struct {
	dinerMenu   *diner.DinerMenu
	pancakeMenu *pancake.PancakeHouseMenu
	cafeMenu    *cafe.CafeMenu
}

func NewWaitress(dinerMenu *diner.DinerMenu, pancakeMenu *pancake.PancakeHouseMenu, cafeMenu *cafe.CafeMenu) Waitress {
	return Waitress{
		dinerMenu:   dinerMenu,
		pancakeMenu: pancakeMenu,
		cafeMenu:    cafeMenu,
	}
}

func (wts Waitress) PrintMenu() {
	dinerIterator := wts.dinerMenu.CreateIterator()
	pancakeIterator := wts.pancakeMenu.CreateIterator()
	cafeIterator := wts.cafeMenu.CreateIterator()

	fmt.Printf("%60s\n", "----- CAFE MENU -----")
	printMenu(cafeIterator)
	fmt.Printf("%60s\n", "----- PANCAKE MENU -----")
	printMenu(pancakeIterator)
	fmt.Printf("%60s\n", "----- DINER MENU -----")
	printMenu(dinerIterator)
}

func printMenu(it iterator.Iterator) {
	for it.HasNext() {
		item := it.Next()
		fmt.Printf("%30s \t-- %v \t -- %s\n", item.GetName(), item.GetPrice(), item.GetDescription())
	}
	fmt.Println("")
}
