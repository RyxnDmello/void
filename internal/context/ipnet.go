package context

import (
	"fmt"
	"net"
)

func (context *Context) MustGetIPNetFlag(name string) *net.IPNet {
	value, ok := context.MustGetFlag(name).(*net.IPNet)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A CIDR / Network Range", name))
	}

	return value
}

func (context *Context) MustGetIPNetListFlag(name string) []*net.IPNet {
	value, ok := context.MustGetFlag(name).([]*net.IPNet)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A List Of CIDRs / Network Ranges", name))
	}

	return value
}
