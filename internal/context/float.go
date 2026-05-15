package context

import (
	"fmt"
)

func (context *Context) MustGetFloatArg(name string) float64 {
	value, ok := context.MustGetArg(name).(float64)

	if !ok {
		panic(fmt.Sprintf("Argument '%s' Cannot Be Converted Into A Float", name))
	}

	return value
}

func (context *Context) MustGetFloatListArg(name string) []float64 {
	value, ok := context.MustGetArg(name).([]float64)

	if !ok {
		panic(fmt.Sprintf("Argument '%s' Cannot Be Converted Into A List Of Floats", name))
	}

	return value
}

func (context *Context) MustGetFloatFlag(name string) float64 {
	value, ok := context.MustGetFlag(name).(float64)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A Float", name))
	}

	return value
}

func (context *Context) MustGetFloatListFlag(name string) []float64 {
	value, ok := context.MustGetFlag(name).([]float64)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A List Of Floats", name))
	}

	return value
}
