package flag

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func NewDirFlag(name string, create bool) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil:
			return v

		case string:
			return mustGetDir(v, create)

		default:
			return fmt.Sprintf("Flag '%s' Must Be A Path To A Directory", v)
		}
	})

	return flag
}

func NewDirListFlag(name string, create bool) *Flag {
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
				paths = append(paths, mustGetDir(v[i], create))
			}

			return paths

		default:
			return fmt.Sprintf("Flag '%s' Must Be A List Of Directory Paths", v)
		}
	})

	flag.variadic = true

	return flag
}

func NewDirEnumFlag(name string, create bool, enum []string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	if len(enum) == 0 {
		panic(fmt.Sprintf("Flag '%s' Must Have At Least 1 Option", flag.name))
	}

	for i := range enum {
		flag.enum = append(flag.enum, mustGetDir(enum[i], create))
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
			if path := mustGetDir(v, false); slices.Contains(options, path) {
				return path
			}

			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))
		}
	})

	return flag
}

func mustGetDir(value string, create bool) string {
	path, err := filepath.Abs(value)

	if err != nil {
		panic(fmt.Sprintf("'%s' Not A Valid Path", value))
	}

	if create {
		if err := os.MkdirAll(path, os.ModePerm); err == nil {
			return path
		}

		panic(fmt.Sprintf("Failed To Create Directory '%s'", path))
	}

	info, err := os.Stat(path)

	if err == nil {
		if info.IsDir() {
			return path
		}

		panic(fmt.Sprintf("Path '%s' Exists As A File", path))
	}

	panic(fmt.Sprintf("'%s' Directory Path Cannot Be Found", path))
}
