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
			duration, err := time.ParseDuration(v)

			if err != nil {
				panic(fmt.Sprintf("Flag '%s' Must Be A Valid Duration", flag.name))
			}

			return duration

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A Valid Duration", flag.name))
		}
	})

	return flag
}

func NewDurationEnumFlag(name string, enum []time.Duration) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	if len(enum) == 0 {
		panic(fmt.Sprintf("Flag '%s' Must Have At Least 1 Option", flag.name))
	}

	for i := range enum {
		flag.enum = append(flag.enum, enum[i])
	}

	flag.SetConvert(func(value any) any {
		if value == nil {
			return nil
		}

		enum := make([]string, 0)

		for i := range flag.enum {
			enum = append(enum, flag.enum[i].(time.Duration).String())
		}

		switch v := value.(type) {
		case string:
			duration, err := time.ParseDuration(v)

			if err != nil || !slices.Contains(enum, duration.String()) {
				panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(enum, ", ")))
			}

			return duration

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(enum, ", ")))
		}
	})

	return flag
}

func NewDurationListFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.variadic = true

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

	return flag
}
