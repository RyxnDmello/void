package void

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

const (
	InvalidFloat       ErrorCode = "INVALID_FLOAT"
	InvalidFloatList   ErrorCode = "INVALID_FLOAT_LIST"
	InvalidFloatOption ErrorCode = "INVALID_FLOAT_OPTION"
)

func NewFloatArgument(name string) *Argument {
	return NewCustomArgument(name, func(argument argumentView, value any) any {
		switch v := value.(type) {
		case nil, float64:
			return v

		case int, float32, string:
			if float, err := strconv.ParseFloat(fmt.Sprintf("%v", v), 64); err == nil {
				return float
			}

			panic(
				NewUsageError(InvalidFloat, nil).
					WithArgument(argument).
					WithMessage(
						"Invalid float argument detected.",
					).
					WithExpected(
						"Expected a valid float value.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid float value.",
					),
			)

		default:
			panic(
				NewUsageError(InvalidFloat, nil).
					WithArgument(argument).
					WithMessage(
						"Invalid float argument detected.",
					).
					WithExpected(
						"Expected a valid float value.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid float value.",
					),
			)
		}
	})
}

func NewFloatListArgument(name string) *Argument {
	return NewCustomListArgument(name, func(argument argumentView, value any) any {
		switch v := value.(type) {
		case nil, []float64:
			return v

		case []int:
			floats := make([]float64, 0)

			for i := range v {
				float, err := strconv.ParseFloat(fmt.Sprintf("%v", v[i]), 64)

				if err != nil {
					panic(
						NewUsageError(InvalidFloatList, nil).
							WithArgument(argument).
							WithMessage(
								"Invalid float list argument detected.",
							).
							WithExpected(
								"Expected a list of valid float values.",
							).
							WithReceived(
								fmt.Sprintf("Invalid float value: %v", v[i]),
							).
							WithSuggestion(
								"Provide a valid list of float values.",
							),
					)
				}

				floats = append(floats, float)
			}

			return floats

		case []float32:
			floats := make([]float64, 0)

			for i := range v {
				float, err := strconv.ParseFloat(fmt.Sprintf("%v", v[i]), 64)

				if err != nil {
					panic(
						NewUsageError(InvalidFloatList, nil).
							WithArgument(argument).
							WithMessage(
								"Invalid float list argument detected.",
							).
							WithExpected(
								"Expected a list of valid float values.",
							).
							WithReceived(
								fmt.Sprintf("Invalid float value: %v", v[i]),
							).
							WithSuggestion(
								"Provide a valid list of float values.",
							),
					)
				}

				floats = append(floats, float)
			}

			return floats

		case []string:
			floats := make([]float64, 0)

			for i := range v {
				float, err := strconv.ParseFloat(v[i], 64)

				if err != nil {
					panic(
						NewUsageError(InvalidFloatList, nil).
							WithArgument(argument).
							WithMessage(
								"Invalid float list argument detected.",
							).
							WithExpected(
								"Expected a list of valid float values.",
							).
							WithReceived(
								fmt.Sprintf("Invalid float value: %s", v[i]),
							).
							WithSuggestion(
								"Provide a valid list of float values.",
							),
					)
				}

				floats = append(floats, float)
			}

			return floats

		default:
			panic(
				NewUsageError(InvalidFloatList, nil).
					WithArgument(argument).
					WithMessage(
						"Invalid float list argument detected.",
					).
					WithExpected(
						"Expected a list of valid float values.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid list of float values.",
					),
			)
		}
	})
}

func NewFloatEnumArgument(name string, enum []float64) *Argument {
	return NewCustomEnumArgument(name, enum, func(argument argumentView, value any) any {
		if value == nil {
			return nil
		}

		options := make([]string, 0)

		for i := range enum {
			options = append(options, fmt.Sprintf("%v", enum[i]))
		}

		switch v := value.(type) {
		case int, float32, float64, string:
			raw := fmt.Sprintf("%v", v)

			float, err := strconv.ParseFloat(raw, 64)

			if err != nil {
				panic(
					NewUsageError(InvalidFloatOption, nil).
						WithArgument(argument).
						WithMessage(
							"Invalid float enum argument detected.",
						).
						WithExpected(
							"Expected a valid float value.",
						).
						WithReceived(
							fmt.Sprintf("Received value: %s", raw),
						).
						WithSuggestion(
							"Provide a valid float value.",
						),
				)
			}

			if slices.Contains(options, raw) {
				return float
			}

			panic(
				NewUsageError(InvalidFloatOption, nil).
					WithArgument(argument).
					WithMessage(
						"Unsupported float enum argument detected.",
					).
					WithExpected(
						fmt.Sprintf("Expected one of: %s", strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %f", float),
					).
					WithSuggestion(
						"Provide one of the supported enum float values.",
					),
			)

		default:
			panic(
				NewUsageError(InvalidFloatOption, nil).
					WithArgument(argument).
					WithMessage(
						"Unsupported float enum argument detected.",
					).
					WithExpected(
						fmt.Sprintf("Argument '%s' must receive one of: %s", argument.Name(), strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide one of the supported enum float values.",
					),
			)
		}
	})
}

func NewFloatFlag(name string) *Flag {
	return NewCustomFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil, float64:
			return v

		case int, float32, string:
			if float, err := strconv.ParseFloat(fmt.Sprintf("%v", v), 64); err == nil {
				return float
			}

			panic(
				NewUsageError(InvalidFloat, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid float flag detected.",
					).
					WithExpected(
						"Expected a valid float value.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid float value.",
					),
			)

		default:
			panic(
				NewUsageError(InvalidFloat, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid float flag detected.",
					).
					WithExpected(
						"Expected a valid float value.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid float value.",
					),
			)
		}
	})
}

