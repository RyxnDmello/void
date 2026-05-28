package void

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

const (
	InvalidSemver       ErrorCode = "INVALID_SEMVER"
	InvalidSemverTag    ErrorCode = "INVALID_SEMVER_TAG"
	InvalidSemverList   ErrorCode = "INVALID_SEMVER_LIST"
	InvalidSemverOption ErrorCode = "INVALID_SEMVER_OPTION"
)

var (
	versionRegexp = regexp.MustCompile(`^v?([0-9]+)\.([0-9]+)\.([0-9]+)$`)
	releaseRegexp = regexp.MustCompile(`^[0-9A-Za-z]+(?:[.-][0-9A-Za-z]+)*$`)
)

func NewSemverFlag(name string) *Flag {
	return NewCustomFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil:
			return v

		case string:
			return mustParseSemver(v)

		default:
			panic(
				NewUsageError(InvalidSemver, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid semantic version detected.",
					).
					WithExpected(
						"Expected a valid semantic version.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid semantic version such as '1.0.0' or 'v1.2.3-beta'.",
					),
			)
		}
	})
}

func NewSemverListFlag(name string) *Flag {
	return NewCustomListFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil:
			return v

		case []string:
			versions := make([]string, 0)

			for i := range v {
				versions = append(versions, mustParseSemver(v[i]))
			}

			return versions

		default:
			panic(
				NewUsageError(InvalidSemverList, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid semantic version list detected.",
					).
					WithExpected(
						"Expected a list of valid semantic versions.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide valid semantic versions such as '1.0.0' or 'v1.2.3-beta'.",
					),
			)
		}
	})
}

func NewSemverEnumFlag(name string, enum []string) *Flag {
	flag := NewCustomEnumFlag(name, enum, func(flag flagView, value any) any {
		if value == nil {
			return nil
		}

		_, enum := flag.Enumerable()

		options := make([]string, 0)

		for i := range enum {
			options = append(options, enum[i].(string))
		}

		switch v := value.(type) {
		case string:
			version := mustParseSemver(v)

			if slices.Contains(options, version) {
				return version
			}

			panic(
				NewUsageError(InvalidSemverOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported semantic version detected.",
					).
					WithExpected(
						fmt.Sprintf("Expected one of: %s", strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %s", version),
					).
					WithSuggestion(
						"Provide one of the supported enum semantic versions.",
					),
			)

		default:
			panic(
				NewUsageError(InvalidSemverOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported semantic version detected.",
					).
					WithExpected(
						fmt.Sprintf("Flag '%s' must receive one of: %s", flag.Name(), strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide one of the supported enum semantic versions.",
					),
			)
		}
	})

	flag.enum = make([]any, 0)

	for i := range enum {
		flag.enum = append(flag.enum, mustParseSemver(enum[i]))
	}

	return flag
}

func mustParseSemver(value string) string {
	version := strings.ToLower(strings.TrimSpace(value))

	sections := strings.SplitN(version, "-", 2)

	semver := sections[0]

	if !versionRegexp.MatchString(semver) {
		NewUsageError(InvalidSemver, nil).
			WithMessage(
				"Invalid semantic version detected.",
			).
			WithExpected(
				"Expected a valid semantic version in the format 'MAJOR.MINOR.PATCH'.",
			).
			WithReceived(
				fmt.Sprintf("Received version: %s", semver),
			).
			WithSuggestion(
				"Provide a valid semantic version such as '1.0.0' or 'v1.2.3'.",
			).
			Fatal()
	}

	if len(sections) > 1 {
		release := sections[1]

		if !releaseRegexp.MatchString(release) {
			NewUsageError(InvalidSemverTag, nil).
				WithMessage(
					"Invalid semantic version release tag detected.",
				).
				WithExpected(
					"Expected a valid semantic version release tag.",
				).
				WithReceived(
					fmt.Sprintf("Received release tag: %s", release),
				).
				WithSuggestion(
					"Use release tags containing only alphanumeric characters, dots, or hyphens.",
				).
				Fatal()
		}
	}

	return version
}
