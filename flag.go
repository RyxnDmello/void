package void

import (
	"fmt"
	"strings"
)

type flagView interface {
	Name() string
	Tagline() string
	Description() string
	Required() bool
	Variadic() bool
	Enumerable() (bool, []any)
}

type Flag struct {
	name        string
	tagline     string
	description string
	required    bool
	variadic    bool
	enumerable  bool
	value       any
	enum        []any
	convert     func(flag flagView, value any) any
	validate    func(flag flagView, value any) any
}

func NewFlag(name string) *Flag {
	return NewStringFlag(name)
}

func NewListFlag(name string) *Flag {
	return NewStringListFlag(name)
}

func NewMapFlag(name string) *Flag {
	return NewStringMapFlag(name)
}

func NewEnumFlag(name string, enum []string) *Flag {
	return NewStringEnumFlag(name, enum)
}

func (flag *Flag) Name() string {
	return flag.name
}

func (flag *Flag) Tagline() string {
	return flag.tagline
}

func (flag *Flag) SetTagline(tagline string) *Flag {
	flag.tagline = strings.Trim(tagline, " ")
	return flag
}

func (flag *Flag) Description() string {
	return flag.description
}

func (flag *Flag) SetDescription(description string) *Flag {
	flag.description = strings.Trim(description, " ")
	return flag
}

func (flag *Flag) Required() bool {
	return flag.required
}

func (flag *Flag) Variadic() bool {
	return flag.variadic
}

func (flag *Flag) Enumerable() (bool, []any) {
	return flag.enumerable, flag.enum
}

func (flag *Flag) SetDefault(value any) *Flag {
	flag.value = value
	return flag
}

func (flag *Flag) MarkRequired() *Flag {
	flag.value = nil
	flag.required = true
	return flag
}

func (flag *Flag) HandleConvert(convert func(flag flagView, value any) any) *Flag {
	if convert == nil {
		return nil
	}

	flag.convert = func(_ flagView, _ any) any {
		value := convert(flag, flag.value)

		if flag.validate == nil {
			return value
		}

		return flag.validate(flag, value)
	}

	return flag
}

func (flag *Flag) HandleValidate(validate func(flag flagView, value any) any) *Flag {
	flag.validate = validate
	return flag
}

func (flag *Flag) mustGetValue() any {
	return flag.convert(flag, flag.value)
}

func (flag *Flag) mustSetValue(value any) *Flag {
	if value == nil {
		panic(
			NewDefinitionError(FlagValueRequired, nil).
				WithFlag(flag).
				WithMessage(
					"Flag value is required.",
				).
				WithExpected(
					fmt.Sprintf("Flag '%s' must have a value or be marked as optional.", flag.name),
				).
				WithSuggestion(
					"Assign a value or mark the flag as required.",
				),
		)
	}

	flag.value = value
	return flag
}

func (flag *Flag) mustValidate() {
	if len(flag.name) == 0 {
		panic(
			NewDefinitionError(FlagNameRequired, nil).
				WithFlag(flag).
				WithMessage(
					"Flag name is required.",
				).
				WithExpected(
					"Every flag must define a name.",
				),
		)
	}

	if len(flag.tagline) == 0 {
		panic(
			NewDefinitionError(FlagTaglineRequired, nil).
				WithFlag(flag).
				WithMessage(
					"Flag tagline is required.",
				).
				WithExpected(
					fmt.Sprintf("Flag '%s' must define a tagline.", flag.name),
				).
				WithSuggestion(
					"Define a tagline using SetTagline().",
				),
		)
	}

	if len(flag.description) == 0 {
		panic(
			NewDefinitionError(FlagDescriptionRequired, nil).
				WithFlag(flag).
				WithMessage(
					"Flag description is required.",
				).
				WithExpected(
					fmt.Sprintf("Flag '%s' must define a description.", flag.name),
				).
				WithSuggestion(
					"Define a description using SetDescription().",
				),
		)
	}

	if flag.convert == nil {
		panic(
			NewDefinitionError(FlagTypeHandlerRequired, nil).
				WithFlag(flag).
				WithMessage(
					"Flag type handler is required.",
				).
				WithExpected(
					fmt.Sprintf("Flag '%s' must define a type handler.", flag.name),
				).
				WithSuggestion(
					"Attach a type handler to the flag definition.",
				),
		)
	}

	if flag.enumerable && len(flag.enum) == 0 {
		panic(
			NewDefinitionError(FlagEnumRequired, nil).
				WithFlag(flag).
				WithMessage(
					"Flag enum options are required.",
				).
				WithExpected(
					fmt.Sprintf("Enumerable flag '%s' must define at least one enum value.", flag.name),
				).
				WithSuggestion(
					"Provide one or more enum options for the flag definition.",
				),
		)
	}
}
