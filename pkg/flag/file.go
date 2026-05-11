package flag

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func NewFileFlag(name string, create, force bool) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil:
			return v

		case string:
			return mustParseFile(v, create, force)

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A Path To A File", v))
		}
	})

	return flag
}

func NewFileListFlag(name string, create, force bool) *Flag {
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
				paths = append(paths, mustParseFile(v[i], create, force))
			}

			return paths

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A List Of File Paths", v))
		}
	})

	flag.variadic = true

	return flag
}

func NewFileEnumFlag(name string, create, force bool, enum []string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	if len(enum) == 0 {
		panic(fmt.Sprintf("Flag '%s' Must Have At Least 1 Option", flag.name))
	}

	for i := range enum {
		flag.enum = append(flag.enum, mustParseFile(enum[i], create, force))
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
			if path := mustParseFile(v, false, false); slices.Contains(options, path) {
				return path
			}

			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))
		}
	})

	return flag
}

func mustParseFile(value string, create, force bool) string {
	path, err := filepath.Abs(value)

	if err != nil {
		panic(fmt.Sprintf("'%s' Not A Valid Path", value))
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

			panic(fmt.Sprintf("Path '%s' Exists As A Directory", path))
		}

		return mustCreateFile(path)
	}

	if err == nil {
		if !info.IsDir() {
			return path
		}

		panic(fmt.Sprintf("Path '%s' Exists As A Directory", path))
	}

	panic(fmt.Sprintf("'%s' File Path Cannot Be Found", path))
}

func mustCreateFile(path string) string {
	failure := fmt.Sprintf("Failed To Create File At Path '%s'", path)

	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		panic(failure)
	}

	file, err := os.Create(path)

	if err != nil {
		panic(failure)
	}

	if err := file.Close(); err == nil {
		return path
	}

	panic(failure)
}
