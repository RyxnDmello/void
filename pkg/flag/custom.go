package flag

import (
	"fmt"
	"regexp"
	"strings"
)

var nameRegexp = regexp.MustCompile(`^(?:--)?([a-zA-Z]+)$`)

func NewCustomFlag(name string, convert func(value any) any) *Flag {
	name = MustParseFlagName(name)

	if convert == nil {
		panic(fmt.Sprintf("Flag '%s' Must Have A Type Conversion Handler", name))
	}

	return &Flag{
		name:        name,
		tagline:     "",
		description: "",
		required:    false,
		variadic:    false,
		value:       nil,
		enum:        nil,
		covert:      nil,
		validate:    nil,
	}
}

func MustParseFlagName(name string) string {
	name = strings.ToLower(strings.TrimSpace(name))

	if matches := nameRegexp.FindStringSubmatch(name); len(matches) != 0 {
		return "--" + matches[1]
	}

	panic(fmt.Sprintf("Flag Cannot Have A Name Such As '%s'", name))
}
