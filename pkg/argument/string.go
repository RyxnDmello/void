package argument

import (
	"fmt"
	"slices"
	"strings"
)

func NewStringArgument(name string) *Argument {
	argument := NewCustomArgument(name, func(value any) any {
		return nil
	})

	argument.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, string:
			return v

		default:
			return fmt.Sprintf("%v", v)
		}
	})

	return argument
}

func NewStringListArgument(name string) *Argument {
	argument := NewCustomArgument(name, func(value any) any {
		return nil
	})

	argument.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, []string:
			return v

		default:
			panic(fmt.Sprintf("Argument '%s' Must Be A List Of Strings", argument.name))
		}
	})

	argument.variadic = true

	return argument
}

func NewStringEnumArgument(name string, enum []string) *Argument {
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
			enum = append(enum, argument.enum[i].(string))
		}

		switch v := value.(type) {
		case string:
			if !slices.Contains(enum, v) {
				panic(fmt.Sprintf("Argument '%s' Must Be One Of %s", argument.name, strings.Join(enum, ", ")))
			}

			return v

		default:
			panic(fmt.Sprintf("Argument '%s' Must Be One Of %s", argument.name, strings.Join(enum, ", ")))
		}
	})

	return argument
}
