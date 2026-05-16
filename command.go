package void

import (
	"fmt"
	"path/filepath"
	"regexp"
	"runtime"
	"slices"
	"strings"
)

type Command struct {
	use         string
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
		use:         name,
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

func (command *Command) SetTagline(tagline string) *Command {
	command.tagline = strings.Trim(tagline, " ")
	return command
}

func (command *Command) SetDescription(description string) *Command {
	command.description = strings.Trim(description, " ")
	return command
}

func (command *Command) SetArguments(arguments ...*Argument) *Command {
	var variadic *Argument

	names := make([]string, 0)
	required := make([]*Argument, 0)
	optional := make([]*Argument, 0)

	for i := range arguments {
		arguments[i].mustValidate()

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

func (command *Command) SetFlags(flags ...*Flag) *Command {
	for i := range flags {
		flags[i].mustValidate()
	}

	command.flags = flags
	return command
}

func (command *Command) MarkDisabled() *Command {
	command.enabled = false
	return command
}

func (command *Command) HandleExecute(execute func(ctx *Context) error) *Command {
	if execute == nil {
		panic("Command Execution Handler Must Be Provided")
	}

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

func (command *Command) getArgs() ([]*Argument, []*Argument, *Argument) {
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

func (command *Command) getExecute(ctx *Context) error {
	if ctx == nil {
		panic("Command Execution Handler Must Have A Context")
	}

	if command.execute == nil {
		panic("Command Execution Handler Must Be Provided")
	}

	return command.execute(ctx)
}

func (command *Command) getPreExecute(ctx *Context) error {
	if ctx == nil {
		panic("Command Pre Execution Handler Must Have A Context")
	}

	if command.preExecute == nil {
		return nil
	}

	return command.preExecute(ctx)
}

func (command *Command) getPostExecute(ctx *Context) error {
	if ctx == nil {
		panic("Command Post Execution Handler Must Have A Context")
	}

	if command.postExecute == nil {
		return nil
	}

	return command.postExecute(ctx)
}

func (command *Command) mustValidate() {
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

func mustResolveCommandPath() ([]string, string) {
	filePath, err := mustResolveCommandFile()

	if err != nil {
		panic(err)
	}

	root := strings.Split(filepath.Dir(filePath), " ")
	base := strings.Split(filepath.Base(filePath), ".")[0]

	if len(root) == 0 || root[0] == "." {
		return []string{base}, base
	}

	if base == "root" {
		return root, root[len(root)-1]
	}

	root = append(root, base)

	return root, base
}

func mustResolveCommandFile() (string, error) {
	workPath, err := filepath.Abs(".")

	if err != nil {
		return "", CommandNotFound
	}

	_, filePath, _, ok := runtime.Caller(3)

	if !ok {
		return "", CommandNotFound
	}

	pattern := regexp.MustCompile(`^.*/cmd/[^/]+/commands(/.*)?\.go$`)

	valid := pattern.MatchString(filePath)

	if !valid {
		return "", CommandNotFound
	}

	filePath = filepath.Clean(filePath)

	path, err := filepath.Rel(workPath, filePath)

	if err != nil {
		return "", CommandNotFound
	}

	parts := strings.Split(path, "\\commands\\")[1:]

	path = strings.Join(parts, fmt.Sprintf("%v", filepath.Separator))

	return path, nil
}
