package flag

import (
	"fmt"
	"strconv"
	"strings"
)

func NewStringMapFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, map[string]string:
			return v

		case []string:
			record := make(map[string]string)

			for i := range v {
				parts := strings.Split(v[i], "=")

				if len(parts) != 2 {
					panic(fmt.Sprintf("Flag '%s' Must Be A Map Of Strings", flag.name))
				}

				record[parts[0]] = parts[1]
			}

			return record

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A Map Of Strings", flag.name))
		}
	})

	flag.variadic = true

	return flag
}

func NewIntMapFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, map[string]int:
			return v

		case []string:
			record := make(map[string]int)

			for i := range v {
				parts := strings.Split(v[i], "=")

				if len(parts) != 2 {
					panic(fmt.Sprintf("Flag '%s' Must Be A Map Of Integers", flag.name))
				}

				integer, err := strconv.Atoi(parts[1])

				if err != nil {
					panic(fmt.Sprintf("Flag '%s' Must Be A Map Of Integers", flag.name))
				}

				record[parts[0]] = integer
			}

			return record

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A Map Of Integers", flag.name))
		}
	})

	flag.variadic = true

	return flag
}

func NewFloatMapFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, map[string]float64:
			return v

		case []string:
			record := make(map[string]float64)

			for i := range v {
				parts := strings.Split(v[i], "=")

				if len(parts) != 2 {
					panic(fmt.Sprintf("Flag '%s' Must Be A Map Of Floats", flag.name))
				}

				float, err := strconv.ParseFloat(parts[1], 64)

				if err != nil {
					panic(fmt.Sprintf("Flag '%s' Must Be A Map Of Floats", flag.name))
				}

				record[parts[0]] = float
			}

			return record

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A Map Of Floats", flag.name))
		}
	})

	flag.variadic = true

	return flag
}

func NewBoolMapFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, map[string]bool:
			return v

		case []string:
			record := make(map[string]bool)

			for i := range v {
				parts := strings.Split(v[i], "=")

				if len(parts) != 2 {
					panic(fmt.Sprintf("Flag '%s' Must Be A Map Of Booleans", flag.name))
				}

				boolean, err := strconv.ParseBool(strings.ToLower(parts[1]))

				if err != nil {
					panic(fmt.Sprintf("Flag '%s' Must Be A Map Of Booleans", flag.name))
				}

				record[parts[0]] = boolean
			}

			return record

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A Map Of Booleans", flag.name))
		}
	})

	flag.variadic = true

	return flag
}

func NewComplexMapFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, map[string]any:
			return v

		case []string:
			record := make(map[string]any)

			for i := range v {
				parts := strings.Split(v[i], "=")

				if len(parts) != 2 {
					panic(fmt.Sprintf("Flag '%s' Must Be A Map Of Strings, Integers, Floats, Or Booleans", flag.name))
				}

				key := parts[0]
				val := parts[1]

				if boolean, err := strconv.ParseBool(val); err == nil {
					record[key] = boolean
					continue
				}

				if len(strings.Split(val, ".")) == 1 {
					integer, err := strconv.Atoi(val)

					if err != nil {
						record[key] = val
						continue
					}

					record[key] = integer
					continue
				}

				if len(strings.Split(val, ".")) == 2 {
					float, err := strconv.ParseFloat(val, 64)

					if err != nil {
						record[key] = val
						continue
					}

					record[key] = float
					continue
				}

				record[key] = val
			}

			return record

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A Map Of Strings, Integers, Floats, Or Booleans", flag.name))
		}
	})

	flag.variadic = true

	return flag
}
