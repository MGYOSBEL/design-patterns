package main

import (
	"fmt"

	"github.com/MGYOSBEL/design-patterns/iterator/diner"
	"github.com/MGYOSBEL/design-patterns/iterator/pancake"
)

func main() {
	fmt.Println("Hello")
	pancakeMenu := pancake.InitPancakeHouseMenu()
	pancakeItems := pancakeMenu.GetMenuItems()

	for item := range pancakeItems {
		fmt.Printf("%s -- %f -- %s", item.GetName(), item.GetPrice(), item.GetDescription())
	}

	dinerMenu := diner.InitDinerMenu()
	dinerItems := dinerMenu.GetMenuItems()

	for item := range dinerItems {
		fmt.Printf("%s -- %f -- %s", item.GetName(), item.GetPrice(), item.GetDescription())
	}

}
