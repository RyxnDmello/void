package flag

import (
	"fmt"
	"slices"
	"strings"
	"time"
)

func NewDurationFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, time.Duration:
			return v

		case string:
			if duration, err := time.ParseDuration(v); err == nil {
				return duration
			}

			panic(fmt.Sprintf("Flag '%s' Must Be A Valid Duration", flag.name))

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A Valid Duration", flag.name))
		}
	})

	return flag
}

func NewDurationListFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, []time.Duration:
			return v

		case []string:
			durations := make([]time.Duration, 0)

			for i := range v {
				duration, err := time.ParseDuration(v[i])

				if err != nil {
					panic(fmt.Sprintf("Flag '%s' Must Be A Valid Duration", flag.name))
				}

				durations = append(durations, duration)
			}

			return durations

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A List Of Durations", flag.name))
		}
	})

	flag.variadic = true

	return flag
}

func NewDurationEnumFlag(name string, enum []string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	if len(enum) == 0 {
		panic(fmt.Sprintf("Flag '%s' Must Have At Least 1 Option", flag.name))
	}

	for i := range enum {
		duration, err := time.ParseDuration(enum[i])

		if err != nil {
			panic(fmt.Sprintf("Flag '%s' Has An Invalid Option: %s", flag.name, enum[i]))
		}

		flag.enum = append(flag.enum, duration)
	}

	flag.SetConvert(func(value any) any {
		if value == nil {
			return nil
		}

		options := make([]string, 0)

		for i := range flag.enum {
			options = append(options, flag.enum[i].(time.Duration).String())
		}

		switch v := value.(type) {
		case time.Duration:
			if slices.Contains(options, v.String()) {
				return v
			}

			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))

		case string:
			duration, err := time.ParseDuration(v)

			if err != nil {
				panic(fmt.Sprintf("%s Not A Valid Duration", v))
			}

			if !slices.Contains(options, duration.String()) {
				panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))
			}

			return duration

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))
		}
	})

	return flag
}
