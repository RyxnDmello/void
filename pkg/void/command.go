package void

import (
	"strings"
)

type Command struct {
	path        []string
	name        string
	tagline     string
	enabled     bool
	flags       []string
	arguments   []string
	execute     func(context *Context, arguments, flags []string) error
	preExecute  func(context *Context, arguments, flags []string) error
	postExecute func(context *Context, arguments, flags []string) error
}

func NewPath(path string) ([]string, string) {
	newPath := make([]string, 0)

	parts := strings.SplitSeq(path, " ")

	for part := range parts {
		part = strings.ToLower(part)
		part = strings.Trim(part, " ")
		newPath = append(newPath, part)
	}

	newName := strings.Join(newPath, " ")

	return newPath, newName
}

func NewCommand(path string) *Command {
	newPath, newName := NewPath(path)

	return &Command{
		path:      newPath,
		name:      newName,
		tagline:   "A command",
		enabled:   true,
		flags:     nil,
		arguments: nil,
		execute: func(context *Context, arguments, flags []string) error {
			return nil
		},
		preExecute: func(context *Context, arguments, flags []string) error {
			return nil
		},
		postExecute: func(context *Context, arguments, flags []string) error {
			return nil
		},
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

func (command *Command) Enabled() bool {
	return command.enabled
}

func (command *Command) Disable() *Command {
	command.enabled = false
	return command
}

func (command *Command) Flags() []string {
	return command.flags
}

func (command *Command) SetFlags(flags []string) *Command {
	command.flags = flags
	return command
}

func (command *Command) Arguments() []string {
	return command.arguments
}

func (command *Command) SetArguments(arguments []string) *Command {
	command.arguments = arguments
	return command
}

func (command *Command) Execute(execute func(context *Context, arguments, flags []string) error) *Command {
	command.execute = execute
	return command
}

func (command *Command) PreExecute(preExecute func(context *Context, arguments, flags []string) error) *Command {
	command.preExecute = preExecute
	return command
}

func (command *Command) PostExecute(postExecute func(context *Context, arguments, flags []string) error) *Command {
	command.postExecute = postExecute
	return command
}
