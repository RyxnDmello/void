package void

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

	argument.HandleConvert(func(value any) any {
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

func NewIntFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.HandleConvert(func(value any) any {
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

func NewIntListArgument(name string) *Argument {
	argument := NewCustomArgument(name, func(value any) any {
		return nil
	})

	argument.HandleConvert(func(value any) any {
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

func NewIntListFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.HandleConvert(func(value any) any {
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

	argument.HandleConvert(func(value any) any {
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

	flag.HandleConvert(func(value any) any {
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

			if slices.Contains(options, raw) {
				return integer
			}

			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))
		}
	})

	return flag
}
