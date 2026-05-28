package void

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

const (
	InvalidRegex       ErrorCode = "INVALID_REGEX"
	InvalidRegexList   ErrorCode = "INVALID_REGEX_LIST"
	InvalidRegexOption ErrorCode = "INVALID_REGEX_OPTION"
)

func NewRegexFlag(name string) *Flag {
	return NewCustomFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil:
			return v

		case string:
			if exp, err := regexp.Compile(v); err == nil {
				return exp
			}

			panic(
				NewUsageError(InvalidRegex, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid regular expression detected.",
					).
					WithExpected(
						"Expected a valid regular expression.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %s", v),
					).
					WithSuggestion(
						"Provide a valid regular expression pattern.",
					),
			)

		default:
			panic(
				NewUsageError(InvalidRegex, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid regular expression detected.",
					).
					WithExpected(
						"Expected a valid regular expression.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid regular expression pattern.",
					),
			)
		}
	})
}

func NewRegexListFlag(name string) *Flag {
	return NewCustomListFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil, []*regexp.Regexp:
			return v

		case []regexp.Regexp:
			exps := make([]*regexp.Regexp, 0)

			for i := range v {
				exps = append(exps, &v[i])
			}

			return exps

		case []string:
			exps := make([]*regexp.Regexp, 0)

			for i := range v {
				exp, err := regexp.Compile(v[i])

				if err != nil {
					panic(
						NewUsageError(InvalidRegexList, nil).
							WithFlag(flag).
							WithMessage(
								"Invalid regular expression detected.",
							).
							WithExpected(
								"Expected a valid regular expression within the regular expression list.",
							).
							WithReceived(
								fmt.Sprintf("Invalid regular expression: %s", v[i]),
							).
							WithSuggestion(
								"Provide valid regular expression patterns.",
							),
					)
				}

				exps = append(exps, exp)
			}

			return exps

		default:
			panic(
				NewUsageError(InvalidRegexList, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid regular expression list detected.",
					).
					WithExpected(
						"Expected a list of valid regular expressions.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid list of regular expression patterns.",
					),
			)
		}
	})
}

func NewRegexEnumFlag(name string, enum []string) *Flag {
	flag := NewCustomEnumFlag(name, enum, func(flag flagView, value any) any {
		if value == nil {
			return nil
		}

		_, enum := flag.Enumerable()

		options := make([]string, 0)

		for i := range enum {
			options = append(options, enum[i].(*regexp.Regexp).String())
		}

		switch v := value.(type) {
		case *regexp.Regexp:
			if slices.Contains(options, v.String()) {
				return v
			}

			panic(
				NewUsageError(InvalidRegexOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported regular expression detected.",
					).
					WithExpected(
						fmt.Sprintf("Expected one of: %s", strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %s", v.String()),
					).
					WithSuggestion(
						"Provide one of the supported enum regular expressions.",
					),
			)

		case regexp.Regexp:
			if slices.Contains(options, v.String()) {
				return &v
			}

			panic(
				NewUsageError(InvalidRegexOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported regular expression detected.",
					).
					WithExpected(
						fmt.Sprintf("Expected one of: %s", strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %s", v.String()),
					).
					WithSuggestion(
						"Provide one of the supported enum regular expressions.",
					),
			)

		case string:
			exp, err := regexp.Compile(v)

			if err != nil {
				panic(
					NewUsageError(InvalidRegex, nil).
						WithFlag(flag).
						WithMessage(
							"Invalid regular expression detected.",
						).
						WithExpected(
							"Expected a valid regular expression.",
						).
						WithReceived(
							fmt.Sprintf("Received value: %s", v),
						).
						WithSuggestion(
							"Provide a valid regular expression pattern.",
						),
				)
			}

			if slices.Contains(options, v) {
				return exp
			}

			panic(
				NewUsageError(InvalidRegexOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported regular expression detected.",
					).
					WithExpected(
						fmt.Sprintf("Expected one of: %s", strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %s", v),
					).
					WithSuggestion(
						"Provide one of the supported enum regular expressions.",
					),
			)

		default:
			panic(
				NewUsageError(InvalidRegexOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported regular expression detected.",
					).
					WithExpected(
						fmt.Sprintf("Flag '%s' must receive one of: %s", flag.Name(), strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide one of the supported enum regular expressions.",
					),
			)
		}
	})

	flag.enum = make([]any, 0)

	for i := range enum {
		exp, err := regexp.Compile(enum[i])

		if err != nil {
			NewDefinitionError(InvalidRegexOption, nil).
				WithFlag(flag).
				WithMessage(
					"Invalid regular expression enum value detected.",
				).
				WithExpected(
					"All enum options must be valid regular expressions.",
				).
				WithReceived(
					fmt.Sprintf("Invalid enum value: %s", enum[i]),
				).
				WithSuggestion(
					"Provide valid regular expression patterns.",
				).
				Fatal()
		}

		flag.enum = append(flag.enum, exp)
	}

	return flag
}
