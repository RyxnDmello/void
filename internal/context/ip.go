package context

import (
	"fmt"
	"net"
)

func (context *Context) MustGetIPFlag(name string) net.IP {
	value, ok := context.MustGetFlag(name).(net.IP)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into An IPv4 Or IPv6 Address", name))
	}

	return value
}

func (context *Context) MustGetIPListFlag(name string) []net.IP {
	value, ok := context.MustGetFlag(name).([]net.IP)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A List Of IPv4 Or IPv6 Addresses", name))
	}

	return value
}
