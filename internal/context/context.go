package context

import (
	"fmt"

	"github.com/ryxndmello/void/pkg/argument"
	"github.com/ryxndmello/void/pkg/flag"
)

type Context struct {
	path      string
	flags     map[string]any
	arguments map[string]any
}

func NewContext(path string, arguments map[string]any, flags map[string]any) *Context {
	return &Context{
		path:      path,
		flags:     flags,
		arguments: arguments,
	}
}

func (context *Context) Path() string {
	return context.path
}

func (context *Context) Args() map[string]any {
	return context.arguments
}

func (context *Context) GetArg(name string) any {
	name = argument.MustParseArgumentName(name)

	if len(context.arguments) == 0 {
		return nil
	}

	return context.arguments[name]
}

func (context *Context) MustGetArg(name string) any {
	value := context.GetArg(name)

	if value == nil {
		panic(fmt.Sprintf("Argument '%s' Does Not Exist", name))
	}

	return value
}

func (context *Context) Flags() map[string]any {
	return context.flags
}

func (context *Context) GetFlag(name string) any {
	name = flag.MustParseFlagName(name)

	if len(context.flags) == 0 {
		return nil
	}

	return context.flags[name]
}

func (context *Context) MustGetFlag(name string) any {
	value := context.GetFlag(name)

	if value == nil {
		panic(fmt.Sprintf("Flag '%s' Does Not Exist", name))
	}

	return value
}
