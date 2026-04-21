package void

import (
	"fmt"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

var pattern = regexp.MustCompile(`^.*/cmd/[^/]+/commands(/.*)?\.go$`)

func resolveFilePath() (string, error) {
	workPath, err := filepath.Abs(".")

	if err != nil {
		return "", CommandNotFound
	}

	_, filePath, _, ok := runtime.Caller(3)

	if !ok {
		return "", CommandNotFound
	}

	valid := pattern.MatchString(filePath)

	if !valid {
		return "", CommandNotFound
	}

	filePath = path.Clean(filePath)

	path, err := filepath.Rel(workPath, filePath)

	if err != nil {
		return "", CommandNotFound
	}

	parts := strings.Split(path, "\\commands\\")[1:]

	path = strings.Join(parts, fmt.Sprintf("%v", filepath.Separator))

	return path, nil
}

func resolveCommandPath() ([]string, string) {
	filePath, err := resolveFilePath()

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
