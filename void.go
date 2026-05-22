package void

type Void struct {
	commands []*Command
}

func NewApp() *Void {
	return &Void{
		commands: make([]*Command, 0),
	}
}

func (void *Void) SetCommands(commands ...*Command) *Void {
	void.commands = commands
	return void
}

func (void *Void) Run() error {
	defer onError()

	command, arguments, flags := parse(void.commands)

	context := newContext(command, arguments, flags)

	return execute(command, context)
}
