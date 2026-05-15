package context

import (
	"fmt"
)

func (context *Context) MustGetStringArg(name string) string {
	value, ok := context.MustGetArg(name).(string)

	if !ok {
		panic(fmt.Sprintf("Argument '%s' Cannot Be Converted Into A String", name))
	}

	return value
}

func (context *Context) MustGetStringListArg(name string) []string {
	value, ok := context.MustGetArg(name).([]string)

	if !ok {
		panic(fmt.Sprintf("Argument '%s' Cannot Be Converted Into A List Of Strings", name))
	}

	return value
}

func (context *Context) MustGetStringFlag(name string) string {
	value, ok := context.MustGetFlag(name).(string)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A String", name))
	}

	return value
}

func (context *Context) MustGetStringListFlag(name string) []string {
	value, ok := context.MustGetFlag(name).([]string)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A List Of Strings", name))
	}

	return value
}
