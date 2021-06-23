package main

import (
	"fmt"

	bvg "github.com/MGYOSBEL/design-patterns/template-method/caffeine-beverage"
	"github.com/MGYOSBEL/design-patterns/template-method/coffee"
	"github.com/MGYOSBEL/design-patterns/template-method/tea"
)

func main() {
	beverageMachine := new(bvg.CaffeineBeverage)

	c := new(coffee.Coffee)
	beverageMachine.Beverage = c
	beverageMachine.PrepareRecipe()

	fmt.Println("")

	t := new(tea.Tea)
	beverageMachine.Beverage = t
	beverageMachine.PrepareRecipe()
}
