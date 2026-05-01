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

func (context *Context) Args() map[string]any {
	return context.arguments
}

func (context *Context) GetArg(name string) any {
	name = argument.MustParseArgumentName(name)
	return getValue(context.arguments, name)
}

func (context *Context) MustGetArg(name string) any {
	value := context.GetArg(name)

	if value == nil {
		panic(fmt.Sprintf("Argument '%s' Does Not Exist", name))
	}

	return value
}

func (context *Context) MustGetStringArg(name string) string {
	value := getStringValue(context.MustGetArg(name))

	if value == nil {
		panic(fmt.Sprintf("Argument '%s' Cannot Be Converted Into A String", name))
	}

	return *value
}

func (context *Context) MustGetIntArg(name string) int {
	value := getIntValue(context.MustGetArg(name))

	if value == nil {
		panic(fmt.Sprintf("Argument '%s' Cannot Be Converted Into An Integer", name))
	}

	return *value
}

func (context *Context) MustGetFloatArg(name string) float64 {
	value := getFloatValue(context.MustGetArg(name))

	if value == nil {
		panic(fmt.Sprintf("Argument '%s' Cannot Be Converted Into A Float", name))
	}

	return *value
}

func (context *Context) MustGetBoolArg(name string) bool {
	value := getBoolValue(context.MustGetArg(name))

	if value == nil {
		panic(fmt.Sprintf("Argument '%s' Cannot Be Converted Into A Boolean", name))
	}

	return *value
}

func (context *Context) MustGetStringListArg(name string) []string {
	value := getStringListValue(context.MustGetArg(name))

	if value == nil {
		panic(fmt.Sprintf("Argument '%s' Cannot Be Converted Into A List Of Strings", name))
	}

	return *value
}

func (context *Context) MustGetIntListArg(name string) []int {
	value := getIntListValue(context.MustGetArg(name))

	if value == nil {
		panic(fmt.Sprintf("Argument '%s' Cannot Be Converted Into A List Of Integers", name))
	}

	return *value
}

func (context *Context) MustGetFloatListArg(name string) []float64 {
	value := getFloatListValue(context.MustGetArg(name))

	if value == nil {
		panic(fmt.Sprintf("Argument '%s' Cannot Be Converted Into A List Of Floats", name))
	}

	return *value
}

func (context *Context) MustGetBoolListArg(name string) []bool {
	value := getBoolListValue(context.MustGetArg(name))

	if value == nil {
		panic(fmt.Sprintf("Argument '%s' Cannot Be Converted Into A List Of Booleans", name))
	}

	return *value
}

func (context *Context) Flags() map[string]any {
	return context.flags
}

func (context *Context) GetFlag(name string) any {
	name = flag.MustParseFlagName(name)
	return getValue(context.flags, name)
}

func (context *Context) MustGetFlag(name string) any {
	value := context.GetFlag(name)

	if value == nil {
		panic(fmt.Sprintf("Flag '%s' Does Not Exist", name))
	}

	return value
}

func (context *Context) MustGetStringFlag(name string) string {
	value := getStringValue(context.MustGetFlag(name))

	if value == nil {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A String", name))
	}

	return *value
}

func (context *Context) MustGetIntFlag(name string) int {
	value := getIntValue(context.MustGetFlag(name))

	if value == nil {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into An Integer", name))
	}

	return *value
}

func (context *Context) MustGetFloatFlag(name string) float64 {
	value := getFloatValue(context.MustGetFlag(name))

	if value == nil {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A Float", name))
	}

	return *value
}

func (context *Context) MustGetBoolFlag(name string) bool {
	value := getBoolValue(context.MustGetFlag(name))

	if value == nil {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A Boolean", name))
	}

	return *value
}

func (context *Context) MustGetStringListFlag(name string) []string {
	value := getStringListValue(context.MustGetFlag(name))

	if value == nil {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A List Of Strings", name))
	}

	return *value
}

func (context *Context) MustGetIntListFlag(name string) []int {
	value := getIntListValue(context.MustGetFlag(name))

	if value == nil {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A List Of Integers", name))
	}

	return *value
}

func (context *Context) MustGetFloatListFlag(name string) []float64 {
	value := getFloatListValue(context.MustGetFlag(name))

	if value == nil {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A List Of Floats", name))
	}

	return *value
}

func (context *Context) MustGetBoolListFlag(name string) []bool {
	value := getBoolListValue(context.MustGetFlag(name))

	if value == nil {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A List Of Booleans", name))
	}

	return *value
}
