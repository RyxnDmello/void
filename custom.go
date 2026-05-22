package void

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	argNameRegexp = regexp.MustCompile(`^[a-zA-Z]+$`)
	flgNameRegexp = regexp.MustCompile(`^(?:--)?([a-zA-Z]+)$`)
)

func NewCustomArgument(name string, handleConvert func(value any) any) *Argument {
	name = MustParseArgumentName(name)

	if handleConvert == nil {
		panic(fmt.Sprintf("Argument '%s' Must Have A Type Handler", name))
	}

	return &Argument{
		name:        name,
		tagline:     "",
		description: "",
		required:    true,
		variadic:    false,
		value:       nil,
		enum:        nil,
		covert:      handleConvert,
		validate:    nil,
	}
}

func NewCustomFlag(name string, handleConvert func(value any) any) *Flag {
	name = MustParseFlagName(name)

	if handleConvert == nil {
		panic(fmt.Sprintf("Flag '%s' Must Have A Type Handler", name))
	}

	return &Flag{
		name:        name,
		tagline:     "",
		description: "",
		required:    false,
		variadic:    false,
		value:       nil,
		enum:        nil,
		covert:      handleConvert,
		validate:    nil,
	}
}

func MustParseArgumentName(name string) string {
	name = strings.ToLower(strings.TrimSpace(name))

	if argNameRegexp.MatchString(name) {
		return name
	}

	panic(fmt.Sprintf("Invalid Argument Name: '%s'", name))
}

func MustParseFlagName(name string) string {
	name = strings.ToLower(strings.TrimSpace(name))

	if matches := flgNameRegexp.FindStringSubmatch(name); len(matches) != 0 {
		return "--" + matches[1]
	}

	panic(fmt.Sprintf("Invalid Flag Name: '%s'", name))
}
