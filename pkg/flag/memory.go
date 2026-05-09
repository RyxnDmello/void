package flag

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var memoryRegexp = regexp.MustCompile(`^([0-9]+(?:\.[0-9]+)?)([A-Za-z]+)$`)

var unitToMemory = map[string]int64{
	"B": 1,

	"KB": 1000,
	"MB": 1000 * 1000,
	"GB": 1000 * 1000 * 1000,
	"TB": 1000 * 1000 * 1000 * 1000,
	"PB": 1000 * 1000 * 1000 * 1000 * 1000,

	"KiB": 1024,
	"MiB": 1024 * 1024,
	"GiB": 1024 * 1024 * 1024,
	"TiB": 1024 * 1024 * 1024 * 1024,
	"PiB": 1024 * 1024 * 1024 * 1024 * 1024,
}

func NewMemoryFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil:
			return v

		case string:
			return mustConvMemory(v)

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A Unit Of Memory", flag.name))
		}
	})

	return flag
}

func NewMemoryListFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil:
			return v

		case []string:
			memories := make([]float64, 0)

			for i := range v {
				memories = append(memories, mustConvMemory(v[i]))
			}

			return memories

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A List Of Memory Units", flag.name))
		}
	})

	flag.variadic = true

	return flag
}

func NewMemoryEnumFlag(name string, enum []string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	if len(enum) == 0 {
		panic(fmt.Sprintf("Flag '%s' Must Have At Least 1 Option", flag.name))
	}

	for i := range enum {
		flag.enum = append(flag.enum, mustConvMemory(enum[i]))
	}

	flag.SetConvert(func(value any) any {
		if value == nil {
			return nil
		}

		options := make([]string, 0)

		for i := range flag.enum {
			options = append(options, fmt.Sprintf("%v", flag.enum[i].(float64)))
		}

		switch v := value.(type) {
		case string:
			if memory := mustConvMemory(v); slices.Contains(options, fmt.Sprintf("%v", memory)) {
				return memory
			}

			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))
		}
	})

	return flag
}

func mustConvMemory(value string) float64 {
	matches := memoryRegexp.FindStringSubmatch(value)

	if len(matches) == 0 {
		panic(fmt.Sprintf("'%s' Not A Valid Size", value))
	}

	quant := matches[1]
	units := matches[2]

	multiplier, ok := unitToMemory[units]

	if !ok {
		panic(fmt.Sprintf("'%s' Not A Valid Unit Of Memory", units))
	}

	if result, err := strconv.ParseFloat(quant, 64); err == nil {
		return result * float64(multiplier)
	}

	panic(fmt.Sprintf("'%s' Not A Valid Number", quant))
}
