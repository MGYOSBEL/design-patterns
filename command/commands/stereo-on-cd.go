package command

import (
	"github.com/MGYOSBEL/design-patterns/command/devices"
)

type StereoOnWithCDCommand struct {
	stereo devices.Stereo
}

func NewStereoOnWithCDCommand(stereo devices.Stereo) Command {
	return &StereoOnWithCDCommand{
		stereo: stereo,
	}
}

func (cmd StereoOnWithCDCommand) Execute() {
	cmd.stereo.On()
	cmd.stereo.SetCD()
	cmd.stereo.SetVoulme(11)
}
