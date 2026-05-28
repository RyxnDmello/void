package void

import (
	"fmt"
	"strconv"
)

const (
	InvalidBoolean     ErrorCode = "INVALID_BOOLEAN"
	InvalidBooleanList ErrorCode = "INVALID_BOOLEAN_LIST"
)

func NewBoolArgument(name string) *Argument {
	argument := NewCustomArgument(name, func(argument argumentView, value any) any {
		switch v := value.(type) {
		case nil, bool:
			return v

		case string:
			if boolean, err := strconv.ParseBool(v); err == nil {
				return boolean
			}

			panic(
				NewUsageError(InvalidBoolean, nil).
					WithArgument(argument).
					WithMessage(
						"Invalid boolean argument detected.",
					).
					WithExpected(
						"Expected a valid boolean value.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %s", v),
					).
					WithSuggestion(
						"Provide a valid boolean value such as 'true' or 'false'.",
					),
			)

		default:
			panic(
				NewUsageError(InvalidBoolean, nil).
					WithArgument(argument).
					WithMessage(
						"Invalid boolean argument detected.",
					).
					WithExpected(
						"Expected a valid boolean value.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid boolean value such as 'true' or 'false'.",
					),
			)
		}
	})

	argument.value = false
	return argument
}

func NewBoolListArgument(name string) *Argument {
	return NewCustomListArgument(name, func(argument argumentView, value any) any {
		switch v := value.(type) {
		case nil, []bool:
			return v

		case []string:
			booleans := make([]bool, 0)

			for i := range v {
				boolean, err := strconv.ParseBool(v[i])

				if err != nil {
					panic(
						NewUsageError(InvalidBooleanList, nil).
							WithArgument(argument).
							WithMessage(
								"Invalid boolean list argument detected.",
							).
							WithExpected(
								"Expected a list of valid boolean values.",
							).
							WithReceived(
								fmt.Sprintf("Invalid boolean value: %s", v[i]),
							).
							WithSuggestion(
								"Provide valid boolean values such as 'true' or 'false'.",
							),
					)
				}

				booleans = append(booleans, boolean)
			}

			return booleans

		default:
			panic(
				NewUsageError(InvalidBooleanList, nil).
					WithArgument(argument).
					WithMessage(
						"Invalid boolean list argument detected.",
					).
					WithExpected(
						"Expected a list of valid boolean values.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid list of boolean values.",
					),
			)
		}
	})
}

func NewBoolFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil, bool:
			return v

		case string:
			if boolean, err := strconv.ParseBool(v); err == nil {
				return boolean
			}

			panic(
				NewUsageError(InvalidBoolean, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid boolean flag detected.",
					).
					WithExpected(
						"Expected a valid boolean value.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %s", v),
					).
					WithSuggestion(
						"Provide a valid boolean value such as 'true' or 'false'.",
					),
			)

		default:
			panic(
				NewUsageError(InvalidBoolean, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid boolean flag detected.",
					).
					WithExpected(
						"Expected a valid boolean value.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid boolean value such as 'true' or 'false'.",
					),
			)
		}
	})

	flag.value = false
	return flag
}

func NewBoolListFlag(name string) *Flag {
	return NewCustomListFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil, []bool:
			return v

		case []string:
			booleans := make([]bool, 0)

			for i := range v {
				boolean, err := strconv.ParseBool(v[i])

				if err != nil {
					panic(
						NewUsageError(InvalidBooleanList, nil).
							WithFlag(flag).
							WithMessage(
								"Invalid boolean list flag detected.",
							).
							WithExpected(
								"Expected a list of valid boolean values.",
							).
							WithReceived(
								fmt.Sprintf("Invalid boolean value: %s", v[i]),
							).
							WithSuggestion(
								"Provide valid boolean values such as 'true' or 'false'.",
							),
					)
				}

				booleans = append(booleans, boolean)
			}

			return booleans

		default:
			panic(
				NewUsageError(InvalidBooleanList, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid boolean list flag detected.",
					).
					WithExpected(
						"Expected a list of valid boolean values.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid list of boolean values.",
					),
			)
		}
	})
}
