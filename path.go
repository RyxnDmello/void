package void

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

const (
	InvalidPath       ErrorCode = "INVALID_PATH"
	PathNotFound      ErrorCode = "PATH_NOT_FOUND"
	InvalidPathList   ErrorCode = "INVALID_PATH_LIST"
	InvalidPathOption ErrorCode = "INVALID_PATH_OPTION"
)

func NewPathFlag(name string, validate bool) *Flag {
	return NewCustomFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil:
			return v

		case string:
			return mustParsePath(v, validate)

		default:
			panic(
				NewUsageError(InvalidPath, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid system path detected.",
					).
					WithExpected(
						"Expected a valid system path.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid system path.",
					),
			)
		}
	})
}

func NewPathListFlag(name string, validate bool) *Flag {
	return NewCustomListFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil:
			return v

		case []string:
			paths := make([]string, 0)

			for i := range v {
				paths = append(paths, mustParsePath(v[i], validate))
			}

			return paths

		default:
			panic(
				NewUsageError(InvalidPathList, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid system path list detected.",
					).
					WithExpected(
						"Expected a list of valid system paths.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid list of system paths.",
					),
			)
		}
	})
}

func NewPathEnumFlag(name string, validate bool, enum []string) *Flag {
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
			if path := mustParsePath(v, false); slices.Contains(options, path) {
				return path
			}

			panic(
				NewUsageError(InvalidPathOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported system path detected.",
					).
					WithExpected(
						fmt.Sprintf("Expected one of: %s", strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %s", v),
					).
					WithSuggestion(
						"Provide one of the supported enum system paths.",
					),
			)

		default:
			panic(
				NewUsageError(InvalidPathOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported system path detected.",
					).
					WithExpected(
						fmt.Sprintf("Flag '%s' must receive one of: %s", flag.Name(), strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide one of the supported enum system paths.",
					),
			)
		}
	})

	flag.enum = make([]any, 0)

	for i := range enum {
		flag.enum = append(flag.enum, mustParsePath(enum[i], validate))
	}

	return flag
}

func mustParsePath(value string, validate bool) string {
	path, err := filepath.Abs(value)

	if err != nil {
		NewUsageError(InvalidPath, err).
			WithMessage(
				"Invalid system path detected.",
			).
			WithExpected(
				"Expected a valid system path.",
			).
			WithReceived(
				fmt.Sprintf("Received path: %s", value),
			).
			WithSuggestion(
				"Provide a valid system path.",
			).
			Fatal()
	}

	if !validate {
		return path
	}

	if _, err := os.Stat(path); err == nil {
		return path
	}

	NewUsageError(PathNotFound, nil).
		WithMessage(
			"System path could not be found.",
		).
		WithExpected(
			"Expected an existing system path.",
		).
		WithReceived(
			fmt.Sprintf("Path not found: %s", path),
		).
		WithSuggestion(
			"Provide a valid existing system path.",
		).
		Fatal()

	return ""
}
