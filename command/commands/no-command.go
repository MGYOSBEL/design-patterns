package command

type NoCommand struct{}

func (NoCommand) Execute() {}

func NewNoCommand() Command {
	return &NoCommand{}
}
