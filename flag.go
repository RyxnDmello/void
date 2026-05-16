package void

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

func (flag *Flag) SetTagline(tagline string) *Flag {
	flag.tagline = strings.Trim(tagline, " ")
	return flag
}

func (flag *Flag) SetDescription(description string) *Flag {
	flag.description = strings.Trim(description, " ")
	return flag
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

func (flag *Flag) HandleConvert(convert func(value any) any) *Flag {
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

func (flag *Flag) HandleValidate(validate func(value any) any) *Flag {
	flag.validate = validate
	return flag
}

func (flag *Flag) getValue() any {
	return flag.covert(flag.value)
}

func (flag *Flag) setValue(value any) *Flag {
	if value == nil {
		panic(fmt.Sprintf("Flag '%s' Must Either Be Assigned A Value Or Set As Optional", flag.name))
	}

	flag.value = value
	return flag
}

func (flag *Flag) mustValidate() {
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
