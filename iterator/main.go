package main

import (
	"fmt"

	"github.com/MGYOSBEL/design-patterns/iterator/diner"
	"github.com/MGYOSBEL/design-patterns/iterator/iterator"
	"github.com/MGYOSBEL/design-patterns/iterator/pancake"
)

func main() {
	pancakeMenu := pancake.InitPancakeHouseMenu()
	pancakeIterator := pancakeMenu.CreateIterator()

	fmt.Printf("%60s\n", "----- PANCAKE MENU -----")
	printMenu(pancakeIterator)

	dinerMenu := diner.InitDinerMenu()
	dinerIterator := dinerMenu.CreateIterator()

	fmt.Printf("%60s\n", "----- DINER MENU -----")
	printMenu(dinerIterator)

}

func printMenu(it iterator.Iterator) {
	for it.HasNext() {
		item := it.Next()
		fmt.Printf("%30s \t-- %v \t -- %s\n", item.GetName(), item.GetPrice(), item.GetDescription())
	}
	fmt.Println("\v")
}
