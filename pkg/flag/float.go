package flag

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func NewFloatFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, float64:
			return v

		case int, string:
			raw := fmt.Sprintf("%v", v)

			if result, err := strconv.ParseFloat(raw, 64); err == nil {
				return result
			}

			panic(fmt.Sprintf("Flag '%s' Must Be A Float", flag.name))

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A Float", flag.name))
		}
	})

	return flag
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

	flag.SetConvert(func(value any) any {
		if value == nil {
			return nil
		}

		enum := make([]string, 0)

		for i := range flag.enum {
			enum = append(enum, fmt.Sprintf("%v", flag.enum[i]))
		}

		switch v := value.(type) {
		case int, float64, string:
			raw := fmt.Sprintf("%v", v)

			result, err := strconv.ParseFloat(raw, 64)

			if err != nil {
				panic(fmt.Sprintf("Flag '%s' Must Be A Float", flag.name))
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

func NewFloatListFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.variadic = true

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, []float64:
			return v

		case []int:
			result := make([]float64, 0)

			for i := range v {
				raw := fmt.Sprintf("%v", v[i])

				value, err := strconv.ParseFloat(raw, 64)

				if err != nil {
					panic(fmt.Sprintf("Flag '%s' Must Be A List Of Floats", flag.name))
				}

				result = append(result, value)
			}

			return result

		case []string:
			result := make([]float64, 0)

			for i := range v {
				value, err := strconv.ParseFloat(v[i], 64)

				if err != nil {
					panic(fmt.Sprintf("Flag '%s' Must Be A List Of Floats", flag.name))
				}

				result = append(result, value)
			}

			return result

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A List Of Floats", flag.name))
		}
	})

	return flag
}
