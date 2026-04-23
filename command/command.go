package command

import (
	"strings"
)

type Command struct {
	path        []string
	name        string
	tagline     string
	description string
	enabled     bool
	flags       []string
	arguments   []*Argument
	execute     func(path string, cmd *Command, args map[string]*Argument, flags map[string]string) error
	preExecute  func(path string, cmd *Command, args map[string]*Argument, flags map[string]string) error
	postExecute func(path string, cmd *Command, args map[string]*Argument, flags map[string]string) error
}

func NewCommand() *Command {
	path, name := resolvePath()

	return &Command{
		path:        path,
		name:        name,
		tagline:     "A command.",
		description: "A command.",
		enabled:     true,
		flags:       make([]string, 0),
		arguments:   make([]*Argument, 0),
		execute:     nil,
		preExecute:  nil,
		postExecute: nil,
	}
}

func (command *Command) Path() string {
	return strings.Join(command.path, " ")
}

func (command *Command) Name() string {
	return command.name
}

func (command *Command) SetName(name string) *Command {
	command.name = name
	return command
}

func (command *Command) Tagline() string {
	return command.tagline
}

func (command *Command) SetTagline(tagline string) *Command {
	command.tagline = tagline
	return command
}

func (command *Command) Description() string {
	return command.description
}

func (command *Command) SetDescription(description string) *Command {
	command.description = description
	return command
}

func (command *Command) Disable() *Command {
	command.enabled = false
	return command
}

func (command *Command) Flags() []string {
	return command.flags
}

func (command *Command) AddFlag(flag string) *Command {
	command.flags = append(command.flags, flag)
	return command
}

func (command *Command) Arguments() ([]*Argument, []*Argument) {
	required := make([]*Argument, 0)
	optional := make([]*Argument, 0)

	for i := range command.arguments {
		if command.arguments[i].required {
			required = append(required, command.arguments[i])
		} else {
			optional = append(optional, command.arguments[i])
		}
	}

	return required, optional
}

func (command *Command) SetArguments(arguments []*Argument) *Command {
	required := make([]*Argument, 0)
	optional := make([]*Argument, 0)

	for i := range arguments {
		if arguments[i].required {
			required = append(required, arguments[i])
		} else {
			optional = append(optional, arguments[i])
		}
	}

	command.arguments = append(command.arguments, required...)
	command.arguments = append(command.arguments, optional...)

	return command
}

func (command *Command) Execute() func(path string, cmd *Command, args map[string]*Argument, flags map[string]string) error {
	return command.execute
}

func (command *Command) HandleExecute(execute func(path string, cmd *Command, args map[string]*Argument, flags map[string]string) error) *Command {
	if execute == nil {
		panic("Command Execute Handler Must Be Provided")
	}

	command.execute = execute
	return command
}

func (command *Command) PreExecute() func(path string, cmd *Command, args map[string]*Argument, flags map[string]string) error {
	return command.preExecute
}

func (command *Command) HandlePreExecute(preExecute func(path string, cmd *Command, args map[string]*Argument, flags map[string]string) error) *Command {
	command.preExecute = preExecute
	return command
}

func (command *Command) PostExecute() func(path string, cmd *Command, args map[string]*Argument, flags map[string]string) error {
	return command.postExecute
}

func (command *Command) HandlePostExecute(postExecute func(path string, cmd *Command, args map[string]*Argument, flags map[string]string) error) *Command {
	command.postExecute = postExecute
	return command
}
