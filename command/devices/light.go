package devices

import "fmt"

type Light struct {
	state bool
	name  string
}

func NewLight(name string) Light {
	return Light{
		name:  name,
		state: false,
	}
}

func (lgt Light) On() {
	lgt.state = true
	fmt.Printf("The %s light is on\n", lgt.name)
}

func (lgt Light) Off() {
	lgt.state = false
	fmt.Printf("The %s light is off\n", lgt.name)
}
