package void

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

const (
	InvalidInteger       ErrorCode = "INVALID_INTEGER"
	InvalidIntegerList   ErrorCode = "INVALID_INTEGER_LIST"
	InvalidIntegerOption ErrorCode = "INVALID_INTEGER_OPTION"
)

func NewIntArgument(name string) *Argument {
	return NewCustomArgument(name, func(argument argumentView, value any) any {
		switch v := value.(type) {
		case nil, int:
			return v

		case string:
			if integer, err := strconv.Atoi(v); err == nil {
				return integer
			}

			panic(
				NewUsageError(InvalidInteger, nil).
					WithArgument(argument).
					WithMessage(
						"Invalid integer argument detected.",
					).
					WithExpected(
						"Expected a valid integer value.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %s", v),
					).
					WithSuggestion(
						"Provide a valid integer value.",
					),
			)

		default:
			panic(
				NewUsageError(InvalidInteger, nil).
					WithArgument(argument).
					WithMessage(
						"Invalid integer argument detected.",
					).
					WithExpected(
						"Expected a valid integer value.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid integer value.",
					),
			)
		}
	})
}

func NewIntListArgument(name string) *Argument {
	return NewCustomListArgument(name, func(argument argumentView, value any) any {
		switch v := value.(type) {
		case nil, []int:
			return v

		case []string:
			integers := make([]int, 0)

			for i := range v {
				integer, err := strconv.Atoi(v[i])

				if err != nil {
					panic(
						NewUsageError(InvalidIntegerList, nil).
							WithArgument(argument).
							WithMessage(
								"Invalid integer list argument detected.",
							).
							WithExpected(
								"Expected a list of valid integer values.",
							).
							WithReceived(
								fmt.Sprintf("Invalid integer value: %s", v[i]),
							).
							WithSuggestion(
								"Provide a valid list of integer values.",
							),
					)
				}

				integers = append(integers, integer)
			}

			return integers

		default:
			panic(
				NewUsageError(InvalidIntegerList, nil).
					WithArgument(argument).
					WithMessage(
						"Invalid integer list argument detected.",
					).
					WithExpected(
						"Expected a list of valid integer values.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid list of integer values.",
					),
			)
		}
	})
}

func NewIntEnumArgument(name string, enum []int) *Argument {
	return NewCustomEnumArgument(name, enum, func(argument argumentView, value any) any {
		if value == nil {
			return nil
		}

		options := make([]string, 0)

		for i := range enum {
			options = append(options, fmt.Sprintf("%v", enum[i]))
		}

		switch v := value.(type) {
		case int, string:
			raw := fmt.Sprintf("%v", v)

			integer, err := strconv.Atoi(raw)

			if err != nil {
				panic(
					NewUsageError(InvalidIntegerOption, nil).
						WithArgument(argument).
						WithMessage(
							"Invalid integer enum argument detected.",
						).
						WithExpected(
							"Expected a valid integer value.",
						).
						WithReceived(
							fmt.Sprintf("Received value: %s", raw),
						).
						WithSuggestion(
							"Provide a valid integer value.",
						),
				)
			}

			if slices.Contains(options, raw) {
				return integer
			}

			panic(
				NewUsageError(InvalidIntegerOption, nil).
					WithArgument(argument).
					WithMessage(
						"Unsupported integer enum argument detected.",
					).
					WithExpected(
						fmt.Sprintf("Expected one of: %s", strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %d", integer),
					).
					WithSuggestion(
						"Provide one of the supported enum integer values.",
					),
			)

		default:
			panic(
				NewUsageError(InvalidIntegerOption, nil).
					WithArgument(argument).
					WithMessage(
						"Unsupported integer enum argument detected.",
					).
					WithExpected(
						fmt.Sprintf("Argument '%s' must receive one of: %s", argument.Name(), strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide one of the supported enum integer values.",
					),
			)
		}
	})
}

func NewIntFlag(name string) *Flag {
	return NewCustomFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil, int:
			return v

		case string:
			if integer, err := strconv.Atoi(v); err == nil {
				return integer
			}

			panic(
				NewUsageError(InvalidInteger, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid integer flag detected.",
					).
					WithExpected(
						"Expected a valid integer value.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %s", v),
					).
					WithSuggestion(
						"Provide a valid integer value.",
					),
			)

		default:
			panic(
				NewUsageError(InvalidInteger, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid integer flag detected.",
					).
					WithExpected(
						"Expected a valid integer value.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid integer value.",
					),
			)
		}
	})
}

func NewIntListFlag(name string) *Flag {
	return NewCustomListFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil, []int:
			return v

		case []string:
			integers := make([]int, 0)

			for i := range v {
				integer, err := strconv.Atoi(v[i])

				if err != nil {
					panic(
						NewUsageError(InvalidIntegerList, nil).
							WithFlag(flag).
							WithMessage(
								"Invalid integer list flag detected.",
							).
							WithExpected(
								"Expected a list of valid integer values.",
							).
							WithReceived(
								fmt.Sprintf("Invalid integer value: %s", v[i]),
							).
							WithSuggestion(
								"Provide a valid list of integer values.",
							),
					)
				}

				integers = append(integers, integer)
			}

			return integers

		default:
			panic(
				NewUsageError(InvalidIntegerList, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid integer list flag detected.",
					).
					WithExpected(
						"Expected a list of valid integer values.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid list of integer values.",
					),
			)
		}
	})
}

func NewIntEnumFlag(name string, enum []int) *Flag {
	return NewCustomEnumFlag(name, enum, func(flag flagView, value any) any {
		if value == nil {
			return nil
		}

		options := make([]string, 0)

		for i := range enum {
			options = append(options, fmt.Sprintf("%v", enum[i]))
		}

		switch v := value.(type) {
		case int, string:
			raw := fmt.Sprintf("%v", v)

			integer, err := strconv.Atoi(raw)

			if err != nil {
				panic(
					NewUsageError(InvalidIntegerOption, nil).
						WithFlag(flag).
						WithMessage(
							"Invalid integer enum flag detected.",
						).
						WithExpected(
							"Expected a valid integer value.",
						).
						WithReceived(
							fmt.Sprintf("Received value: %s", raw),
						).
						WithSuggestion(
							"Provide a valid integer value.",
						),
				)
			}

			if slices.Contains(options, raw) {
				return integer
			}

			panic(
				NewUsageError(InvalidIntegerOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported integer enum flag detected.",
					).
					WithExpected(
						fmt.Sprintf("Expected one of: %s", strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %d", integer),
					).
					WithSuggestion(
						"Provide one of the supported enum integer values.",
					),
			)

		default:
			panic(
				NewUsageError(InvalidIntegerOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported integer enum flag detected.",
					).
					WithExpected(
						fmt.Sprintf("Flag '%s' must receive one of: %s", flag.Name(), strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide one of the supported enum integer values.",
					),
			)
		}
	})
}
