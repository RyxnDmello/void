package void

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

const (
	InvalidFloatMap   ErrorCode = "INVALID_FLOAT_MAP"
	InvalidStringMap  ErrorCode = "INVALID_STRING_MAP"
	InvalidIntegerMap ErrorCode = "INVALID_INTEGER_MAP"
	InvalidBooleanMap ErrorCode = "INVALID_BOOLEAN_MAP"
	InvalidComplexMap ErrorCode = "INVALID_COMPLEX_MAP"
)

func NewStringMapFlag(name string) *Flag {
	return NewCustomListFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil, map[string]string:
			return v

		case []string:
			record := make(map[string]string)

			for i := range v {
				parts := strings.Split(v[i], "=")

				if len(parts) != 2 {
					panic(
						NewUsageError(InvalidStringMap, nil).
							WithFlag(flag).
							WithMessage(
								"Invalid string map entry detected.",
							).
							WithExpected(
								"Expected map entries in the format 'key=value'.",
							).
							WithReceived(
								fmt.Sprintf("Received entry: %s", v[i]),
							).
							WithSuggestion(
								"Provide valid string map entries such as 'name=john'.",
							),
					)
				}

				record[parts[0]] = parts[1]
			}

			return record

		default:
			panic(
				NewUsageError(InvalidStringMap, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid string map detected.",
					).
					WithExpected(
						"Expected a map of string key-value pairs.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide valid string map entries in the format 'key=value'.",
					),
			)
		}
	})
}

func NewIntMapFlag(name string) *Flag {
	return NewCustomListFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil, map[string]int:
			return v

		case []string:
			record := make(map[string]int)

			for i := range v {
				parts := strings.Split(v[i], "=")

				if len(parts) != 2 {
					panic(
						NewUsageError(InvalidIntegerMap, nil).
							WithFlag(flag).
							WithMessage(
								"Invalid integer map entry detected.",
							).
							WithExpected(
								"Expected map entries in the format 'key=value'.",
							).
							WithReceived(
								fmt.Sprintf("Received entry: %s", v[i]),
							).
							WithSuggestion(
								"Provide valid integer map entries such as 'port=3000'.",
							),
					)
				}

				integer, err := strconv.Atoi(parts[1])

				if err != nil {
					panic(
						NewUsageError(InvalidIntegerMap, nil).
							WithFlag(flag).
							WithMessage(
								"Invalid integer map value detected.",
							).
							WithExpected(
								"Expected integer map values.",
							).
							WithReceived(
								fmt.Sprintf("Received value: %s", parts[1]),
							).
							WithSuggestion(
								"Provide valid integer values such as 'count=10'.",
							),
					)
				}

				record[parts[0]] = integer
			}

			return record

		default:
			panic(
				NewUsageError(InvalidIntegerMap, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid integer map detected.",
					).
					WithExpected(
						"Expected a map of integer key-value pairs.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide valid integer map entries in the format 'key=value'.",
					),
			)
		}
	})
}

func NewFloatMapFlag(name string) *Flag {
	return NewCustomListFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil, map[string]float64:
			return v

		case []string:
			record := make(map[string]float64)

			for i := range v {
				parts := strings.Split(v[i], "=")

				if len(parts) != 2 {
					panic(
						NewUsageError(InvalidFloatMap, nil).
							WithFlag(flag).
							WithMessage(
								"Invalid float map entry detected.",
							).
							WithExpected(
								"Expected map entries in the format 'key=value'.",
							).
							WithReceived(
								fmt.Sprintf("Received entry: %s", v[i]),
							).
							WithSuggestion(
								"Provide valid float map entries such as 'ratio=1.5'.",
							),
					)
				}

				float, err := strconv.ParseFloat(parts[1], 64)

				if err != nil {
					panic(
						NewUsageError(InvalidFloatMap, nil).
							WithFlag(flag).
							WithMessage(
								"Invalid float map value detected.",
							).
							WithExpected(
								"Expected float map values.",
							).
							WithReceived(
								fmt.Sprintf("Received value: %s", parts[1]),
							).
							WithSuggestion(
								"Provide valid float values such as 'ratio=1.5'.",
							),
					)
				}

				record[parts[0]] = float
			}

			return record

		default:
			panic(
				NewUsageError(InvalidFloatMap, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid float map detected.",
					).
					WithExpected(
						"Expected a map of float key-value pairs.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide valid float map entries in the format 'key=value'.",
					),
			)
		}
	})
}

func NewBoolMapFlag(name string) *Flag {
	return NewCustomListFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil, map[string]bool:
			return v

		case []string:
			record := make(map[string]bool)

			for i := range v {
				parts := strings.Split(v[i], "=")

				if len(parts) != 2 {
					panic(
						NewUsageError(InvalidBooleanMap, nil).
							WithFlag(flag).
							WithMessage(
								"Invalid boolean map entry detected.",
							).
							WithExpected(
								"Expected map entries in the format 'key=value'.",
							).
							WithReceived(
								fmt.Sprintf("Received entry: %s", v[i]),
							).
							WithSuggestion(
								"Provide valid boolean map entries such as 'enabled=true'.",
							),
					)
				}

				boolean, err := strconv.ParseBool(parts[1])

				if err != nil {
					panic(
						NewUsageError(InvalidBooleanMap, nil).
							WithFlag(flag).
							WithMessage(
								"Invalid boolean map value detected.",
							).
							WithExpected(
								"Expected boolean map values.",
							).
							WithReceived(
								fmt.Sprintf("Received value: %s", parts[1]),
							).
							WithSuggestion(
								"Provide valid boolean values such as 'enabled=true'.",
							),
					)
				}

				record[parts[0]] = boolean
			}

			return record

		default:
			panic(
				NewUsageError(InvalidBooleanMap, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid boolean map detected.",
					).
					WithExpected(
						"Expected a map of boolean key-value pairs.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide valid boolean map entries in the format 'key=value'.",
					),
			)
		}
	})
}

func NewComplexMapFlag(name string) *Flag {
	return NewCustomListFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil, map[string]any:
			return v

		case []string:
			record := make(map[string]any)

			for i := range v {
				parts := strings.Split(v[i], "=")

				if len(parts) != 2 {
					panic(
						NewUsageError(InvalidComplexMap, nil).
							WithFlag(flag).
							WithMessage(
								"Invalid complex map entry detected.",
							).
							WithExpected(
								"Expected map entries in the format 'key=value'.",
							).
							WithReceived(
								fmt.Sprintf("Received entry: %s", v[i]),
							).
							WithSuggestion(
								"Provide valid map entries such as 'name=john' or 'enabled=true'.",
							),
					)
				}

				key := parts[0]
				val := parts[1]

				if integer, err := strconv.Atoi(val); err == nil {
					record[key] = integer
					continue
				}

				if float, err := strconv.ParseFloat(val, 64); err == nil {
					record[key] = float
					continue
				}

				if boolean, err := strconv.ParseBool(val); slices.Contains([]string{"true", "false", "TRUE", "FALSE"}, val) && err == nil {
					record[key] = boolean
					continue
				}

				record[key] = val
			}

			return record

		default:
			panic(
				NewUsageError(InvalidComplexMap, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid complex map detected.",
					).
					WithExpected(
						"Expected a map of strings, integers, floats, or booleans.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide valid map entries in the format 'key=value'.",
					),
			)
		}
	})
}
