package command

import (
	"github.com/MGYOSBEL/design-patterns/command/devices"
)

type LightOffCommand struct {
	light devices.Light
}

func NewLightOffCommand(light devices.Light) Command {
	return &LightOffCommand{
		light: light,
	}
}

func (cmd LightOffCommand) Execute() {
	cmd.light.Off()
}
