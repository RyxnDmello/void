package void

import (
	"fmt"
	"strconv"
)

func NewBoolArgument(name string) *Argument {
	argument := NewCustomArgument(name, func(value any) any {
		return nil
	})

	argument.HandleConvert(func(value any) any {
		switch v := value.(type) {
		case nil, bool:
			return v

		case string:
			if v == argument.name {
				return true
			}

			if result, err := strconv.ParseBool(v); err == nil {
				return result
			}

			panic(fmt.Sprintf("Argument '%s' Must Be A Boolean", argument.name))

		default:
			panic(fmt.Sprintf("Argument '%s' Must Be A Boolean", argument.name))
		}
	})

	return argument
}

func NewBoolFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.HandleConvert(func(value any) any {
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

func NewBoolListArgument(name string) *Argument {
	argument := NewCustomArgument(name, func(value any) any {
		return nil
	})

	argument.HandleConvert(func(value any) any {
		switch v := value.(type) {
		case nil, []bool:
			return v

		case []string:
			result := make([]bool, 0)

			for i := range v {
				value, err := strconv.ParseBool(v[i])

				if err != nil {
					panic(fmt.Sprintf("Argument '%s' Must Be A List Of Booleans", argument.name))
				}

				result = append(result, value)
			}

			return result

		default:
			panic(fmt.Sprintf("Argument '%s' Must Be A List Of Booleans", argument.name))
		}
	})

	argument.variadic = true

	return argument
}

func NewBoolListFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.HandleConvert(func(value any) any {
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
