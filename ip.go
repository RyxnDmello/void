package void

import (
	"fmt"
	"net"
	"slices"
	"strings"
)

func NewIPFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.HandleConvert(func(value any) any {
		switch v := value.(type) {
		case nil, net.IP:
			return v

		case string:
			if ip := net.ParseIP(v); ip != nil {
				return ip
			}

			panic(fmt.Sprintf("'%s' Not A Valid IPv4 or IPv6 Address", v))

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be Either An IPv4 or IPv6 Address", flag.name))
		}
	})

	return flag
}

func NewIPListFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.HandleConvert(func(value any) any {
		switch v := value.(type) {
		case nil, []net.IP:
			return v

		case []string:
			ips := make([]net.IP, 0)

			for i := range v {
				ip := net.ParseIP(v[i])

				if ip == nil {
					panic(fmt.Sprintf("'%s' Not A Valid IPv4 or IPv6 Address", v[i]))
				}

				ips = append(ips, ip)
			}

			return ips

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A List Of IPv4 or IPv6 Addresses", flag.name))
		}
	})

	flag.variadic = true

	return flag
}

func NewIPEnumFlag(name string, enum []string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	if len(enum) == 0 {
		panic(fmt.Sprintf("Flag '%s' Must Have At Least 1 Option", flag.name))
	}

	for i := range enum {
		ip := net.ParseIP(enum[i])

		if ip == nil {
			panic(fmt.Sprintf("'%s' Not A Valid IPv4 or IPv6 Address", enum[i]))
		}

		flag.enum = append(flag.enum, ip)
	}

	flag.HandleConvert(func(value any) any {
		if value == nil {
			return nil
		}

		options := make([]string, 0)

		for i := range flag.enum {
			options = append(options, flag.enum[i].(net.IP).String())
		}

		switch v := value.(type) {
		case net.IP:
			if slices.Contains(options, v.String()) {
				return v
			}

			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))

		case string:
			ip := net.ParseIP(v)

			if ip == nil {
				panic(fmt.Sprintf("'%s' Not A Valid IPv4 or IPv6 Address", v))
			}

			if slices.Contains(options, ip.String()) {
				return ip
			}

			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))
		}
	})

	return flag
}
