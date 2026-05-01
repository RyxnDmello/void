package argument

import (
	"fmt"
	"regexp"
	"strings"
)

var nameRegexp = regexp.MustCompile(`^[a-zA-Z]+$`)

func NewCustomArgument(name string, convert func(value any) any) *Argument {
	name = MustParseArgumentName(name)

	if convert == nil {
		panic(fmt.Sprintf("Argument '%s' Must Have A Type Conversion Handler", name))
	}

	return &Argument{
		name:        name,
		tagline:     "",
		description: "",
		required:    true,
		variadic:    false,
		value:       nil,
		enum:        nil,
		covert:      nil,
		validate:    nil,
	}
}

func MustParseArgumentName(name string) string {
	name = strings.ToLower(strings.TrimSpace(name))

	if valid := nameRegexp.MatchString(name); !valid {
		panic(fmt.Sprintf("Argument Cannot Have A Name Such As '%s'", name))
	}

	return name
}
