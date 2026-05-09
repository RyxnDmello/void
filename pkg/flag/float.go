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

func NewFloatListFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
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

		options := make([]string, 0)

		for i := range enum {
			options = append(options, fmt.Sprintf("%v", enum[i]))
		}

		switch v := value.(type) {
		case int, float64, string:
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
