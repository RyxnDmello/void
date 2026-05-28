package void

import (
	"fmt"
	"slices"
	"strings"
	"time"
)

const (
	InvalidDuration       ErrorCode = "INVALID_DURATION"
	InvalidDurationList   ErrorCode = "INVALID_DURATION_LIST"
	InvalidDurationOption ErrorCode = "INVALID_DURATION_OPTION"
)

func NewDurationFlag(name string) *Flag {
	return NewCustomFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil, time.Duration:
			return v

		case string:
			if duration, err := time.ParseDuration(v); err == nil {
				return duration
			}

			panic(
				NewUsageError(InvalidDuration, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid duration value detected.",
					).
					WithExpected(
						"Expected a valid duration value.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %s", v),
					).
					WithSuggestion(
						"Use a valid duration format such as '1h', '30m', or '15s'.",
					),
			)

		default:
			panic(
				NewUsageError(InvalidDuration, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid duration value detected.",
					).
					WithExpected(
						"Expected a valid duration value.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Use a valid duration value.",
					),
			)
		}
	})
}

func NewDurationListFlag(name string) *Flag {
	return NewCustomListFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil, []time.Duration:
			return v

		case []string:
			durations := make([]time.Duration, 0)

			for i := range v {
				duration, err := time.ParseDuration(v[i])

				if err != nil {
					panic(
						NewUsageError(InvalidDurationList, nil).
							WithFlag(flag).
							WithMessage(
								"Invalid duration value detected.",
							).
							WithExpected(
								"Expected a valid duration value within the duration list.",
							).
							WithReceived(
								fmt.Sprintf("Invalid duration value: %s", v[i]),
							).
							WithSuggestion(
								"Use valid duration formats such as '1h', '30m', or '15s'.",
							),
					)
				}

				durations = append(durations, duration)
			}

			return durations

		default:
			panic(
				NewUsageError(InvalidDurationList, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid duration list detected.",
					).
					WithExpected(
						fmt.Sprintf("Flag '%s' must receive a list of valid duration values.", flag.Name()),
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid list of duration values.",
					),
			)
		}
	})
}

func NewDurationEnumFlag(name string, enum []string) *Flag {
	flag := NewCustomEnumFlag(name, enum, func(flag flagView, value any) any {
		if value == nil {
			return nil
		}

		_, enum := flag.Enumerable()

		options := make([]string, 0)

		for i := range enum {
			options = append(options, enum[i].(time.Duration).String())
		}

		switch v := value.(type) {
		case time.Duration:
			if slices.Contains(options, v.String()) {
				return v
			}

			panic(
				NewUsageError(InvalidDurationOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported duration value detected.",
					).
					WithExpected(
						fmt.Sprintf("Expected one of: %s", strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %s", v.String()),
					).
					WithSuggestion(
						"Provide one of the supported enum duration values.",
					),
			)

		case string:
			duration, err := time.ParseDuration(v)

			if err != nil {
				panic(
					NewUsageError(InvalidDuration, nil).
						WithFlag(flag).
						WithMessage(
							"Invalid duration value detected.",
						).
						WithExpected(
							"Expected a valid duration value.",
						).
						WithReceived(
							fmt.Sprintf("Received value: %s", v),
						).
						WithSuggestion(
							"Use a valid duration format such as '1h', '30m', or '15s'.",
						),
				)
			}

			if slices.Contains(options, duration.String()) {
				return duration
			}

			panic(
				NewUsageError(InvalidDurationOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported duration value detected.",
					).
					WithExpected(
						fmt.Sprintf("Expected one of: %s", strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %s", duration.String()),
					).
					WithSuggestion(
						"Provide one of the supported enum duration values.",
					),
			)

		default:
			panic(
				NewUsageError(InvalidDurationOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported duration value detected.",
					).
					WithExpected(
						fmt.Sprintf("Flag '%s' must receive one of: %s", flag.Name(), strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide one of the supported enum duration values.",
					),
			)
		}
	})

	flag.enum = make([]any, 0)

	for i := range enum {
		duration, err := time.ParseDuration(enum[i])

		if err != nil {
			NewDefinitionError(InvalidDurationOption, nil).
				WithFlag(flag).
				WithMessage(
					"Invalid duration enum value detected.",
				).
				WithExpected(
					"All enum options must be valid duration values.",
				).
				WithReceived(
					fmt.Sprintf("Invalid enum value: %s", enum[i]),
				).
				WithSuggestion(
					"Use valid duration formats such as '1h', '30m', or '15s'.",
				).
				Fatal()
		}

		flag.enum = append(flag.enum, duration)
	}

	return flag
}
