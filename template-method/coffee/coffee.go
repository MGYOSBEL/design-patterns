package coffee

import (
	"fmt"

	bvg "github.com/MGYOSBEL/design-patterns/template-method/caffeine-beverage"
)

type Coffee struct {
	beverage bvg.CaffeineBeverage
}

func (*Coffee) Brew() {
	fmt.Println("Dripping coffee through filter")
}

func (*Coffee) AddCondiments() {
	fmt.Println("Adding sugar and milk")
}
