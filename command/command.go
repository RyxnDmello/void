package command

import (
	"fmt"
	"slices"
	"strings"

	"github.com/ryxndmello/void/internal/context"
)

type Command struct {
	use         string
	name        string
	tagline     string
	description string
	enabled     bool
	path        []string
	flags       []string
	arguments   []*Argument
	execute     func(ctx *context.Context) error
	preExecute  func(ctx *context.Context) error
	postExecute func(ctx *context.Context) error
}

func NewCommand() *Command {
	path, name := resolvePath()

	return &Command{
		use:         name,
		name:        name,
		tagline:     "",
		description: "",
		enabled:     true,
		path:        path,
		flags:       make([]string, 0),
		arguments:   make([]*Argument, 0),
		execute:     nil,
		preExecute:  nil,
		postExecute: nil,
	}
}

func (command *Command) Use() string {
	return command.use
}

func (command *Command) Name() string {
	return command.name
}

func (command *Command) Tagline() string {
	return command.tagline
}

func (command *Command) SetTagline(tagline string) *Command {
	command.tagline = strings.Trim(tagline, " ")
	return command
}

func (command *Command) Description() string {
	return command.description
}

func (command *Command) SetDescription(description string) *Command {
	command.description = strings.Trim(description, " ")
	return command
}

func (command *Command) IsEnabled() bool {
	return command.enabled
}

func (command *Command) Disable() *Command {
	command.enabled = false
	return command
}

func (command *Command) Path() string {
	return strings.Join(command.path, " ")
}

func (command *Command) Flags() []string {
	return command.flags
}

func (command *Command) SetFlags(flags ...string) *Command {
	command.flags = flags
	return command
}

func (command *Command) Arguments() ([]*Argument, []*Argument, *Argument) {
	var variadic *Argument

	required := make([]*Argument, 0)
	optional := make([]*Argument, 0)

	for i := range command.arguments {
		if command.arguments[i].variadic {
			variadic = command.arguments[i]
			break
		}

		if command.arguments[i].required {
			required = append(required, command.arguments[i])
		} else {
			optional = append(optional, command.arguments[i])
		}
	}

	return required, optional, variadic
}

func (command *Command) SetArguments(arguments ...*Argument) *Command {
	var variadic *Argument

	names := make([]string, 0)
	required := make([]*Argument, 0)
	optional := make([]*Argument, 0)

	for i := range arguments {
		arguments[i].MustValidate()

		if slices.Contains(names, arguments[i].name) {
			panic(fmt.Sprintf("Command '%s' Must Have Unique Arguments", command.name))
		}

		if arguments[i].variadic {
			if variadic != nil {
				panic(fmt.Sprintf("Command '%s' Already Contains A Variadic Argument '%s'", command.name, variadic.name))
			}

			variadic = arguments[i]
			continue
		}

		if arguments[i].required {
			required = append(required, arguments[i])
		} else {
			optional = append(optional, arguments[i])
		}

		names = append(names, arguments[i].name)
	}

	command.arguments = append(command.arguments, required...)
	command.arguments = append(command.arguments, optional...)

	if variadic != nil {
		command.arguments = append(command.arguments, variadic)
	}

	return command
}

func (command *Command) Execute(ctx *context.Context) error {
	if ctx == nil {
		panic("Command Execution Handler Must Have A Context")
	}

	if command.execute == nil {
		panic("Command Execution Handler Must Be Provided")
	}

	return command.execute(ctx)
}

func (command *Command) SetExecute(execute func(ctx *context.Context) error) *Command {
	if execute == nil {
		panic("Command Execution Handler Must Be Provided")
	}

	command.execute = execute
	return command
}

func (command *Command) PreExecute(ctx *context.Context) error {
	if ctx == nil {
		panic("Command Pre Execution Handler Must Have A Context")
	}

	if command.preExecute == nil {
		return nil
	}

	return command.preExecute(ctx)
}

func (command *Command) SetPreExecute(preExecute func(ctx *context.Context) error) *Command {
	command.preExecute = preExecute
	return command
}

func (command *Command) PostExecute(ctx *context.Context) error {
	if ctx == nil {
		panic("Command Post Execution Handler Must Have A Context")
	}

	if command.postExecute == nil {
		return nil
	}

	return command.postExecute(ctx)
}

func (command *Command) SetPostExecute(postExecute func(ctx *context.Context) error) *Command {
	command.postExecute = postExecute
	return command
}

func (command *Command) MustValidate() {
	if len(command.path) == 0 {
		panic("Every Command Must Have A Path")
	}

	if len(command.name) == 0 {
		panic("Every Command Must Have A Name")
	}

	if len(command.tagline) == 0 {
		panic(fmt.Sprintf("Command '%s' Must Have A Tagline", command.name))
	}

	if len(command.description) == 0 {
		panic(fmt.Sprintf("Command '%s' Must Have A Description", command.name))
	}

	if command.execute == nil {
		panic(fmt.Sprintf("Command '%s' Must Have An Execution Handler", command.name))
	}
}
