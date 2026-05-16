package void

import (
	"fmt"
	"strings"
)

type Argument struct {
	name        string
	tagline     string
	description string
	required    bool
	variadic    bool
	value       any
	enum        []any
	covert      func(value any) any
	validate    func(value any) any
}

func NewArgument(name string) *Argument {
	return NewStringArgument(name)
}

func NewListArgument(name string) *Argument {
	return NewStringListArgument(name)
}

func NewEnumArgument(name string, enum []string) *Argument {
	return NewStringEnumArgument(name, enum)
}

func (argument *Argument) SetTagline(tagline string) *Argument {
	argument.tagline = strings.Trim(tagline, " ")
	return argument
}

func (argument *Argument) SetDescription(description string) *Argument {
	argument.description = strings.Trim(description, " ")
	return argument
}

func (argument *Argument) HandleConvert(convert func(value any) any) *Argument {
	if convert == nil {
		panic(fmt.Sprintf("Argument '%s' Must Have A Type Handler", argument.name))
	}

	argument.covert = func(value any) any {
		result := convert(value)

		if argument.validate != nil {
			return argument.validate(result)
		}

		return result
	}

	return argument
}

func (argument *Argument) HandleValidate(validate func(value any) any) *Argument {
	argument.validate = validate
	return argument
}

func (argument *Argument) MarkOptional(value any) *Argument {
	argument.value = value
	argument.required = false
	return argument
}

func (argument *Argument) getValue() any {
	return argument.covert(argument.value)
}

func (argument *Argument) setValue(value any) *Argument {
	if value == nil {
		panic(fmt.Sprintf("Argument '%s' Must Either Be Assigned A Value Or Set As Optional", argument.name))
	}

	argument.value = value
	return argument
}

func (argument *Argument) mustValidate() {
	if len(argument.name) == 0 {
		panic("Every Argument Must Have A Name")
	}

	if len(argument.tagline) == 0 {
		panic(fmt.Sprintf("Argument '%s' Must Have A Tagline", argument.name))
	}

	if len(argument.description) == 0 {
		panic(fmt.Sprintf("Argument '%s' Must Have A Tagline", argument.name))
	}

	if argument.covert == nil {
		panic(fmt.Sprintf("Argument '%s' Must Have A Type Conversion Handler", argument.name))
	}
}
