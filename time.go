package void

import (
	"fmt"
	"slices"
	"strings"
	"time"
)

const (
	InvalidDateTime         ErrorCode = "INVALID_DATE_TIME"
	InvalidDateTimeList     ErrorCode = "INVALID_DATE_TIME_LIST"
	InvalidDateTimeOption   ErrorCode = "INVALID_DATE_TIME_OPTION"
	DateTimeLayoutsRequired ErrorCode = "DATE_TIME_LAYOUTS_REQUIRED"
)

var (
	timeLayouts = []string{
		"15:04",
		"15:04:05",
	}
	dateLayouts = []string{
		"2006-01-02",
		"2006/01/02",
	}
	dateTimeLayouts = []string{
		time.RFC3339,
		"2006-01-02 15:04:05",
		"2006-01-02 15:04",
		"2006/01/02 15:04:05",
		"2006/01/02 15:04",
	}
)

func NewTimeFlag(name string) *Flag {
	return NewCustomDateTimeFlag(name, timeLayouts)
}

func NewTimeListFlag(name string) *Flag {
	return NewCustomDateTimeListFlag(name, timeLayouts)
}

func NewTimeEnumFlag(name string, enum []string) *Flag {
	return NewCustomDateTimeEnumFlag(name, enum, timeLayouts)
}

func NewDateFlag(name string) *Flag {
	return NewCustomDateTimeFlag(name, dateLayouts)
}

func NewDateListFlag(name string) *Flag {
	return NewCustomDateTimeListFlag(name, dateLayouts)
}

func NewDateEnumFlag(name string, enum []string) *Flag {
	return NewCustomDateTimeEnumFlag(name, enum, dateLayouts)
}

func NewDateTimeFlag(name string) *Flag {
	return NewCustomDateTimeFlag(name, dateTimeLayouts)
}

func NewDateTimeListFlag(name string) *Flag {
	return NewCustomDateTimeListFlag(name, dateTimeLayouts)
}

func NewDateTimeEnumFlag(name string, enum []string) *Flag {
	return NewCustomDateTimeEnumFlag(name, enum, dateTimeLayouts)
}

func NewCustomDateTimeFlag(name string, layouts []string) *Flag {
	flag := NewCustomFlag(name, func(flag flagView, value any) any {
		if value == nil {
			return nil
		}

		return mustConvDateTime(flag, value, layouts)
	})

	if len(layouts) == 0 {
		NewDefinitionError(DateTimeLayoutsRequired, nil).
			WithFlag(flag).
			WithMessage(
				"Date-time layouts are required.",
			).
			WithExpected(
				fmt.Sprintf("Flag '%s' must define at least one supported date-time layout.", flag.name),
			).
			WithSuggestion(
				"Provide one or more valid date-time layouts.",
			).
			Fatal()
	}

	return flag
}

func NewCustomDateTimeListFlag(name string, layouts []string) *Flag {
	flag := NewCustomListFlag(name, func(flag flagView, value any) any {
		if value == nil {
			return nil
		}

		return mustConvDateTimeList(flag, value, layouts)
	})

	if len(layouts) == 0 {
		NewDefinitionError(DateTimeLayoutsRequired, nil).
			WithFlag(flag).
			WithMessage(
				"Date-time layouts are required.",
			).
			WithExpected(
				fmt.Sprintf("Flag '%s' must define at least one supported date-time layout.", flag.name),
			).
			WithSuggestion(
				"Provide one or more valid date-time layouts.",
			).
			Fatal()
	}

	return flag
}

func NewCustomDateTimeEnumFlag(name string, enum []string, layouts []string) *Flag {
	flag := NewCustomEnumFlag(name, enum, func(flag flagView, value any) any {
		if value == nil {
			return nil
		}

		return mustConvDateTimeEnum(flag, value, enum, layouts)
	})

	if len(layouts) == 0 {
		NewDefinitionError(DateTimeLayoutsRequired, nil).
			WithFlag(flag).
			WithMessage(
				"Date-time layouts are required.",
			).
			WithExpected(
				fmt.Sprintf("Flag '%s' must define at least one supported date-time layout.", flag.name),
			).
			WithSuggestion(
				"Provide one or more valid date-time layouts.",
			).
			Fatal()
	}

	flag.enum = make([]any, 0)

	for i := range enum {
		for j := range layouts {
			if time, err := time.Parse(layouts[j], enum[i]); err == nil {
				flag.enum = append(flag.enum, time)
				break
			}
		}

		if len(flag.enum) != i+1 {
			NewDefinitionError(InvalidDateTimeOption, nil).
				WithFlag(flag).
				WithMessage(
					"Invalid date-time enum value detected.",
				).
				WithExpected(
					"All enum options must match one of the supported date-time layouts.",
				).
				WithReceived(
					fmt.Sprintf("Invalid enum option: %s", enum[i]),
				).
				WithSuggestion(
					"Ensure all enum options use valid date-time formats.",
				).
				Fatal()
		}
	}

	return flag
}

