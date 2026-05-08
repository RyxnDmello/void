package context

import (
	"fmt"

	"github.com/ryxndmello/void/pkg/argument"
)

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

func (context *Context) MustGetStringArg(name string) string {
	value, ok := context.MustGetArg(name).(string)

	if !ok {
		panic(fmt.Sprintf("Argument '%s' Cannot Be Converted Into A String", name))
	}

	return value
}

func (context *Context) MustGetIntArg(name string) int {
	value, ok := context.MustGetArg(name).(int)

	if !ok {
		panic(fmt.Sprintf("Argument '%s' Cannot Be Converted Into An Integer", name))
	}

	return value
}

func (context *Context) MustGetFloatArg(name string) float64 {
	value, ok := context.MustGetArg(name).(float64)

	if !ok {
		panic(fmt.Sprintf("Argument '%s' Cannot Be Converted Into A Float", name))
	}

	return value
}

func (context *Context) MustGetBoolArg(name string) bool {
	value, ok := context.MustGetArg(name).(bool)

	if !ok {
		panic(fmt.Sprintf("Argument '%s' Cannot Be Converted Into A Boolean", name))
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

func (context *Context) MustGetIntListArg(name string) []int {
	value, ok := context.MustGetArg(name).([]int)

	if !ok {
		panic(fmt.Sprintf("Argument '%s' Cannot Be Converted Into A List Of Integers", name))
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

func (context *Context) MustGetBoolListArg(name string) []bool {
	value, ok := context.MustGetArg(name).([]bool)

	if !ok {
		panic(fmt.Sprintf("Argument '%s' Cannot Be Converted Into A List Of Booleans", name))
	}

	return value
}
