package context

import (
	"fmt"
	"time"
)

func (context *Context) MustGetDurationFlag(name string) time.Duration {
	value, ok := context.MustGetFlag(name).(time.Duration)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A Duration", name))
	}

	return value
}

func (context *Context) MustGetDurationListFlag(name string) []time.Duration {
	value, ok := context.MustGetFlag(name).([]time.Duration)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A List Of Durations", name))
	}

	return value
}
