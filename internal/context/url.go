package context

import (
	"fmt"
	"net/url"
)

func (context *Context) MustGetURLFlag(name string) *url.URL {
	value, ok := context.MustGetFlag(name).(*url.URL)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A URL", name))
	}

	return value
}

func (context *Context) MustGetURLListFlag(name string) []*url.URL {
	value, ok := context.MustGetFlag(name).([]*url.URL)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A List Of URLs", name))
	}

	return value
}
