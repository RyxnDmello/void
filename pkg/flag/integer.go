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
			if result, err := strconv.Atoi(v); err == nil {
				return result
			}

			panic(fmt.Sprintf("Flag '%s' Must Be An Integer", flag.name))

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be An Integer", flag.name))
		}
	})

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

		enum := make([]string, 0)

		for i := range flag.enum {
			enum = append(enum, fmt.Sprintf("%v", flag.enum[i]))
		}

		switch v := value.(type) {
		case int, string:
			raw := fmt.Sprintf("%v", v)

			result, err := strconv.Atoi(raw)

			if err != nil {
				panic(fmt.Sprintf("Flag '%s' Must Be An Integer", flag.name))
			}

			if !slices.Contains(enum, raw) {
				panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(enum, ", ")))
			}

			return result

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(enum, ", ")))
		}
	})

	return flag
}

func NewIntListFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.variadic = true

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, []int:
			return v

		case []string:
			result := make([]int, 0)

			for i := range v {
				value, err := strconv.Atoi(v[i])

				if err != nil {
					panic(fmt.Sprintf("Flag '%s' Must Be A List Of Integers", flag.name))
				}

				result = append(result, value)
			}

			return result

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A List Of Integers", flag.name))
		}
	})

	return flag
}
