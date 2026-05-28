package void

import (
	"fmt"
	"slices"
	"strings"
)

const (
	InvalidString       ErrorCode = "INVALID_STRING_LIST"
	InvalidStringList   ErrorCode = "INVALID_STRING_LIST"
	InvalidStringOption ErrorCode = "INVALID_STRING_OPTION"
)

func NewStringArgument(name string) *Argument {
	return NewCustomArgument(name, func(_ argumentView, value any) any {
		switch v := value.(type) {
		case nil:
			return v

		default:
			return fmt.Sprintf("%v", v)
		}
	})
}

func NewStringListArgument(name string) *Argument {
	return NewCustomListArgument(name, func(argument argumentView, value any) any {
		switch v := value.(type) {
		case nil, []string:
			return v

		default:
			panic(
				NewUsageError(InvalidStringList, nil).
					WithArgument(argument).
					WithMessage(
						"Invalid string list argument detected.",
					).
					WithExpected(
						"Expected a list of string values.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid list of string values.",
					),
			)
		}
	})
}

func NewStringEnumArgument(name string, enum []string) *Argument {
	return NewCustomEnumArgument(name, enum, func(argument argumentView, value any) any {
		switch v := value.(type) {
		case nil:
			return v

		case string:
			if !slices.Contains(enum, v) {
				panic(
					NewUsageError(InvalidStringOption, nil).
						WithArgument(argument).
						WithMessage(
							"Unsupported string argument detected.",
						).
						WithExpected(
							fmt.Sprintf("Expected one of: %s", strings.Join(enum, ", ")),
						).
						WithReceived(
							fmt.Sprintf("Received value: %s", v),
						).
						WithSuggestion(
							"Provide one of the supported enum string values.",
						),
				)
			}

			return v

		default:
			panic(
				NewUsageError(InvalidStringOption, nil).
					WithArgument(argument).
					WithMessage(
						"Unsupported string argument detected.",
					).
					WithExpected(
						fmt.Sprintf("Argument '%s' must receive one of: %s", argument.Name(), strings.Join(enum, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide one of the supported enum string values.",
					),
			)
		}
	})
}

func NewStringFlag(name string) *Flag {
	return NewCustomFlag(name, func(_ flagView, value any) any {
		switch v := value.(type) {
		case nil:
			return v

		default:
			return fmt.Sprintf("%v", v)
		}
	})
}

func NewStringListFlag(name string) *Flag {
	return NewCustomListFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil, []string:
			return v

		default:
			panic(
				NewUsageError(InvalidString, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid string list flag detected.",
					).
					WithExpected(
						"Expected a list of string values.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid list of string values.",
					),
			)
		}
	})
}

func NewStringEnumFlag(name string, enum []string) *Flag {
	return NewCustomEnumFlag(name, enum, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil:
			return v

		case string:
			if !slices.Contains(enum, v) {
				panic(
					NewUsageError(InvalidStringOption, nil).
						WithFlag(flag).
						WithMessage(
							"Unsupported string flag value detected.",
						).
						WithExpected(
							fmt.Sprintf("Expected one of: %s", strings.Join(enum, ", ")),
						).
						WithReceived(
							fmt.Sprintf("Received value: %s", v),
						).
						WithSuggestion(
							"Provide one of the supported enum string values.",
						),
				)
			}

			return v

		default:
			panic(
				NewUsageError(InvalidStringOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported string flag value detected.",
					).
					WithExpected(
						fmt.Sprintf("Flag '%s' must receive one of: %s", flag.Name(), strings.Join(enum, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide one of the supported enum string values.",
					),
			)
		}
	})
}
