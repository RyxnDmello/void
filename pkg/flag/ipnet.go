package flag

import (
	"fmt"
	"net"
	"slices"
	"strings"
)

func NewIPNetFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, *net.IPNet:
			return v

		case net.IPNet:
			return &v

		case string:
			if _, net, err := net.ParseCIDR(v); err == nil {
				return net
			}

			panic(fmt.Sprintf("'%s' Not A Valid CIDR / Network Range", v))

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A CIDR / Network Range", flag.name))
		}
	})

	return flag
}

func NewIPNetListFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, []*net.IPNet:
			return v

		case []net.IPNet:
			nets := make([]*net.IPNet, 0)

			for i := range v {
				nets = append(nets, &v[i])
			}

			return nets

		case []string:
			nets := make([]*net.IPNet, 0)

			for i := range v {
				_, net, err := net.ParseCIDR(v[i])

				if err != nil {
					panic(fmt.Sprintf("'%s' Not A Valid CIDR / Network Range", v[i]))
				}

				nets = append(nets, net)
			}

			return nets

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A List Of CIDRs / Network Ranges", flag.name))
		}
	})

	flag.variadic = true

	return flag
}

func NewIPNetEnumFlag(name string, enum []string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	if len(enum) == 0 {
		panic(fmt.Sprintf("Flag '%s' Must Have At Least 1 Option", flag.name))
	}

	for i := range enum {
		_, net, err := net.ParseCIDR(enum[i])

		if err != nil {
			panic(fmt.Sprintf("'%s' Not A Valid CIDR / Network Range", enum[i]))
		}

		flag.enum = append(flag.enum, net)
	}

	flag.SetConvert(func(value any) any {
		if value == nil {
			return nil
		}

		options := make([]string, 0)

		for i := range flag.enum {
			options = append(options, flag.enum[i].(*net.IPNet).String())
		}

		switch v := value.(type) {
		case *net.IPNet:
			if slices.Contains(options, v.String()) {
				return v
			}

			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))

		case net.IPNet:
			if slices.Contains(options, v.String()) {
				return &v
			}

			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))

		case string:
			_, net, err := net.ParseCIDR(v)

			if err != nil {
				panic(fmt.Sprintf("'%s' Not A Valid CIDR / Network Range", v))
			}

			if slices.Contains(options, net.String()) {
				return net
			}

			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))
		}
	})

	return flag
}
