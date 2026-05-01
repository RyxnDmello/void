package argument

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func NewFloatArgument(name string) *Argument {
	argument := NewCustomArgument(name, func(value any) any {
		return nil
	})

	argument.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, float64:
			return v

		case int, string:
			raw := fmt.Sprintf("%v", v)

			if result, err := strconv.ParseFloat(raw, 64); err == nil {
				return result
			}

			panic(fmt.Sprintf("Argument '%s' Must Be A Float", argument.name))

		default:
			panic(fmt.Sprintf("Argument '%s' Must Be A Float", argument.name))
		}
	})

	return argument
}

func NewFloatEnumArgument(name string, enum []float64) *Argument {
	argument := NewCustomArgument(name, func(value any) any {
		return nil
	})

	if len(enum) == 0 {
		panic(fmt.Sprintf("Argument '%s' Of Type Enum Must Have At Least 1 Option", argument.name))
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
		case int, float64, string:
			raw := fmt.Sprintf("%v", v)

			result, err := strconv.ParseFloat(raw, 64)

			if err != nil {
				panic(fmt.Sprintf("Argument '%s' Must Be A Float", argument.name))
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

func NewFloatListArgument(name string) *Argument {
	argument := NewCustomArgument(name, func(value any) any {
		return nil
	})

	argument.variadic = true

	argument.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, []float64:
			return v

		case []int:
			result := make([]float64, 0)

			for i := range v {
				raw := fmt.Sprintf("%v", v[i])

				value, err := strconv.ParseFloat(raw, 64)

				if err != nil {
					panic(fmt.Sprintf("Argument '%s' Must Be A List Of Floats", argument.name))
				}

				result = append(result, value)
			}

			return result

		case []string:
			result := make([]float64, 0)

			for i := range v {
				value, err := strconv.ParseFloat(v[i], 64)

				if err != nil {
					panic(fmt.Sprintf("Argument '%s' Must Be A List Of Floats", argument.name))
				}

				result = append(result, value)
			}

			return result

		default:
			panic(fmt.Sprintf("Argument '%s' Must Be A List Of Floats", argument.name))
		}
	})

	return argument
}
