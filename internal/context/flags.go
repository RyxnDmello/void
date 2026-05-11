package context

import (
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"time"

	"github.com/ryxndmello/void/pkg/flag"
)

func (context *Context) GetFlag(name string) any {
	name = flag.MustParseFlagName(name)

	if len(context.flags) == 0 {
		return nil
	}

	return context.flags[name]
}

func (context *Context) MustGetFlag(name string) any {
	value := context.GetFlag(name)

	if value == nil {
		panic(fmt.Sprintf("Flag '%s' Does Not Exist", name))
	}

	return value
}

func (context *Context) MustGetStringFlag(name string) string {
	value, ok := context.MustGetFlag(name).(string)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A String", name))
	}

	return value
}

func (context *Context) MustGetIntFlag(name string) int {
	value, ok := context.MustGetFlag(name).(int)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into An Integer", name))
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

func (context *Context) MustGetBoolFlag(name string) bool {
	value, ok := context.MustGetFlag(name).(bool)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A Boolean", name))
	}

	return value
}

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

func (context *Context) MustGetRegexpFlag(name string) *regexp.Regexp {
	value, ok := context.MustGetFlag(name).(*regexp.Regexp)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A Regular Expression", name))
	}

	return value
}

func (context *Context) MustGetTimeFlag(name string) time.Time {
	value, ok := context.MustGetFlag(name).(time.Time)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A Date Time", name))
	}

	return value
}

func (context *Context) MustGetDurationFlag(name string) time.Duration {
	value, ok := context.MustGetFlag(name).(time.Duration)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A Duration", name))
	}

	return value
}

func (context *Context) MustGetURLFlag(name string) *url.URL {
	value, ok := context.MustGetFlag(name).(*url.URL)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A URL", name))
	}

	return value
}

func (context *Context) MustGetIPFlag(name string) net.IP {
	value, ok := context.MustGetFlag(name).(net.IP)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into An IPv4 Or IPv6 Address", name))
	}

	return value
}

func (context *Context) MustGetIPNetFlag(name string) *net.IPNet {
	value, ok := context.MustGetFlag(name).(*net.IPNet)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A CIDR / Network Range", name))
	}

	return value
}

func (context *Context) MustGetEmailFlag(name string) *mail.Address {
	value, ok := context.MustGetFlag(name).(*mail.Address)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into An Email Address", name))
	}

	return value
}

func (context *Context) MustGetStringListFlag(name string) []string {
	value, ok := context.MustGetFlag(name).([]string)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A List Of Strings", name))
	}

	return value
}

func (context *Context) MustGetIntListFlag(name string) []int {
	value, ok := context.MustGetFlag(name).([]int)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A List Of Integers", name))
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

func (context *Context) MustGetBoolListFlag(name string) []bool {
	value, ok := context.MustGetFlag(name).([]bool)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A List Of Booleans", name))
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

func (context *Context) MustGetTimeListFlag(name string) []time.Time {
	value, ok := context.MustGetFlag(name).([]time.Time)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A List Of Date Times", name))
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

func (context *Context) MustGetURLListFlag(name string) []*url.URL {
	value, ok := context.MustGetFlag(name).([]*url.URL)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A List Of URLs", name))
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

func (context *Context) MustGetIPNetListFlag(name string) []*net.IPNet {
	value, ok := context.MustGetFlag(name).([]*net.IPNet)

	if !ok {
		panic(fmt.Sprintf("Flag '%s' Cannot Be Converted Into A List Of CIDRs / Network Ranges", name))
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
