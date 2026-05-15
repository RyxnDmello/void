package context

import (
	"fmt"
)

func (context *Context) MustGetStringMapFlag(name string) map[string]string {
	value, ok := context.MustGetFlag(name).(map[string]string)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A Map Of Strings", name))
	}

	return value
}

func (context *Context) MustGetIntMapFlag(name string) map[string]int {
	value, ok := context.MustGetFlag(name).(map[string]int)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A Map Of Integers", name))
	}

	return value
}

func (context *Context) MustGetFloatMapFlag(name string) map[string]float64 {
	value, ok := context.MustGetFlag(name).(map[string]float64)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A Map Of Floats", name))
	}

	return value
}

func (context *Context) MustGetBoolMapFlag(name string) map[string]bool {
	value, ok := context.MustGetFlag(name).(map[string]bool)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A Map Of Booleans", name))
	}

	return value
}

func (context *Context) MustGetComplexMapFlag(name string) map[string]any {
	value, ok := context.MustGetFlag(name).(map[string]any)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A Map Of Strings, Integers, Floats, And Booleans", name))
	}

	return value
}
