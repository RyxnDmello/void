package argument

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func NewIntArgument(name string) *Argument {
	argument := NewCustomArgument(name, func(value any) any {
		return nil
	})

	argument.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, int:
			return v

		case string:
			if result, err := strconv.Atoi(v); err == nil {
				return result
			}

			panic(fmt.Sprintf("Argument '%s' Must Be An Integer", argument.name))

		default:
			panic(fmt.Sprintf("Argument '%s' Must Be An Integer", argument.name))
		}
	})

	return argument
}

func NewIntListArgument(name string) *Argument {
	argument := NewCustomArgument(name, func(value any) any {
		return nil
	})

	argument.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, []int:
			return v

		case []string:
			result := make([]int, 0)

			for i := range v {
				value, err := strconv.Atoi(v[i])

				if err != nil {
					panic(fmt.Sprintf("Argument '%s' Must Be A List Of Integers", argument.name))
				}

				result = append(result, value)
			}

			return result

		default:
			panic(fmt.Sprintf("Argument '%s' Must Be A List Of Integers", argument.name))
		}
	})

	argument.variadic = true

	return argument
}

func NewIntEnumArgument(name string, enum []int) *Argument {
	argument := NewCustomArgument(name, func(value any) any {
		return nil
	})

	if len(enum) == 0 {
		panic(fmt.Sprintf("Argument '%s' Must Have At Least 1 Option", argument.name))
	}

	for i := range enum {
		argument.enum = append(argument.enum, enum[i])
	}

	argument.SetConvert(func(value any) any {
		if value == nil {
			return nil
		}

		enum := make([]string, 0)

		for i := range argument.enum {
			enum = append(enum, fmt.Sprintf("%v", argument.enum[i]))
		}

		switch v := value.(type) {
		case int, string:
			raw := fmt.Sprintf("%v", v)

			result, err := strconv.Atoi(raw)

			if err != nil {
				panic(fmt.Sprintf("Argument '%s' Must Be An Integer", argument.name))
			}

			if !slices.Contains(enum, raw) {
				panic(fmt.Sprintf("Argument '%s' Must Be One Of %s", argument.name, strings.Join(enum, ", ")))
			}

			return result

		default:
			panic(fmt.Sprintf("Argument '%s' Must Be One Of %s", argument.name, strings.Join(enum, ", ")))
		}
	})

	return argument
}
