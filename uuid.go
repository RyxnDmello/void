package void

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

const (
	InvalidUUID       ErrorCode = "INVALID_UUID"
	InvalidUUIDList   ErrorCode = "INVALID_UUID_LIST"
	InvalidUUIDOption ErrorCode = "INVALID_UUID_OPTION"
)

var uuidRegexp = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[1-5][0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`)

func NewUUIDFlag(name string) *Flag {
	return NewCustomFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil:
			return v

		case string:
			if uuidRegexp.MatchString(v) {
				return v
			}

			panic(
				NewUsageError(InvalidUUID, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid UUID detected.",
					).
					WithExpected(
						"Expected a valid UUID.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %s", v),
					).
					WithSuggestion(
						"Provide a valid UUID such as '550e8400-e29b-41d4-a716-446655440000'.",
					),
			)

		default:
			panic(
				NewUsageError(InvalidUUID, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid UUID detected.",
					).
					WithExpected(
						"Expected a valid UUID.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid UUID value.",
					),
			)
		}
	})
}

func NewUUIDListFlag(name string) *Flag {
	return NewCustomListFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil:
			return v

		case []string:
			uuids := make([]string, 0)

			for i := range v {
				if !uuidRegexp.MatchString(v[i]) {
					panic(
						NewUsageError(InvalidUUIDList, nil).
							WithFlag(flag).
							WithMessage(
								"Invalid UUID detected.",
							).
							WithExpected(
								"Expected a valid UUID within the UUID list.",
							).
							WithReceived(
								fmt.Sprintf("Invalid UUID: %s", v[i]),
							).
							WithSuggestion(
								"Provide valid UUID values.",
							),
					)
				}

				uuids = append(uuids, v[i])
			}

			return uuids

		default:
			panic(
				NewUsageError(InvalidUUIDList, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid UUID list detected.",
					).
					WithExpected(
						"Expected a list of valid UUID values.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid list of UUID values.",
					),
			)
		}
	})
}

func NewUUIDEnumFlag(name string, enum []string) *Flag {
	flag := NewCustomEnumFlag(name, enum, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil:
			return v

		case string:
			if slices.Contains(enum, v) {
				return v
			}

			panic(
				NewUsageError(InvalidUUIDOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported UUID detected.",
					).
					WithExpected(
						fmt.Sprintf("Expected one of: %s", strings.Join(enum, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %s", v),
					).
					WithSuggestion(
						"Provide one of the supported enum UUID values.",
					),
			)

		default:
			panic(
				NewUsageError(InvalidUUIDOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported UUID detected.",
					).
					WithExpected(
						fmt.Sprintf("Flag '%s' must receive one of: %s", flag.Name(), strings.Join(enum, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide one of the supported enum UUID values.",
					),
			)
		}
	})

	for i := range enum {
		if !uuidRegexp.MatchString(enum[i]) {
			NewDefinitionError(InvalidUUIDOption, nil).
				WithFlag(flag).
				WithMessage(
					"Invalid UUID enum value detected.",
				).
				WithExpected(
					"All enum options must be valid UUID values.",
				).
				WithReceived(
					fmt.Sprintf("Invalid enum value: %s", enum[i]),
				).
				WithSuggestion(
					"Provide valid UUID values such as '550e8400-e29b-41d4-a716-446655440000'.",
				).
				Fatal()
		}
	}

	return flag
}