func mustConvDateTime(flag flagView, value any, layouts []string) time.Time {
	switch v := value.(type) {
	case time.Time:
		return v

	case string:
		for i := range layouts {
			if time, err := time.Parse(layouts[i], v); err == nil {
				return time
			}
		}

		panic(
			NewUsageError(InvalidDateTime, nil).
				WithFlag(flag).
				WithMessage(
					"Invalid date-time value detected.",
				).
				WithExpected(
					"Expected a valid date-time value matching one of the supported layouts.",
				).
				WithReceived(
					fmt.Sprintf("Received value: %s", v),
				).
				WithSuggestion(
					"Provide a valid date-time value using a supported date-time format.",
				),
		)

	default:
		panic(
			NewUsageError(InvalidDateTime, nil).
				WithFlag(flag).
				WithMessage(
					"Invalid date-time value detected.",
				).
				WithExpected(
					"Expected a valid date-time value.",
				).
				WithReceived(
					fmt.Sprintf("Received value: %v", v),
				).
				WithSuggestion(
					"Provide a valid date-time value.",
				),
		)
	}
}

func mustConvDateTimeList(flag flagView, value any, layouts []string) []time.Time {
	switch v := value.(type) {
	case []time.Time:
		return v

	case []string:
		times := make([]time.Time, 0)

		for i := range v {
			for j := range layouts {
				if time, err := time.Parse(layouts[j], v[i]); err == nil {
					times = append(times, time)
					break
				}
			}

			if len(times) != i+1 {
				panic(
					NewUsageError(InvalidDateTimeList, nil).
						WithFlag(flag).
						WithMessage(
							"Invalid date-time list detected.",
						).
						WithExpected(
							"Expected a list of valid date-time values matching one of the supported layouts.",
						).
						WithReceived(
							fmt.Sprintf("Invalid date-time value: %s", v[i]),
						).
						WithSuggestion(
							"Provide valid date-time values using supported date-time formats.",
						),
				)
			}
		}

		return times

	default:
		panic(
			NewUsageError(InvalidDateTimeList, nil).
				WithFlag(flag).
				WithMessage(
					"Invalid date-time list detected.",
				).
				WithExpected(
					fmt.Sprintf("Flag '%s' must receive a list of valid date-time values.", flag.Name()),
				).
				WithReceived(
					fmt.Sprintf("Received value: %v", v),
				).
				WithSuggestion(
					"Provide a valid list of date-time values.",
				),
		)
	}
}

func mustConvDateTimeEnum(flag flagView, value any, enum []string, layouts []string) time.Time {
	options := make([]string, 0)

	for i := range enum {
		for j := range layouts {
			if time, err := time.Parse(layouts[j], enum[i]); err == nil {
				options = append(options, time.String())
				break
			}
		}

		if len(options) != i+1 {
			panic(
				NewUsageError(InvalidDateTimeOption, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid date-time enum value detected.",
					).
					WithExpected(
						"All enum options must match one of the supported date-time layouts.",
					).
					WithReceived(
						fmt.Sprintf("Invalid enum value: %s", enum[i]),
					).
					WithSuggestion(
						"Ensure all enum options use valid date-time formats.",
					),
			)
		}
	}

	switch v := value.(type) {
	case time.Time:
		if slices.Contains(options, v.String()) {
			return v
		}

		panic(
			NewUsageError(InvalidDateTimeOption, nil).
				WithFlag(flag).
				WithMessage(
					"Unsupported date-time value detected.",
				).
				WithExpected(
					fmt.Sprintf("Expected one of: %s", strings.Join(enum, ", ")),
				).
				WithReceived(
					fmt.Sprintf("Received value: %s", v.String()),
				).
				WithSuggestion(
					"Provide one of the supported enum date-time values.",
				),
		)

	case string:
		for i := range layouts {
			if time, err := time.Parse(layouts[i], v); err == nil {
				if slices.Contains(options, time.String()) {
					return time
				}

				panic(
					NewUsageError(InvalidDateTimeOption, nil).
						WithFlag(flag).
						WithMessage(
							"Unsupported date-time value detected.",
						).
						WithExpected(
							fmt.Sprintf("Expected one of: %s", strings.Join(enum, ", ")),
						).
						WithReceived(
							fmt.Sprintf("Received value: %s", v),
						).
						WithSuggestion(
							"Provide one of the supported enum date-time values.",
						),
				)
			}
		}

		panic(
			NewUsageError(InvalidDateTime, nil).
				WithFlag(flag).
				WithMessage(
					"Invalid date-time value detected.",
				).
				WithExpected(
					"Expected a valid date-time value matching one of the supported layouts.",
				).
				WithReceived(
					fmt.Sprintf("Received value: %s", v),
				).
				WithSuggestion(
					"Provide a valid date-time value using a supported date-time format.",
				),
		)

	default:
		panic(
			NewUsageError(InvalidDateTimeOption, nil).
				WithFlag(flag).
				WithMessage(
					"Unsupported date-time value detected.",
				).
				WithExpected(
					fmt.Sprintf("Flag '%s' must receive one of: %s", flag.Name(), strings.Join(enum, ", ")),
				).
				WithReceived(
					fmt.Sprintf("Received value: %v", v),
				).
				WithSuggestion(
					"Provide one of the supported enum date-time values.",
				),
		)
	}
}
