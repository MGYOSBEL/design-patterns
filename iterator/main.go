package main

import (
	"github.com/MGYOSBEL/design-patterns/iterator/cafe"
	"github.com/MGYOSBEL/design-patterns/iterator/diner"
	"github.com/MGYOSBEL/design-patterns/iterator/pancake"
	"github.com/MGYOSBEL/design-patterns/iterator/waitress"
)

func main() {
	pancakeMenu := pancake.InitPancakeHouseMenu()
	dinerMenu := diner.InitDinerMenu()
	cafeMenu := cafe.InitCafeMenu()

	wts := waitress.NewWaitress(dinerMenu, pancakeMenu, cafeMenu)
	wts.PrintMenu()
}
