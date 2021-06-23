package main

import (
	"fmt"

	command "github.com/MGYOSBEL/design-patterns/command/commands"
	"github.com/MGYOSBEL/design-patterns/command/control"
	"github.com/MGYOSBEL/design-patterns/command/devices"
)

func main() {
	fmt.Println("TESTING THE COMMAND PATTERN... GO!")

	remoteControl := control.NewRemoteControl()

	bedRoomLight := devices.NewLight("bedroom")
	diningRoomLight := devices.NewLight("diningRoom")
	stereo := devices.NewStereo()

	bedroomLightOnCommand := command.NewLightOnCommand(bedRoomLight)
	bedroomLightOffCommand := command.NewLightOffCommand(bedRoomLight)

	diningRoomLightOnCommand := command.NewLightOnCommand(diningRoomLight)
	diningRoomLightOffCommand := command.NewLightOffCommand(diningRoomLight)

	stereoOnCommand := command.NewStereoOnWithCDCommand(stereo)
	stereoOffCommand := command.NewStereoOffCommand(stereo)

	remoteControl.SetCommand(0, bedroomLightOnCommand, bedroomLightOffCommand)
	remoteControl.SetCommand(1, diningRoomLightOnCommand, diningRoomLightOffCommand)
	remoteControl.SetCommand(2, stereoOnCommand, stereoOffCommand)

	remoteControl.OnButtonWasPushed(0)
	remoteControl.OffButtonWasPushed(0)
	remoteControl.OnButtonWasPushed(1)
	remoteControl.OffButtonWasPushed(1)
	remoteControl.OnButtonWasPushed(2)
	remoteControl.OffButtonWasPushed(2)
}
