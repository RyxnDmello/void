package context

import (
	"fmt"
	"regexp"
)

func (context *Context) MustGetRegexpFlag(name string) *regexp.Regexp {
	value, ok := context.MustGetFlag(name).(*regexp.Regexp)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A Regular Expression", name))
	}

	return value
}

func (context *Context) MustGetRegexpListFlag(name string) []*regexp.Regexp {
	value, ok := context.MustGetFlag(name).([]*regexp.Regexp)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A List Of Regular Expressions", name))
	}

	return value
}
