package void

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

	argument.HandleConvert(func(value any) any {
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

func NewFloatFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.HandleConvert(func(value any) any {
		switch v := value.(type) {
		case nil, float64:
			return v

		case int, float32, string:
			if float, err := strconv.ParseFloat(fmt.Sprintf("%v", v), 64); err == nil {
				return float
			}

			panic(fmt.Sprintf("Flag '%s' Must Be A Float", flag.name))

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A Float", flag.name))
		}
	})

	return flag
}

func NewFloatListArgument(name string) *Argument {
	argument := NewCustomArgument(name, func(value any) any {
		return nil
	})

	argument.HandleConvert(func(value any) any {
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

	argument.variadic = true

	return argument
}

func NewFloatListFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.HandleConvert(func(value any) any {
		switch v := value.(type) {
		case nil, []float64:
			return v

		case []int:
			floats := make([]float64, 0)

			for i := range v {
				float, err := strconv.ParseFloat(fmt.Sprintf("%v", v[i]), 64)

				if err != nil {
					panic(fmt.Sprintf("Flag '%s' Must Be A List Of Floats", flag.name))
				}

				floats = append(floats, float)
			}

			return floats

		case []float32:
			floats := make([]float64, 0)

			for i := range v {
				float, err := strconv.ParseFloat(fmt.Sprintf("%v", v[i]), 64)

				if err != nil {
					panic(fmt.Sprintf("Flag '%s' Must Be A List Of Floats", flag.name))
				}

				floats = append(floats, float)
			}

			return floats

		case []string:
			floats := make([]float64, 0)

			for i := range v {
				float, err := strconv.ParseFloat(v[i], 64)

				if err != nil {
					panic(fmt.Sprintf("Flag '%s' Must Be A List Of Floats", flag.name))
				}

				floats = append(floats, float)
			}

			return floats

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A List Of Floats", flag.name))
		}
	})

	flag.variadic = true

	return flag
}

func NewFloatEnumArgument(name string, enum []float64) *Argument {
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

func NewFloatEnumFlag(name string, enum []float64) *Flag {
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
		case int, float32, float64, string:
			raw := fmt.Sprintf("%v", v)

			float, err := strconv.ParseFloat(raw, 64)

			if err != nil {
				panic(fmt.Sprintf("Flag '%s' Must Be A Float", flag.name))
			}

			if slices.Contains(options, raw) {
				return float
			}

			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))
		}
	})

	return flag
}
