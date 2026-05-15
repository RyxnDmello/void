package command

import (
	"fmt"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/ryxndmello/void/pkg/errors"
)

func mustResolvePath() ([]string, string) {
	filePath, err := mustResolveFile()

	if err != nil {
		panic(err)
	}

	var root = strings.Split(filepath.Dir(filePath), " ")
	var base = strings.Split(filepath.Base(filePath), ".")[0]

	if len(root) == 0 || root[0] == "." {
		return []string{base}, base
	}

	if base == "root" {
		return root, root[len(root)-1]
	}

	root = append(root, base)

	return root, base
}

func mustResolveFile() (string, error) {
	workPath, err := filepath.Abs(".")

	if err != nil {
		return "", errors.CommandNotFound
	}

	_, filePath, _, ok := runtime.Caller(3)

	if !ok {
		return "", errors.CommandNotFound
	}

	pattern := regexp.MustCompile(`^.*/cmd/[^/]+/commands(/.*)?\.go$`)

	valid := pattern.MatchString(filePath)

	if !valid {
		return "", errors.CommandNotFound
	}

	filePath = filepath.Clean(filePath)

	path, err := filepath.Rel(workPath, filePath)

	if err != nil {
		return "", errors.CommandNotFound
	}

	parts := strings.Split(path, "\\commands\\")[1:]

	path = strings.Join(parts, fmt.Sprintf("%v", filepath.Separator))

	return path, nil
}
