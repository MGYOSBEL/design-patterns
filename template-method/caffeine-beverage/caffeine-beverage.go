package caffeine_beverage

import "fmt"

type ICaffeineBeverage interface {
	Brew()
	AddCondiments()
}

type CaffeineBeverage struct {
	Beverage ICaffeineBeverage
}

func (bvg *CaffeineBeverage) PrepareRecipe() {
	bvg.BoilWater()
	bvg.Beverage.Brew()
	bvg.PourInCup()
	bvg.Beverage.AddCondiments()
}

func (bvg *CaffeineBeverage) BoilWater() {
	fmt.Println("Boiling water")
}

func (bvg *CaffeineBeverage) PourInCup() {
	fmt.Println("Pouring into cup")
}
