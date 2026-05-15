package flag

import (
	"fmt"
	"strconv"
)

func NewBoolFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, bool:
			return v

		case string:
			if boolean, err := strconv.ParseBool(v); err == nil {
				return boolean
			}

			panic(fmt.Sprintf("Flag '%s' Must Be A Boolean", flag.name))

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A Boolean", flag.name))
		}
	})

	flag.value = false

	return flag
}

func NewBoolListFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, []bool:
			return v

		case []string:
			booleans := make([]bool, 0)

			for i := range v {
				boolean, err := strconv.ParseBool(v[i])

				if err != nil {
					panic(fmt.Sprintf("Flag '%s' Must Be A List Of Booleans", flag.name))
				}

				booleans = append(booleans, boolean)
			}

			return booleans

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A List Of Booleans", flag.name))
		}
	})

	flag.variadic = true

	return flag
}