func NewFloatListFlag(name string) *Flag {
	return NewCustomListFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil, []float64:
			return v

		case []int:
			floats := make([]float64, 0)

			for i := range v {
				float, err := strconv.ParseFloat(fmt.Sprintf("%v", v[i]), 64)

				if err != nil {
					panic(
						NewUsageError(InvalidFloatList, nil).
							WithFlag(flag).
							WithMessage(
								"Invalid float list flag detected.",
							).
							WithExpected(
								"Expected a list of valid float values.",
							).
							WithReceived(
								fmt.Sprintf("Invalid float value: %v", v[i]),
							).
							WithSuggestion(
								"Provide a valid list of float values.",
							),
					)
				}

				floats = append(floats, float)
			}

			return floats

		case []float32:
			floats := make([]float64, 0)

			for i := range v {
				float, err := strconv.ParseFloat(fmt.Sprintf("%v", v[i]), 64)

				if err != nil {
					panic(
						NewUsageError(InvalidFloatList, nil).
							WithFlag(flag).
							WithMessage(
								"Invalid float list flag detected.",
							).
							WithExpected(
								"Expected a list of valid float values.",
							).
							WithReceived(
								fmt.Sprintf("Invalid float value: %v", v[i]),
							).
							WithSuggestion(
								"Provide a valid list of float values.",
							),
					)
				}

				floats = append(floats, float)
			}

			return floats

		case []string:
			floats := make([]float64, 0)

			for i := range v {
				float, err := strconv.ParseFloat(v[i], 64)

				if err != nil {
					panic(
						NewUsageError(InvalidFloatList, nil).
							WithFlag(flag).
							WithMessage(
								"Invalid float list flag detected.",
							).
							WithExpected(
								"Expected a list of valid float values.",
							).
							WithReceived(
								fmt.Sprintf("Invalid float value: %s", v[i]),
							).
							WithSuggestion(
								"Provide a valid list of float values.",
							),
					)
				}

				floats = append(floats, float)
			}

			return floats

		default:
			panic(
				NewUsageError(InvalidFloatList, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid float list flag detected.",
					).
					WithExpected(
						"Expected a list of valid float values.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid list of float values.",
					),
			)
		}
	})
}

func NewFloatEnumFlag(name string, enum []float64) *Flag {
	return NewCustomEnumFlag(name, enum, func(flag flagView, value any) any {
		if value == nil {
			return nil
		}

		options := make([]string, 0)

		for i := range enum {
			options = append(options, fmt.Sprintf("%v", enum[i]))
		}

		switch v := value.(type) {
		case int, float32, float64, string:
			raw := fmt.Sprintf("%v", v)

			float, err := strconv.ParseFloat(raw, 64)

			if err != nil {
				panic(
					NewUsageError(InvalidFloatOption, nil).
						WithFlag(flag).
						WithMessage(
							"Invalid float enum flag detected.",
						).
						WithExpected(
							"Expected a valid float value.",
						).
						WithReceived(
							fmt.Sprintf("Received value: %s", raw),
						).
						WithSuggestion(
							"Provide a valid float value.",
						),
				)
			}

			if slices.Contains(options, raw) {
				return float
			}

			panic(
				NewUsageError(InvalidFloatOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported float enum flag detected.",
					).
					WithExpected(
						fmt.Sprintf("Expected one of: %s", strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %f", float),
					).
					WithSuggestion(
						"Provide one of the supported enum float values.",
					),
			)

		default:
			panic(
				NewUsageError(InvalidFloatOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported float enum flag detected.",
					).
					WithExpected(
						fmt.Sprintf("Flag '%s' must receive one of: %s", flag.Name(), strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide one of the supported enum float values.",
					),
			)
		}
	})
}
