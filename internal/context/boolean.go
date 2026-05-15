package context

import (
	"fmt"
)

func (context *Context) MustGetBoolArg(name string) bool {
	value, ok := context.MustGetArg(name).(bool)

	if !ok {
		panic(fmt.Sprintf("Argument '%s' Cannot Be Converted Into A Boolean", name))
	}

	return value
}

func (context *Context) MustGetBoolListArg(name string) []bool {
	value, ok := context.MustGetArg(name).([]bool)

	if !ok {
		panic(fmt.Sprintf("Argument '%s' Cannot Be Converted Into A List Of Booleans", name))
	}

	return value
}

func (context *Context) MustGetBoolFlag(name string) bool {
	value, ok := context.MustGetFlag(name).(bool)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A Boolean", name))
	}

	return value
}

func (context *Context) MustGetBoolListFlag(name string) []bool {
	value, ok := context.MustGetFlag(name).([]bool)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A List Of Booleans", name))
	}

	return value
}
