package context

import (
	"fmt"
	"net/mail"
)

func (context *Context) MustGetEmailFlag(name string) *mail.Address {
	value, ok := context.MustGetFlag(name).(*mail.Address)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into An Email Address", name))
	}

	return value
}

func (context *Context) MustGetEmailListFlag(name string) []*mail.Address {
	value, ok := context.MustGetFlag(name).([]*mail.Address)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A List Of Email Addresses", name))
	}

	return value
}
