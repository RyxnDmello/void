package void

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

const (
	DirectoryNotFound          ErrorCode = "DIRECTORY_NOT_FOUND"
	InvalidDirectoryPath       ErrorCode = "INVALID_DIRECTORY_PATH"
	FailedToCreateDirectory    ErrorCode = "FAILED_TO_CREATE_DIRECTORY"
	InvalidDirectoryPathList   ErrorCode = "INVALID_DIRECTORY_PATH_LIST"
	InvalidDirectoryPathOption ErrorCode = "INVALID_DIRECTORY_PATH_OPTION"
)

func NewDirFlag(name string, create bool) *Flag {
	return NewCustomFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil:
			return v

		case string:
			return mustParseDir(v, create)

		default:
			panic(
				NewUsageError(InvalidDirectoryPath, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid directory path detected.",
					).
					WithExpected(
						"Expected a valid directory path.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid directory path.",
					),
			)
		}
	})
}

func NewDirListFlag(name string, create bool) *Flag {
	return NewCustomListFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil:
			return v

		case []string:
			paths := make([]string, 0)

			for i := range v {
				paths = append(paths, mustParseDir(v[i], create))
			}

			return paths

		default:
			panic(
				NewUsageError(InvalidDirectoryPathList, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid directory path list detected.",
					).
					WithExpected(
						"Expected a list of valid directory paths.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid list of directory paths.",
					),
			)
		}
	})
}

func NewDirEnumFlag(name string, create bool, enum []string) *Flag {
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
			if path := mustParseDir(v, false); slices.Contains(options, path) {
				return path
			}

			panic(
				NewUsageError(InvalidDirectoryPathOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported directory path detected.",
					).
					WithExpected(
						fmt.Sprintf("Expected one of: %s", strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %s", v),
					).
					WithSuggestion(
						"Provide one of the supported enum directory paths.",
					),
			)

		default:
			panic(
				NewUsageError(InvalidDirectoryPathOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported directory path detected.",
					).
					WithExpected(
						fmt.Sprintf("Flag '%s' must receive one of: %s", flag.Name(), strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide one of the supported enum directory paths.",
					),
			)
		}
	})

	flag.enum = make([]any, 0)

	for i := range enum {
		flag.enum = append(flag.enum, mustParseDir(enum[i], create))
	}

	return flag
}

func mustParseDir(value string, create bool) string {
	path, err := filepath.Abs(value)

	if err != nil {
		NewUsageError(InvalidDirectoryPath, err).
			WithMessage(
				"Invalid directory path detected.",
			).
			WithExpected(
				"Expected a valid directory path.",
			).
			WithReceived(
				fmt.Sprintf("Received path: %s", value),
			).
			WithSuggestion(
				"Provide a valid directory path.",
			).
			Fatal()
	}

	if create {
		if err := os.MkdirAll(path, os.ModePerm); err == nil {
			return path
		}

		NewRuntimeError(FailedToCreateDirectory, err).
			WithMessage(
				"Failed to create directory.",
			).
			WithExpected(
				"The directory should be created successfully.",
			).
			WithReceived(
				fmt.Sprintf("Failed to create directory: %s", path),
			).
			WithSuggestion(
				"Ensure the process has permission to create directories at the target location.",
			).
			Fatal()
	}

	if info, err := os.Stat(path); err == nil {
		if info.IsDir() {
			return path
		}

		NewUsageError(InvalidDirectoryPath, nil).
			WithMessage(
				"Path is not a directory.",
			).
			WithExpected(
				"Expected the provided path to resolve to a directory.",
			).
			WithReceived(
				fmt.Sprintf("Received file path: %s", path),
			).
			WithSuggestion(
				"Provide a valid directory path instead of a file path.",
			).
			Fatal()
	}

	NewUsageError(DirectoryNotFound, nil).
		WithMessage(
			"Directory could not be found.",
		).
		WithExpected(
			"Expected an existing directory path.",
		).
		WithReceived(
			fmt.Sprintf("Directory path not found: %s", path),
		).
		WithSuggestion(
			"Provide a valid existing directory path or enable directory creation.",
		).
		Fatal()

	return ""
}
