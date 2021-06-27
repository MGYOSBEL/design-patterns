package main

import (
	"fmt"

	diner "github.com/MGYOSBEL/design-patterns/iterator/diner-menu"
	pancake "github.com/MGYOSBEL/design-patterns/iterator/pancake-house"
)

func main() {
	fmt.Println("Hello")
}

func initPancakeHouseMenu() {
	pancakeMenu := pancake.NewPancakeHouseMenu()
	pancakeMenu.AddItem("K&B's Pancake Breakfast", "Pancakes with scrambled eggs, and toast", true, 2.99)
	pancakeMenu.AddItem("Regular Pancake Breakfast", "Pancakes with fried eggs, sausage", false, 2.99)
	pancakeMenu.AddItem("Blueberry Pancakes", "Pancakes mades with fresh blueberries", true, 3.49)
	pancakeMenu.AddItem("Wafles", "Wafles, with your choice of blueberries or strawberries", true, 3.59)
}

func initDinerMenu() {
	dinerMenu := diner.NewDinerMenu()
	dinerMenu.AddItem("Vegetarian BLT", "(Fakin') Bacon with lettuce & tomato on whole wheat", true, 2.99)
	dinerMenu.AddItem("BLT", "Bacon with lettuce & tomato on whole wheat", false, 2.99)
	dinerMenu.AddItem("Soup of the day", "Soup of the day, with a side of potato salad", false, 3.29)
	dinerMenu.AddItem("Hotdog", "A hot dog, with saurkraut, relish, onions, topped with cheese", false, 3.05)
}
