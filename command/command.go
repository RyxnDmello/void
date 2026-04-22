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
	arguments   []string
	execute     func(path string, cmd *Command, args, flags []string) error
	preExecute  func(path string, cmd *Command, args, flags []string) error
	postExecute func(path string, cmd *Command, args, flags []string) error
}

func NewCommand() *Command {
	path, name := resolveCommandPath()

	return &Command{
		path:        path,
		name:        name,
		tagline:     "A command.",
		description: "A command.",
		enabled:     true,
		flags:       make([]string, 0),
		arguments:   make([]string, 0),
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

func (command *Command) Arguments() []string {
	return command.arguments
}

func (command *Command) AddArgument(argument string) *Command {
	command.arguments = append(command.arguments, argument)
	return command
}

func (command *Command) Execute() func(path string, cmd *Command, args, flags []string) error {
	return command.execute
}

func (command *Command) HandleExecute(execute func(path string, cmd *Command, args, flags []string) error) *Command {
	command.execute = execute
	return command
}

func (command *Command) PreExecute() func(path string, cmd *Command, args, flags []string) error {
	return command.preExecute
}

func (command *Command) HandlePreExecute(preExecute func(path string, cmd *Command, args, flags []string) error) *Command {
	command.preExecute = preExecute
	return command
}

func (command *Command) PostExecute() func(path string, cmd *Command, args, flags []string) error {
	return command.postExecute
}

func (command *Command) HandlePostExecute(postExecute func(path string, cmd *Command, args, flags []string) error) *Command {
	command.postExecute = postExecute
	return command
}
