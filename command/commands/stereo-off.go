package command

import (
	"github.com/MGYOSBEL/design-patterns/command/devices"
)

type StereoOffCommand struct {
	stereo devices.Stereo
}

func NewStereoOffCommand(stereo devices.Stereo) Command {
	return &StereoOffCommand{
		stereo: stereo,
	}
}

func (cmd StereoOffCommand) Execute() {
	cmd.stereo.Off()
}
