package context

import (
	"fmt"
	"time"
)

func (context *Context) MustGetTimeFlag(name string) time.Time {
	value, ok := context.MustGetFlag(name).(time.Time)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A Date Time", name))
	}

	return value
}

func (context *Context) MustGetTimeListFlag(name string) []time.Time {
	value, ok := context.MustGetFlag(name).([]time.Time)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A List Of Date Times", name))
	}

	return value
}
