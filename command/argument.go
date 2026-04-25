package command

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Argument struct {
	name        string
	tagline     string
	description string
	required    bool
	variadic    bool
	value       any
	enum        []string
	handleValue func(value any) any
}

func NewArgument(name string) *Argument {
	argument := NewCustomArgument(name, func(value any) any {
		return nil
	})

	argument.handleValue = func(value any) any {
		if value == nil {
			return nil
		}

		return fmt.Sprintf("%v", value)
	}

	return argument
}

func NewNumberArgument(name string) *Argument {
	argument := NewCustomArgument(name, func(value any) any {
		return nil
	})

	argument.handleValue = func(value any) any {
		if value == nil {
			return nil
		}

		raw := fmt.Sprintf("%v", value)

		result, err := strconv.ParseFloat(raw, 64)

		if err != nil {
			panic(fmt.Sprintf("Argument '%s' Must Be A Number", argument.name))
		}

		return result
	}

	return argument
}

func NewBooleanArgument(name string) *Argument {
	argument := NewCustomArgument(name, func(value any) any {
		return nil
	})

	argument.handleValue = func(value any) any {
		if value == nil {
			return nil
		}

		raw := fmt.Sprintf("%v", value)

		if raw == name {
			return true
		}

		result, err := strconv.ParseBool(strings.ToLower(raw))

		if err != nil {
			panic(fmt.Sprintf("Argument '%s' Must Be A Boolean", argument.name))
		}

		return result
	}

	return argument
}

func NewEnumArgument(name string, enum []string) *Argument {
	argument := NewCustomArgument(name, func(value any) any {
		return nil
	})

	if len(enum) == 0 {
		panic(fmt.Sprintf("Argument '%s' Of Type Enum Must Have At Least 1 Enum Option", argument.name))
	}

	argument.enum = enum

	argument.handleValue = func(value any) any {
		if value == nil {
			return nil
		}

		if len(argument.enum) == 0 {
			panic(fmt.Sprintf("Argument '%s' Of Type Enum Must Have At Least 1 Enum Option", argument.name))
		}

		raw := fmt.Sprintf("%v", value)

		exists := slices.Contains(argument.enum, raw)

		if !exists {
			panic(fmt.Sprintf("Argument '%s' Must Be One Of %s", argument.name, strings.Join(argument.enum, ", ")))
		}

		return value
	}

	return argument
}

func NewListArgument(name string) *Argument {
	argument := NewCustomArgument(name, func(value any) any {
		return nil
	})

	argument.variadic = true

	argument.handleValue = func(value any) any {
		if value == nil {
			return nil
		}

		result, ok := value.([]string)

		if !ok {
			panic(fmt.Sprintf("Argument '%s' Must Be A List Of Values", argument.name))
		}

		return result
	}

	return argument
}

func NewCustomArgument(name string, handleValue func(value any) any) *Argument {
	name = strings.ToLower(name)
	name = strings.Trim(name, "")

	if handleValue == nil {
		panic(fmt.Sprintf("Argument '%s' Must Have A Value Handler", name))
	}

	return &Argument{
		name:        name,
		tagline:     "An argument.",
		description: "An argument.",
		required:    true,
		variadic:    false,
		value:       nil,
		enum:        nil,
		handleValue: nil,
	}
}

func (argument *Argument) Name() string {
	return argument.name
}

func (argument *Argument) Tagline() string {
	return argument.tagline
}

func (argument *Argument) SetTagline(tagline string) *Argument {
	argument.tagline = tagline
	return argument
}

func (argument *Argument) Description() string {
	return argument.description
}

func (argument *Argument) SetDescription(description string) *Argument {
	argument.description = description
	return argument
}

func (argument *Argument) Optional() *Argument {
	argument.required = false
	return argument
}

func (argument *Argument) Required() bool {
	return argument.required
}

func (argument *Argument) Variadic() bool {
	return argument.variadic
}

func (argument *Argument) Value() any {
	return argument.handleValue(argument.value)
}

func (argument *Argument) SetValue(value any) *Argument {
	argument.value = value
	return argument
}

func (argument *Argument) Enum() []string {
	return argument.enum
}

func (argument *Argument) HandleValue(handleValue func(value any) any) *Argument {
	if handleValue == nil {
		panic(fmt.Sprintf("Argument '%s' Must Have A Value Handler", argument.name))
	}

	argument.handleValue = handleValue
	return argument
}
