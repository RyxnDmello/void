package flag

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func NewPathFlag(name string, validate bool) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil:
			return v

		case string:
			return mustGetPath(v, validate)

		default:
			return fmt.Sprintf("Flag '%s' Must Be A Path", v)
		}
	})

	return flag
}

func NewPathListFlag(name string, validate bool) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil:
			return v

		case []string:
			paths := make([]string, 0)

			for i := range v {
				paths = append(paths, mustGetPath(v[i], validate))
			}

			return paths

		default:
			return fmt.Sprintf("Flag '%s' Must Be A Path", v)
		}
	})

	flag.variadic = true

	return flag
}

func NewPathEnumFlag(name string, validate bool, enum []string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	if len(enum) == 0 {
		panic(fmt.Sprintf("Flag '%s' Must Have At Least 1 Option", flag.name))
	}

	for i := range enum {
		flag.enum = append(flag.enum, mustGetPath(enum[i], validate))
	}

	flag.SetConvert(func(value any) any {
		if value == nil {
			return nil
		}

		options := make([]string, 0)

		for i := range flag.enum {
			options = append(options, flag.enum[i].(string))
		}

		switch v := value.(type) {
		case string:
			if path := mustGetPath(v, false); slices.Contains(options, path) {
				return path
			}

			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))
		}
	})

	return flag
}

func mustGetPath(value string, validate bool) string {
	path, err := filepath.Abs(value)

	if err != nil {
		panic(fmt.Sprintf("'%s' Not A Valid Path", value))
	}

	if !validate {
		return path
	}

	if _, err := os.Stat(path); err == nil {
		return path
	}

	panic(fmt.Sprintf("Path '%s' Cannot Be Found", path))
}
