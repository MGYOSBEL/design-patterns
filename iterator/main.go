package main

import (
	"fmt"

	"github.com/MGYOSBEL/design-patterns/iterator/diner"
	"github.com/MGYOSBEL/design-patterns/iterator/pancake"
)

func main() {
	pancakeMenu := pancake.InitPancakeHouseMenu()
	pancakeItems := pancakeMenu.GetMenuItems()

	fmt.Println("-----PANCAKE HOUSE MENU-----")
	for _, item := range pancakeItems {
		fmt.Printf("%s \t-- %v \t -- %s\n", item.GetName(), item.GetPrice(), item.GetDescription())
	}
	fmt.Println("")

	dinerMenu := diner.InitDinerMenu()
	dinerItems := dinerMenu.GetMenuItems()

	fmt.Println("-----DINER MENU-----")
	for i := 0; i < dinerMenu.GetNumberOfItems(); i++ {
		fmt.Printf("%s \t-- %v \t -- %s\n", dinerItems[i].GetName(), dinerItems[i].GetPrice(), dinerItems[i].GetDescription())
	}
	fmt.Println("")

}
