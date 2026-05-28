package void

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

const (
	FileNotFound          ErrorCode = "FILE_NOT_FOUND"
	InvalidFilePath       ErrorCode = "INVALID_FILE_PATH"
	FailedToCreateFile    ErrorCode = "FAILED_TO_CREATE_FILE"
	InvalidFilePathList   ErrorCode = "INVALID_FILE_PATH_LIST"
	InvalidFilePathOption ErrorCode = "INVALID_FILE_PATH_OPTION"
)

func NewFileFlag(name string, create, force bool) *Flag {
	return NewCustomFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil:
			return v

		case string:
			return mustParseFile(v, create, force)

		default:
			panic(
				NewUsageError(InvalidFilePath, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid file path detected.",
					).
					WithExpected(
						"Expected a valid file path.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid file path.",
					),
			)
		}
	})
}

func NewFileListFlag(name string, create, force bool) *Flag {
	return NewCustomListFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil:
			return v

		case []string:
			paths := make([]string, 0)

			for i := range v {
				paths = append(paths, mustParseFile(v[i], create, force))
			}

			return paths

		default:
			panic(
				NewUsageError(InvalidFilePathList, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid file path list detected.",
					).
					WithExpected(
						"Expected a list of valid file paths.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid list of file paths.",
					),
			)
		}
	})
}

func NewFileEnumFlag(name string, create, force bool, enum []string) *Flag {
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
			if path := mustParseFile(v, false, false); slices.Contains(options, path) {
				return path
			}

			panic(
				NewUsageError(InvalidFilePathOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported file path detected.",
					).
					WithExpected(
						fmt.Sprintf("Expected one of: %s", strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %s", v),
					).
					WithSuggestion(
						"Provide one of the supported enum file paths.",
					),
			)

		default:
			panic(
				NewUsageError(InvalidFilePathOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported file path detected.",
					).
					WithExpected(
						fmt.Sprintf("Flag '%s' must receive one of: %s", flag.Name(), strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide one of the supported enum file paths.",
					),
			)
		}
	})

	flag.enum = make([]any, 0)

	for i := range enum {
		flag.enum = append(flag.enum, mustParseFile(enum[i], create, force))
	}

	return flag
}

func mustParseFile(value string, create, force bool) string {
	path, err := filepath.Abs(value)

	if err != nil {
		NewUsageError(InvalidFilePath, err).
			WithMessage(
				"Invalid file path detected.",
			).
			WithExpected(
				"Expected a valid file path.",
			).
			WithReceived(
				fmt.Sprintf("Received path: %s", value),
			).
			WithSuggestion(
				"Provide a valid file path.",
			).
			Fatal()
	}

	if create && force {
		return mustCreateFile(path)
	}

	info, err := os.Stat(path)

	if create && !force {
		if err == nil {
			if !info.IsDir() {
				return path
			}

			NewUsageError(InvalidFilePath, nil).
				WithMessage(
					"Path is a directory.",
				).
				WithExpected(
					"Expected the provided path to resolve to a file.",
				).
				WithReceived(
					fmt.Sprintf("Received directory path: %s", path),
				).
				WithSuggestion(
					"Provide a valid file path instead of a directory path.",
				).
				Fatal()
		}

		return mustCreateFile(path)
	}

	if err == nil {
		if !info.IsDir() {
			return path
		}

		NewUsageError(InvalidFilePath, nil).
			WithMessage(
				"Path is a directory.",
			).
			WithExpected(
				"Expected the provided path to resolve to a file.",
			).
			WithReceived(
				fmt.Sprintf("Received directory path: %s", path),
			).
			WithSuggestion(
				"Provide a valid file path instead of a directory path.",
			).
			Fatal()
	}

	NewUsageError(FileNotFound, nil).
		WithMessage(
			"File could not be found.",
		).
		WithExpected(
			"Expected an existing file path.",
		).
		WithReceived(
			fmt.Sprintf("File path not found: %s", path),
		).
		WithSuggestion(
			"Provide a valid existing file path or enable file creation.",
		).
		Fatal()

	return ""
}

func mustCreateFile(path string) string {
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		NewRuntimeError(FailedToCreateFile, err).
			WithMessage(
				"Failed to create file directory structure.",
			).
			WithExpected(
				"The parent directory structure should be created successfully.",
			).
			WithReceived(
				fmt.Sprintf("Failed to create file at path '%s'.", path),
			).
			WithSuggestion(
				"Ensure the process has permission to create directories at the target location.",
			).
			Fatal()
	}

	file, err := os.Create(path)

	if err != nil {
		NewRuntimeError(FailedToCreateFile, err).
			WithMessage(
				"Failed to create file.",
			).
			WithExpected(
				"The file should be created successfully.",
			).
			WithReceived(
				fmt.Sprintf("Failed to create file at path '%s'.", path),
			).
			WithSuggestion(
				"Ensure the process has permission to create files at the target location.",
			).
			Fatal()
	}

	if err := file.Close(); err == nil {
		return path
	}

	NewRuntimeError(FailedToCreateFile, err).
		WithMessage(
			"Failed to finalize file creation.",
		).
		WithExpected(
			"The created file should close successfully.",
		).
		WithReceived(
			fmt.Sprintf("Failed to create file at path '%s'.", path),
		).
		WithSuggestion(
			"Ensure the file system is writable and accessible.",
		).
		Fatal()

	return ""
}
