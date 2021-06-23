package control

import (
	"fmt"
	"reflect"

	command "github.com/MGYOSBEL/design-patterns/command/commands"
)

type RemoteControl struct {
	onCommands  []command.Command
	offCommands []command.Command
}

func NewRemoteControl() RemoteControl {
	onCommands := make([]command.Command, 7)
	offCommands := make([]command.Command, 7)

	for idx := range onCommands {
		onCommands[idx] = command.NewNoCommand()
		offCommands[idx] = command.NewNoCommand()
	}
	return RemoteControl{
		onCommands:  onCommands,
		offCommands: offCommands,
	}
}

func (rc *RemoteControl) SetCommand(slot int, onCommand command.Command, offCommand command.Command) {
	rc.offCommands[slot] = offCommand
	rc.onCommands[slot] = onCommand
}

func (rc *RemoteControl) OnButtonWasPushed(slot int) {
	rc.onCommands[slot].Execute()
}

func (rc *RemoteControl) OffButtonWasPushed(slot int) {
	rc.offCommands[slot].Execute()
}

func (rc *RemoteControl) Print() {
	fmt.Println("--------Remote Control---------")
	for i := range rc.offCommands {
		on_cmd := reflect.TypeOf(rc.onCommands[i])
		off_cmd := reflect.TypeOf(rc.offCommands[i])
		slot_str := fmt.Sprintf("[slot %d]: %v %v", i, on_cmd, off_cmd)
		fmt.Println(slot_str)
	}
}
