package void

import (
	"fmt"
	"strings"
)

type argumentView interface {
	Name() string
	Tagline() string
	Description() string
	Required() bool
	Variadic() bool
	Enumerable() (bool, []any)
}

type Argument struct {
	name        string
	tagline     string
	description string
	required    bool
	variadic    bool
	enumerable  bool
	value       any
	enum        []any
	convert     func(argument argumentView, value any) any
	validate    func(argument argumentView, value any) any
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

func (argument *Argument) Name() string {
	return argument.name
}

func (argument *Argument) Tagline() string {
	return argument.tagline
}

func (argument *Argument) SetTagline(tagline string) *Argument {
	argument.tagline = strings.Trim(tagline, " ")
	return argument
}

func (argument *Argument) Description() string {
	return argument.description
}

func (argument *Argument) SetDescription(description string) *Argument {
	argument.description = strings.Trim(description, " ")
	return argument
}

func (argument *Argument) Required() bool {
	return argument.required
}

func (argument *Argument) Variadic() bool {
	return argument.variadic
}

func (argument *Argument) Enumerable() (bool, []any) {
	return argument.enumerable, argument.enum
}

func (argument *Argument) MarkOptional(value any) *Argument {
	argument.value = value
	argument.required = false
	return argument
}

func (argument *Argument) HandleConvert(convert func(argument argumentView, value any) any) *Argument {
	if convert == nil {
		return nil
	}

	argument.convert = func(_ argumentView, _ any) any {
		value := convert(argument, argument.value)

		if argument.validate == nil {
			return value
		}

		return argument.validate(argument, value)
	}

	return argument
}

func (argument *Argument) HandleValidate(validate func(argument argumentView, value any) any) *Argument {
	argument.validate = validate
	return argument
}

func (argument *Argument) mustGetValue() any {
	return argument.convert(argument, argument.value)
}

func (argument *Argument) mustSetValue(value any) *Argument {
	if value == nil {
		panic(
			NewDefinitionError(ArgumentValueRequired, nil).
				WithArgument(argument).
				WithMessage(
					"Argument value is required.",
				).
				WithExpected(
					fmt.Sprintf("Argument '%s' must have a value or be marked as optional.", argument.name),
				).
				WithSuggestion(
					"Assign a value or mark the argument as optional.",
				),
		)
	}

	argument.value = value
	return argument
}

func (argument *Argument) mustValidate() {
	if len(argument.name) == 0 {
		panic(
			NewDefinitionError(ArgumentNameRequired, nil).
				WithArgument(argument).
				WithMessage(
					"Argument name is required.",
				).
				WithExpected(
					"Every argument must define a name.",
				),
		)
	}

	if len(argument.tagline) == 0 {
		panic(
			NewDefinitionError(ArgumentTaglineRequired, nil).
				WithArgument(argument).
				WithMessage(
					"Argument tagline is required.",
				).
				WithExpected(
					fmt.Sprintf("Argument '%s' must define a tagline.", argument.name),
				).
				WithSuggestion(
					"Define a tagline using SetTagline().",
				),
		)
	}

	if len(argument.description) == 0 {
		panic(
			NewDefinitionError(ArgumentDescriptionRequired, nil).
				WithArgument(argument).
				WithMessage(
					"Argument description is required.",
				).
				WithExpected(
					fmt.Sprintf("Argument '%s' must define a description.", argument.name),
				).
				WithSuggestion(
					"Define a description using SetDescription().",
				),
		)
	}

	if argument.convert == nil {
		panic(
			NewDefinitionError(ArgumentTypeHandlerRequired, nil).
				WithArgument(argument).
				WithMessage(
					"Argument type handler is required.",
				).
				WithExpected(
					fmt.Sprintf("Argument '%s' must define a type handler.", argument.name),
				).
				WithSuggestion(
					"Attach a type handler to the argument definition.",
				),
		)
	}

	if argument.enumerable && len(argument.enum) == 0 {
		panic(
			NewDefinitionError(ArgumentEnumRequired, nil).
				WithArgument(argument).
				WithMessage(
					"Argument enum options are required.",
				).
				WithExpected(
					fmt.Sprintf("Enumerable argument '%s' must define at least one enum value.", argument.name),
				).
				WithSuggestion(
					"Provide one or more enum options for the argument definition.",
				),
		)
	}
}
