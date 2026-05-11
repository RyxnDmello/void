package flag

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

var uuidRegexp = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[1-5][0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`)

func NewUUIDFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil:
			return v

		case string:
			if uuidRegexp.MatchString(v) {
				return v
			}

			panic(fmt.Sprintf("'%s' Not A Valid UUID", v))

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A UUID", flag.name))
		}
	})

	return flag
}

func NewUUIDListFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil:
			return v

		case []string:
			uuids := make([]string, 0)

			for i := range v {
				if !uuidRegexp.MatchString(v[i]) {
					panic(fmt.Sprintf("'%s' Not A Valid UUID", v[i]))
				}

				uuids = append(uuids, v[i])
			}

			return uuids

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A List Of UUIDs", flag.name))
		}
	})

	flag.variadic = true

	return flag
}

func NewUUIDEnumFlag(name string, enum []string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	if len(enum) == 0 {
		panic(fmt.Sprintf("Flag '%s' Must Have At Least 1 Option", flag.name))
	}

	for i := range enum {
		if !uuidRegexp.MatchString(enum[i]) {
			panic(fmt.Sprintf("Flag '%s' Has An Invalid Option: '%s'", flag.name, enum[i]))
		}

		flag.enum = append(flag.enum, enum[i])
	}

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil:
			return v

		case string:
			if slices.Contains(enum, v) {
				return v
			}

			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(enum, ", ")))

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(enum, ", ")))
		}
	})

	return flag
}
