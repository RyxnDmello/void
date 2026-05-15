package context

import (
	"fmt"
)

func (context *Context) MustGetIntArg(name string) int {
	value, ok := context.MustGetArg(name).(int)

	if !ok {
		panic(fmt.Sprintf("Argument '%s' Cannot Be Converted Into An Integer", name))
	}

	return value
}

func (context *Context) MustGetIntListArg(name string) []int {
	value, ok := context.MustGetArg(name).([]int)

	if !ok {
		panic(fmt.Sprintf("Argument '%s' Cannot Be Converted Into A List Of Integers", name))
	}

	return value
}

func (context *Context) MustGetIntFlag(name string) int {
	value, ok := context.MustGetFlag(name).(int)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into An Integer", name))
	}

	return value
}

func (context *Context) MustGetIntListFlag(name string) []int {
	value, ok := context.MustGetFlag(name).([]int)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A List Of Integers", name))
	}

	return value
}
