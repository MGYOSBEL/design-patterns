package tea

import (
	"fmt"

	bvg "github.com/MGYOSBEL/design-patterns/template-method/caffeine-beverage"
)

type Tea struct {
	beverage bvg.CaffeineBeverage
}

func (*Tea) Brew() {
	fmt.Println("Steeping the Tea ")
}

func (*Tea) AddCondiments() {
	fmt.Println("Adding lemon")
}
