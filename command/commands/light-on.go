package command

import (
	"github.com/MGYOSBEL/design-patterns/command/devices"
)

type LightOnCommand struct {
	light devices.Light
}

func NewLightOnCommand(light devices.Light) Command {
	return &LightOnCommand{
		light: light,
	}
}

func (cmd LightOnCommand) Execute() {
	cmd.light.On()
}
