package void

import (
	"fmt"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

var routesRegexp = regexp.MustCompile(`^.*/cmd/[^/]+/commands(/.*)?\.go$`)

type commandView interface {
	Name() string
	Path() []string
	Tagline() string
	Description() string
}

type Command struct {
	name        string
	path        []string
	tagline     string
	description string
	enabled     bool
	flags       []*Flag
	arguments   []*Argument
	execute     func(ctx *Context) error
	preExecute  func(ctx *Context) error
	postExecute func(ctx *Context) error
}

func NewCommand() *Command {
	path, name := mustResolveCommandPath()

	return &Command{
		name:        name,
		path:        path,
		tagline:     "",
		description: "",
		enabled:     true,
		flags:       make([]*Flag, 0),
		arguments:   make([]*Argument, 0),
		execute:     nil,
		preExecute:  nil,
		postExecute: nil,
	}
}

func (command *Command) Name() string {
	return command.name
}

func (command *Command) Path() []string {
	return command.path
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

func (command *Command) SetArguments(arguments ...*Argument) *Command {
	command.arguments = arguments
	return command
}

func (command *Command) SetFlags(flags ...*Flag) *Command {
	command.flags = flags
	return command
}

func (command *Command) MarkDisabled() *Command {
	command.enabled = false
	return command
}

func (command *Command) HandleExecute(execute func(ctx *Context) error) *Command {
	command.execute = execute
	return command
}

func (command *Command) HandlePreExecute(preExecute func(ctx *Context) error) *Command {
	command.preExecute = preExecute
	return command
}

func (command *Command) HandlePostExecute(postExecute func(ctx *Context) error) *Command {
	command.postExecute = postExecute
	return command
}

func (command *Command) getPath() string {
	return strings.Join(command.path, " ")
}

func (command *Command) getArguments() ([]*Argument, []*Argument, *Argument) {
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

func (command *Command) getFlags() ([]*Flag, []*Flag) {
	required := make([]*Flag, 0)
	optional := make([]*Flag, 0)

	for i := range command.flags {
		if command.flags[i].required {
			required = append(required, command.flags[i])
			continue
		}

		optional = append(optional, command.flags[i])
	}

	return required, optional
}

func (command *Command) mustValidate() {
	if len(command.path) == 0 {
		panic(
			NewDefinitionError(CommandPathRequired, nil).
				WithCommand(command).
				WithMessage(
					"Command path is required.",
				).
				WithExpected(
					"Every command must define a valid command path.",
				),
		)
	}

	if len(command.name) == 0 {
		panic(
			NewDefinitionError(CommandNameRequired, nil).
				WithCommand(command).
				WithMessage(
					"Command name is required.",
				).
				WithExpected(
					"Every command must define a command name.",
				),
		)
	}

	if len(command.tagline) == 0 {
		panic(
			NewDefinitionError(CommandTaglineRequired, nil).
				WithCommand(command).
				WithMessage(
					"Command tagline is required.",
				).
				WithExpected(
					fmt.Sprintf("Command '%s' must define a tagline.", command.name),
				),
		)
	}

	if len(command.description) == 0 {
		panic(
			NewDefinitionError(CommandDescriptionRequired, nil).
				WithCommand(command).
				WithMessage(
					"Command description is required.",
				).
				WithExpected(
					fmt.Sprintf("Command '%s' must define a description.", command.name),
				),
		)
	}

	if command.execute == nil {
		panic(
			NewDefinitionError(CommandExecuteHandlerRequired, nil).
				WithCommand(command).
				WithMessage(
					"Command execution handler is required.",
				).
				WithExpected(
					"Every command must define an execution handler using HandleExecute().",
				).
				WithSuggestion(
					"Attach an execution handler to the command definition.",
				),
		)
	}
}

func mustResolveCommandPath() ([]string, string) {
	route := mustResolveCommandRoute()

	root := strings.Split(filepath.Dir(route), " ")
	base, _, _ := strings.Cut(filepath.Base(route), ".")

	if len(root) == 0 || root[0] == "." {
		return []string{base}, base
	}

	if base == "root" {
		return root, root[len(root)-1]
	}

	root = append(root, base)

	return root, base
}

func mustResolveCommandRoute() string {
	workPath, err := filepath.Abs(".")

	if err != nil {
		NewFrameworkError(CommandRouteInvalid, err).
			WithMessage(
				"Command route resolution failed.",
			).
			WithExpected(
				"Command files must exist within '/cmd/[app]/commands/*'.",
			).
			WithSuggestion(
				"Ensure the command file is placed inside the command routes directory.",
			).
			Fatal()
	}

	_, filePath, _, ok := runtime.Caller(3)

	if !ok {
		NewFrameworkError(CommandRouteInvalid, nil).
			WithMessage(
				"Command route resolution failed.",
			).
			WithExpected(
				"Command files must exist within '/cmd/[app]/commands/*'.",
			).
			WithSuggestion(
				"Ensure the command file is placed inside the command routes directory.",
			).
			Fatal()
	}

	if !routesRegexp.MatchString(filePath) {
		NewFrameworkError(CommandRouteInvalid, nil).
			WithMessage(
				"Invalid command route detected.",
			).
			WithExpected(
				"Command files must exist within '/cmd/[app]/commands/*'.",
			).
			WithReceived(
				fmt.Sprintf("Resolved invalid command file path: %s", filePath),
			).
			WithSuggestion(
				"Ensure the command file is placed inside the command routes directory.",
			).
			Fatal()
	}

	relFilePath, err := filepath.Rel(workPath, filepath.Clean(filePath))

	if err != nil {
		NewFrameworkError(CommandRouteInvalid, err).
			WithMessage(
				"Failed to resolve command route.",
			).
			WithExpected(
				"Commands must be defined within '/cmd/[app]/commands/*'.",
			).
			WithReceived(
				fmt.Sprintf("Resolved file path '%s' is outside the command routes directory.", filePath),
			).
			WithSuggestion(
				"Move the command file into '/cmd/[app]/commands/*'.",
			).
			Fatal()
	}

	route := strings.Join(strings.Split(relFilePath, "\\commands\\")[1:], "")

	return route
}
