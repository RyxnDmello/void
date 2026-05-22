package void

import (
	"fmt"
	"path/filepath"
	"regexp"
	"runtime"
	"slices"
	"strings"
)

var routesRegexp = regexp.MustCompile(`^.*/cmd/[^/]+/commands(/.*)?\.go$`)

type Command struct {
	name        string
	tagline     string
	description string
	enabled     bool
	path        []string
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
		tagline:     "",
		description: "",
		enabled:     true,
		path:        path,
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

func (command *Command) Path() []string {
	return command.path
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

func (command *Command) mustNormalizeArguments() *Command {
	var variadic *Argument

	names := make([]string, 0)
	required := make([]*Argument, 0)
	optional := make([]*Argument, 0)

	for i := range command.arguments {
		command.arguments[i].mustValidate()

		if slices.Contains(names, command.arguments[i].name) {
			panic(
				NewDefinitionError(DuplicateCommandArgument, nil).
					WithCommand(command).
					WithArgument(command.arguments[i]).
					WithMessage(
						"Duplicate command arguments detected.",
					).
					WithExpected(
						"Each command argument must have a unique name.",
					).
					WithReceived(
						fmt.Sprintf("Duplicate argument: %s", command.arguments[i].name),
					).
					WithSuggestion(
						"Rename or remove duplicate command arguments.",
					),
			)
		}

		if command.arguments[i].variadic {
			if variadic != nil {
				panic(
					NewDefinitionError(MultipleVariadicArguments, nil).
						WithCommand(command).
						WithArgument(command.arguments[i]).
						WithMessage(
							"Multiple variadic arguments were defined.",
						).
						WithExpected(
							"A command may only define a single variadic argument.",
						).
						WithReceived(
							fmt.Sprintf("Existing variadic argument: %s", variadic.name),
						).
						WithSuggestion(
							"Remove additional variadic arguments from the command definition.",
						),
				)
			}

			variadic = command.arguments[i]
			continue
		}

		if command.arguments[i].required {
			required = append(required, command.arguments[i])
		} else {
			optional = append(optional, command.arguments[i])
		}

		names = append(names, command.arguments[i].name)
	}

	command.arguments = make([]*Argument, 0)
	command.arguments = append(command.arguments, required...)
	command.arguments = append(command.arguments, optional...)

	if variadic != nil {
		command.arguments = append(command.arguments, variadic)
	}

	return command
}

func (command *Command) mustNormalizeFlags() *Command {
	names := make([]string, 0)
	required := make([]*Flag, 0)
	optional := make([]*Flag, 0)

	for i := range command.flags {
		command.flags[i].mustValidate()

		if slices.Contains(names, command.flags[i].name) {
			panic(
				NewDefinitionError(DuplicateCommandFlag, nil).
					WithFlag(command.flags[i]).
					WithCommand(command).
					WithMessage(
						"Duplicate flags detected.",
					).
					WithExpected(
						"Each flag must have a unique name.",
					).
					WithReceived(
						fmt.Sprintf("Duplicate flag: %s", command.flags[i].name),
					).
					WithSuggestion(
						"Rename or remove duplicate flags.",
					),
			)
		}

		if command.flags[i].required {
			required = append(required, command.flags[i])
		} else {
			optional = append(optional, command.flags[i])
		}

		names = append(names, command.flags[i].name)
	}

	command.flags = make([]*Flag, 0)
	command.flags = append(command.flags, required...)
	command.flags = append(command.flags, optional...)
	return command
}

func (command *Command) mustValidate() {
	if len(command.path) == 0 {
		panic(
			NewDefinitionError(CommandPathRequired, nil).
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

	command.mustNormalizeArguments().mustNormalizeFlags()
}

func mustResolveCommandPath() ([]string, string) {
	defer onError()

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
		panic(
			NewFrameworkError(InvalidCommandRoute, err).
				WithMessage(
					"Command route resolution failed.",
				),
		)
	}

	_, filePath, _, ok := runtime.Caller(3)

	if !ok {
		panic(
			NewFrameworkError(InvalidCommandRoute, nil).
				WithMessage(
					"Command route resolution failed.",
				),
		)
	}

	if !routesRegexp.MatchString(filePath) {
		panic(
			NewFrameworkError(InvalidCommandRoute, nil).
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
				),
		)
	}

	relFilePath, err := filepath.Rel(workPath, filepath.Clean(filePath))

	if err != nil {
		panic(
			NewFrameworkError(InvalidCommandRoute, err).
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
				),
		)
	}

	route := strings.Join(strings.Split(relFilePath, "\\commands\\")[1:], "")

	return route
}
