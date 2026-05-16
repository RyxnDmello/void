package void

import (
	"fmt"
	"slices"
	"strings"
	"time"
)

var (
	timeLayouts = []string{
		"15:04",
		"15:04:05",
	}

	dateLayouts = []string{
		"2006-01-02",
		"2006/01/02",
	}

	dateTimeLayouts = []string{
		time.RFC3339,
		"2006-01-02 15:04:05",
		"2006-01-02 15:04",
		"2006/01/02 15:04:05",
		"2006/01/02 15:04",
	}
)

func NewTimeFlag(name string) *Flag {
	return NewCustomDateTimeFlag(name, timeLayouts)
}

func NewDateFlag(name string) *Flag {
	return NewCustomDateTimeFlag(name, dateLayouts)
}

func NewDateTimeFlag(name string) *Flag {
	return NewCustomDateTimeFlag(name, dateTimeLayouts)
}

func NewTimeListFlag(name string) *Flag {
	return NewCustomDateTimeListFlag(name, timeLayouts)
}

func NewDateListFlag(name string) *Flag {
	return NewCustomDateTimeListFlag(name, dateLayouts)
}

func NewDateTimeListFlag(name string) *Flag {
	return NewCustomDateTimeListFlag(name, dateTimeLayouts)
}

func NewTimeEnumFlag(name string, enum []string) *Flag {
	return NewCustomDateTimeEnumFlag(name, enum, timeLayouts)
}

func NewDateEnumFlag(name string, enum []string) *Flag {
	return NewCustomDateTimeEnumFlag(name, enum, dateLayouts)
}

func NewDateTimeEnumFlag(name string, enum []string) *Flag {
	return NewCustomDateTimeEnumFlag(name, enum, dateTimeLayouts)
}

func NewCustomDateTimeFlag(name string, layouts []string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	if len(layouts) == 0 {
		panic(fmt.Sprintf("Flag '%s' Must Have At Least 1 Date Time Layout", flag.name))
	}

	flag.HandleConvert(func(value any) any {
		if value == nil {
			return nil
		}

		return mustConvDateTime(value, layouts)
	})

	return flag
}

func NewCustomDateTimeListFlag(name string, layouts []string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	if len(layouts) == 0 {
		panic(fmt.Sprintf("Flag '%s' Must Have At Least 1 Date Time Layout", flag.name))
	}

	flag.HandleConvert(func(value any) any {
		if value == nil {
			return nil
		}

		return mustConvDateTimeList(flag.name, value, layouts)
	})

	flag.variadic = true

	return flag
}

func NewCustomDateTimeEnumFlag(name string, enum []string, layouts []string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	if len(enum) == 0 {
		panic(fmt.Sprintf("Flag '%s' Must Have At Least 1 Option", flag.name))
	}

	if len(layouts) == 0 {
		panic(fmt.Sprintf("Flag '%s' Must Have At Least 1 Date Time Layout", flag.name))
	}

	for i := range enum {
		for j := range layouts {
			if time, err := time.Parse(layouts[j], enum[i]); err == nil {
				flag.enum = append(flag.enum, time)
				break
			}
		}

		if len(flag.enum) != i+1 {
			panic(fmt.Sprintf("%s Not A Valid Option", enum[i]))
		}
	}

	flag.HandleConvert(func(value any) any {
		if value == nil {
			return nil
		}

		return mustConvDateTimeEnum(flag.name, value, enum, layouts)
	})

	return flag
}

func mustConvDateTime(value any, layouts []string) time.Time {
	switch v := value.(type) {
	case time.Time:
		return v

	case string:
		for i := range layouts {
			if time, err := time.Parse(layouts[i], v); err == nil {
				return time
			}
		}

		panic(fmt.Sprintf("'%s' Not A Valid Date Time", v))

	default:
		panic(fmt.Sprintf("'%v' Not A Valid Date Time", v))
	}
}

func mustConvDateTimeList(name string, value any, layouts []string) []time.Time {
	switch v := value.(type) {
	case []time.Time:
		return v

	case []string:
		times := make([]time.Time, 0)

		for i := range v {
			for j := range layouts {
				if time, err := time.Parse(layouts[j], v[i]); err == nil {
					times = append(times, time)
					break
				}
			}

			if len(times) != i+1 {
				panic(fmt.Sprintf("'%s' Not A Valid Date Time", v[i]))
			}
		}

		return times

	default:
		panic(fmt.Sprintf("Flag '%s' Must Be A List Of Date Times", name))
	}
}

func mustConvDateTimeEnum(name string, value any, enum []string, layouts []string) time.Time {
	options := make([]string, 0)

	for i := range enum {
		for j := range layouts {
			if time, err := time.Parse(layouts[j], enum[i]); err == nil {
				options = append(options, time.String())
				break
			}
		}

		if len(options) != i+1 {
			panic(fmt.Sprintf("Flag '%s' Has An Invalid Option: '%s'", name, enum[i]))
		}
	}

	switch v := value.(type) {
	case time.Time:
		if slices.Contains(options, v.String()) {
			return v
		}

		panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", name, strings.Join(enum, ", ")))

	case string:
		for i := range layouts {
			if time, err := time.Parse(layouts[i], v); err == nil {
				if slices.Contains(options, time.String()) {
					return time
				}

				panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", name, strings.Join(enum, ", ")))
			}
		}

		panic(fmt.Sprintf("'%s' Not A Valid Date Time", v))

	default:
		panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", name, strings.Join(enum, ", ")))
	}
}
