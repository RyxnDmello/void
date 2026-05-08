package flag

import (
	"fmt"
	"strings"
)

type Flag struct {
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

func (flag *Flag) IsRequired() bool {
	return flag.required
}

func (flag *Flag) Required() *Flag {
	flag.value = nil
	flag.required = true
	return flag
}

func (flag *Flag) Default(value any) *Flag {
	flag.value = value
	return flag
}

func (flag *Flag) IsVariadic() bool {
	return flag.variadic
}

func (flag *Flag) Value() any {
	return flag.covert(flag.value)
}

func (flag *Flag) SetValue(value any) *Flag {
	if value == nil {
		panic(fmt.Sprintf("Flag '%s' Must Either Be Assigned A Value Or Set As Optional", flag.name))
	}

	flag.value = value
	return flag
}

func (flag *Flag) Enum() []any {
	return flag.enum
}

func (flag *Flag) SetConvert(convert func(value any) any) *Flag {
	if convert == nil {
		panic(fmt.Sprintf("Flag '%s' Must Have A Type Conversion Handler", flag.name))
	}

	flag.covert = func(value any) any {
		result := convert(value)

		if flag.validate != nil {
			return flag.validate(result)
		}

		return result
	}

	return flag
}

func (flag *Flag) SetValidate(validate func(value any) any) *Flag {
	flag.validate = validate
	return flag
}

func (flag *Flag) MustValidate() {
	if len(flag.name) == 0 {
		panic("Every Flag Must Have A Name")
	}

	if len(flag.tagline) == 0 {
		panic(fmt.Sprintf("Flag '%s' Must Have A Tagline", flag.name))
	}

	if len(flag.description) == 0 {
		panic(fmt.Sprintf("Flag '%s' Must Have A Tagline", flag.name))
	}

	if flag.covert == nil {
		panic(fmt.Sprintf("Flag '%s' Must Have A Type Conversion Handler", flag.name))
	}
}
