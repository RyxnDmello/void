package flag

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

var versionRegexp = regexp.MustCompile(`^v?([0-9]+)\.([0-9]+)\.([0-9]+)$`)
var releaseRegexp = regexp.MustCompile(`^[0-9A-Za-z]+(?:[.-][0-9A-Za-z]+)*$`)

func NewSemverFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil:
			return v

		case string:
			return mustParseSemver(v)

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A Version", flag.name))
		}
	})

	return flag
}

func NewSemverListFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
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
			panic(fmt.Sprintf("Flag '%s' Must Be A List Of Versions", flag.name))
		}
	})

	flag.variadic = true

	return flag
}

func NewSemverEnumFlag(name string, enum []string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	if len(enum) == 0 {
		panic(fmt.Sprintf("Flag '%s' Must Have At Least 1 Option", flag.name))
	}

	for i := range enum {
		flag.enum = append(flag.enum, mustParseSemver(enum[i]))
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
			version := mustParseSemver(v)

			if slices.Contains(options, version) {
				return version
			}

			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))
		}
	})

	return flag
}

func mustParseSemver(value string) string {
	version := strings.ToLower(strings.TrimSpace(value))

	sections := strings.SplitN(version, "-", 2)

	semver := sections[0]

	if !versionRegexp.MatchString(semver) {
		panic(fmt.Sprintf("'%s' Not A Valid Semver Version", semver))
	}

	if len(sections) > 1 {
		release := sections[1]

		if !releaseRegexp.MatchString(release) {
			panic(fmt.Sprintf("'-%s' Not A Valid Release Tag In '%s'", release, semver))
		}
	}

	return version
}
