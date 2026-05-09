package flag

import (
	"fmt"
	"slices"
	"strings"
)

func NewStringFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, string:
			return v

		default:
			return fmt.Sprintf("%v", v)
		}
	})

	return flag
}

func NewStringListFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, []string:
			return v

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A List Of Strings", flag.name))
		}
	})

	flag.variadic = true

	return flag
}

func NewStringEnumFlag(name string, enum []string) *Flag {
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
		switch v := value.(type) {
		case nil:
			return v

		case string:
			if slices.Contains(enum, v) {
				return v
			}

			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(enum, ", ")))

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(enum, ", ")))
		}
	})

	return flag
}
