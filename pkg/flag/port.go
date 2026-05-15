package flag

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

const minPort = 1
const maxPort = 65535

func NewPortFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil:
			return v

		case int:
			if minPort <= v && v <= maxPort {
				return v
			}

			panic(fmt.Sprintf("A Network Port Must Be Between %v And %v", minPort, maxPort))

		case string:
			port, err := strconv.Atoi(v)

			if err != nil {
				panic(fmt.Sprintf("A Network Port Must Be An Integer Between %v And %v", minPort, maxPort))
			}

			if minPort <= port && port <= maxPort {
				return port
			}

			panic(fmt.Sprintf("A Network Port Must Be An Integer Between %v And %v", minPort, maxPort))

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A Network Port", flag.name))
		}
	})

	return flag
}

func NewPortListFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil:
			return v

		case []int:
			ports := make([]int, 0)

			for i := range v {
				if v[i] < minPort || v[i] > maxPort {
					panic(fmt.Sprintf("'%v' Not A Valid Port Between %v And %v", v[i], minPort, maxPort))
				}

				ports = append(ports, v[i])
			}

			return ports

		case []string:
			ports := make([]int, 0)

			for i := range v {
				port, err := strconv.Atoi(v[i])

				if err != nil {
					panic(fmt.Sprintf("'%s' Not A Valid Port. Must Be An Integer Between %v And %v", v[i], minPort, maxPort))
				}

				if port < minPort || port > maxPort {
					panic(fmt.Sprintf("'%v' Not A Valid Port Between %v And %v", v[i], minPort, maxPort))
				}

				ports = append(ports, port)
			}

			return ports

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A List Of Network Ports", flag.name))
		}
	})

	flag.variadic = true

	return flag
}

func NewPortEnumFlag(name string, enum []int) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	if len(enum) == 0 {
		panic(fmt.Sprintf("Flag '%s' Must Have At Least 1 Option", flag.name))
	}

	for i := range enum {
		if enum[i] < minPort || enum[i] > maxPort {
			panic(fmt.Sprintf("Flag '%s' Has An Invalid Option: '%v'", flag.name, enum[i]))
		}

		flag.enum = append(flag.enum, enum[i])
	}

	flag.SetConvert(func(value any) any {
		if value == nil {
			return nil
		}

		options := make([]string, 0)

		for i := range enum {
			options = append(options, fmt.Sprintf("%v", enum[i]))
		}

		switch v := value.(type) {
		case int, string:
			raw := fmt.Sprintf("%v", v)

			port, err := strconv.Atoi(raw)

			if err != nil {
				panic(fmt.Sprintf("Flag '%s' Must Be An Integer", flag.name))
			}

			if slices.Contains(options, raw) {
				return port
			}

			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))
		}
	})

	return flag
}
