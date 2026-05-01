package flag

import (
	"fmt"
	"strconv"
)

func NewBoolFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.value = false

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, bool:
			return v

		case string:
			if result, err := strconv.ParseBool(v); err == nil {
				return result
			}

			panic(fmt.Sprintf("Flag '%s' Must Be A Boolean", flag.name))

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A Boolean", flag.name))
		}
	})

	return flag
}

func NewBoolListFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.variadic = true

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, []bool:
			return v

		case []string:
			result := make([]bool, 0)

			for i := range v {
				value, err := strconv.ParseBool(v[i])

				if err != nil {
					panic(fmt.Sprintf("Flag '%s' Must Be A List Of Booleans", flag.name))
				}

				result = append(result, value)
			}

			return result

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A List Of Booleans", flag.name))
		}
	})

	return flag
}
