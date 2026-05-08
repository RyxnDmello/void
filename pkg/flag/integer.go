package flag

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func NewIntFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, int:
			return v

		case string:
			if integer, err := strconv.Atoi(v); err == nil {
				return integer
			}

			panic(fmt.Sprintf("Flag '%s' Must Be An Integer", flag.name))

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be An Integer", flag.name))
		}
	})

	return flag
}

func NewIntListFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, []int:
			return v

		case []string:
			integers := make([]int, 0)

			for i := range v {
				integer, err := strconv.Atoi(v[i])

				if err != nil {
					panic(fmt.Sprintf("Flag '%s' Must Be A List Of Integers", flag.name))
				}

				integers = append(integers, integer)
			}

			return integers

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A List Of Integers", flag.name))
		}
	})

	flag.variadic = true

	return flag
}

func NewIntEnumFlag(name string, enum []int) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	if len(enum) == 0 {
		panic(fmt.Sprintf("Flag '%s' Must Have At Least 1 Option", flag.name))
	}

	for i := range enum {
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

			integer, err := strconv.Atoi(raw)

			if err != nil {
				panic(fmt.Sprintf("Flag '%s' Must Be An Integer", flag.name))
			}

			if !slices.Contains(options, raw) {
				panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))
			}

			return integer

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))
		}
	})

	return flag
}
