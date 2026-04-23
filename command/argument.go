package command

import (
	"fmt"
	"strconv"
)

type Typed int

const (
	String Typed = iota
	Integer
)

type Argument struct {
	name        string
	tagline     string
	description string
	typed       Typed
	value       any
	required    bool
	onValue     func(value any, typed Typed) any
}

func NewArgument(typed Typed, name string) *Argument {
	argument := &Argument{
		name:        name,
		tagline:     "An argument.",
		description: "An argument.",
		required:    false,
		typed:       typed,
		value:       nil,
		onValue:     nil,
	}

	switch argument.typed {

	case Integer:
		argument.onValue = func(value any, typed Typed) any {
			result, err := strconv.Atoi(value.(string))

			if err != nil {
				panic(fmt.Sprintf("Argument '%s' Must Be An Integer", argument.name))
			}

			return result
		}

	default:
		argument.onValue = func(value any, typed Typed) any {
			return value.(string)
		}
	}

	return argument
}

func (argument *Argument) Name() string {
	return argument.name
}

func (argument *Argument) SetName(name string) *Argument {
	argument.name = name
	return argument
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

func (argument *Argument) Required() *Argument {
	argument.required = true
	return argument
}

func (argument *Argument) IsRequired() bool {
	return argument.required
}

func (argument *Argument) Type() string {
	switch argument.typed {

	case Integer:
		return "number"

	default:
		return "string"
	}
}

func (argument *Argument) Value() any {
	return argument.value
}

func (argument *Argument) SetValue(value any) *Argument {
	raw := value.(string)

	switch argument.typed {

	case Integer:
		result, err := strconv.Atoi(raw)

		if err != nil {
			panic(fmt.Sprintf("Argument '%s' Must Be An Integer", argument.name))
		}

		argument.value = result

	default:
		argument.value = raw
	}

	return argument
}

func (argument *Argument) OnValue(onValue func(value any, typed Typed) any) *Argument {
	argument.onValue = onValue
	return argument
}
