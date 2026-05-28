package void

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	argumentNameRegexp = regexp.MustCompile(`^[a-z][a-z0-9]*$`)
	flagNameRegexp     = regexp.MustCompile(`^(?:--)?([a-z][a-z0-9]*(?:[-_][a-z0-9]+)*)$`)
)

func NewCustomArgument(name string, convert func(argument argumentView, value any) any) *Argument {
	name = mustParseArgumentName(name)

	return &Argument{
		name:        name,
		tagline:     "",
		description: "",
		required:    true,
		variadic:    false,
		enumerable:  false,
		value:       nil,
		enum:        nil,
		convert:     convert,
		validate:    nil,
	}
}

func NewCustomListArgument(name string, convert func(argument argumentView, value any) any) *Argument {
	argument := NewCustomArgument(name, convert)
	argument.variadic = true
	return argument
}

func NewCustomEnumArgument[E any](name string, enum []E, convert func(argument argumentView, value any) any) *Argument {
	argument := NewCustomArgument(name, convert)

	for i := range enum {
		argument.enum = append(argument.enum, enum[i])
	}

	argument.enumerable = true
	return argument
}

func NewCustomFlag(name string, convert func(flag flagView, value any) any) *Flag {
	name = mustParseFlagName(name)

	return &Flag{
		name:        name,
		tagline:     "",
		description: "",
		required:    false,
		variadic:    false,
		enumerable:  false,
		value:       nil,
		enum:        nil,
		convert:     convert,
		validate:    nil,
	}
}

func NewCustomListFlag(name string, convert func(flag flagView, value any) any) *Flag {
	flag := NewCustomFlag(name, convert)
	flag.variadic = true
	return flag
}

func NewCustomEnumFlag[E any](name string, enum []E, convert func(flag flagView, value any) any) *Flag {
	flag := NewCustomFlag(name, convert)

	for i := range enum {
		flag.enum = append(flag.enum, enum[i])
	}

	flag.enumerable = true
	return flag
}

func mustParseArgumentName(name string) string {
	name = strings.ToLower(strings.TrimSpace(name))

	if argumentNameRegexp.MatchString(name) {
		return name
	}

	NewDefinitionError(ArgumentNameInvalid, nil).
		WithMessage(
			"Invalid argument name format detected.",
		).
		WithExpected(
			"Argument names must begin with a lowercase letter and may only contain lowercase letters and numbers.",
		).
		WithReceived(
			fmt.Sprintf("Received argument name: %s", name),
		).
		WithSuggestion(
			"Use a valid argument name such as 'path', 'target', or 'build1'.",
		).
		Fatal()

	return ""
}

func mustParseFlagName(name string) string {
	name = strings.ToLower(strings.TrimSpace(name))

	if matches := flagNameRegexp.FindStringSubmatch(name); len(matches) != 0 {
		return "--" + matches[1]
	}

	NewDefinitionError(FlagNameInvalid, nil).
		WithMessage(
			"Invalid flag name format detected.",
		).
		WithExpected(
			"Flag names must optionally begin with '--', start with a lowercase letter, and may contain lowercase letters, numbers, hyphens, or underscores.",
		).
		WithReceived(
			fmt.Sprintf("Invalid flag name: %s", name),
		).
		WithSuggestion(
			"Use a valid flag name such as '--verbose', '--output-dir', or '--cache_size'.",
		).
		Fatal()

	return ""
}
