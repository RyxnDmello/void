package flag

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Flag struct {
	name        string
	tagline     string
	description string
	required    bool
	variadic    bool
	value       any
	enum        []any
	covert      func(value any) any
	validate    func(value any) any
}

func NewFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, string:
			return v

		default:
			return fmt.Sprintf("%v", v)
		}
	})

	return flag
}

func NewIntFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, int:
			return v

		case string:
			if result, err := strconv.Atoi(v); err == nil {
				return result
			}

			panic(fmt.Sprintf("Flag '%s' Must Be An Integer", flag.name))

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be An Integer", flag.name))
		}
	})

	return flag
}

func NewFloatFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, float64:
			return v

		case int, string:
			raw := fmt.Sprintf("%v", v)

			if result, err := strconv.ParseFloat(raw, 64); err == nil {
				return result
			}

			panic(fmt.Sprintf("Flag '%s' Must Be A Float", flag.name))

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A Float", flag.name))
		}
	})

	return flag
}

func NewBoolFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.value = false

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, bool:
			return v

		case string:
			if result, err := strconv.ParseBool(v); err == nil {
				return result
			}

			panic(fmt.Sprintf("Flag '%s' Must Be A Boolean", flag.name))

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A Boolean", flag.name))
		}
	})

	return flag
}

func NewEnumFlag(name string, enum []string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	if len(enum) == 0 {
		panic(fmt.Sprintf("Flag '%s' Of Type Enum Must Have At Least 1 Option", flag.name))
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
			enum = append(enum, flag.enum[i].(string))
		}

		switch v := value.(type) {
		case string:
			if !slices.Contains(enum, v) {
				panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(enum, ", ")))
			}

			return v

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(enum, ", ")))
		}
	})

	return flag
}

func NewIntEnumFlag(name string, enum []int) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	if len(enum) == 0 {
		panic(fmt.Sprintf("Flag '%s' Of Type Enum Must Have At Least 1 Option", flag.name))
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
			enum = append(enum, fmt.Sprintf("%v", flag.enum[i]))
		}

		switch v := value.(type) {
		case int, string:
			raw := fmt.Sprintf("%v", v)

			result, err := strconv.Atoi(raw)

			if err != nil {
				panic(fmt.Sprintf("Flag '%s' Must Be An Integer", flag.name))
			}

			if !slices.Contains(enum, raw) {
				panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(enum, ", ")))
			}

			return result

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(enum, ", ")))
		}
	})

	return flag
}

func NewFloatEnumFlag(name string, enum []float64) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	if len(enum) == 0 {
		panic(fmt.Sprintf("Flag '%s' Of Type Enum Must Have At Least 1 Option", flag.name))
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
			enum = append(enum, fmt.Sprintf("%v", flag.enum[i]))
		}

		switch v := value.(type) {
		case int, float64, string:
			raw := fmt.Sprintf("%v", v)

			result, err := strconv.ParseFloat(raw, 64)

			if err != nil {
				panic(fmt.Sprintf("Flag '%s' Must Be A Float", flag.name))
			}

			if !slices.Contains(enum, raw) {
				panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(enum, ", ")))
			}

			return result

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(enum, ", ")))
		}
	})

	return flag
}

func NewListFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.variadic = true

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, []string:
			return v

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A List Of Strings", flag.name))
		}
	})

	return flag
}

func NewIntListFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.variadic = true

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, []int:
			return v

		case []string:
			result := make([]int, 0)

			for i := range v {
				value, err := strconv.Atoi(v[i])

				if err != nil {
					panic(fmt.Sprintf("Flag '%s' Must Be A List Of Integers", flag.name))
				}

				result = append(result, value)
			}

			return result

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A List Of Integers", flag.name))
		}
	})

	return flag
}

func NewFloatListFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.variadic = true

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, []float64:
			return v

		case []int:
			result := make([]float64, 0)

			for i := range v {
				raw := fmt.Sprintf("%v", v[i])

				value, err := strconv.ParseFloat(raw, 64)

				if err != nil {
					panic(fmt.Sprintf("Flag '%s' Must Be A List Of Floats", flag.name))
				}

				result = append(result, value)
			}

			return result

		case []string:
			result := make([]float64, 0)

			for i := range v {
				value, err := strconv.ParseFloat(v[i], 64)

				if err != nil {
					panic(fmt.Sprintf("Flag '%s' Must Be A List Of Floats", flag.name))
				}

				result = append(result, value)
			}

			return result

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A List Of Floats", flag.name))
		}
	})

	return flag
}

func NewBoolListFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.variadic = true

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, []bool:
			return v

		case []string:
			result := make([]bool, 0)

			for i := range v {
				value, err := strconv.ParseBool(v[i])

				if err != nil {
					panic(fmt.Sprintf("Flag '%s' Must Be A List Of Booleans", flag.name))
				}

				result = append(result, value)
			}

			return result

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A List Of Booleans", flag.name))
		}
	})

	return flag
}

func NewCustomFlag(name string, convert func(value any) any) *Flag {
	name = mustCleanName(name)

	if convert == nil {
		panic(fmt.Sprintf("Flag '%s' Must Have A Type Conversion Handler", name))
	}

	return &Flag{
		name:        name,
		tagline:     "",
		description: "",
		required:    false,
		variadic:    false,
		value:       nil,
		enum:        nil,
		covert:      nil,
		validate:    nil,
	}
}

func (flag *Flag) Name() string {
	return flag.name
}

func (flag *Flag) Tagline() string {
	return flag.tagline
}

func (flag *Flag) SetTagline(tagline string) *Flag {
	flag.tagline = strings.Trim(tagline, " ")
	return flag
}

func (flag *Flag) Description() string {
	return flag.description
}

func (flag *Flag) SetDescription(description string) *Flag {
	flag.description = strings.Trim(description, " ")
	return flag
}

func (flag *Flag) IsRequired() bool {
	return flag.required
}

func (flag *Flag) Required() *Flag {
	flag.value = nil
	flag.required = true
	return flag
}

func (flag *Flag) Default(value any) *Flag {
	flag.value = value
	return flag
}

func (flag *Flag) IsVariadic() bool {
	return flag.variadic
}

func (flag *Flag) Value() any {
	return flag.covert(flag.value)
}

func (flag *Flag) SetValue(value any) *Flag {
	if value == nil {
		panic(fmt.Sprintf("Flag '%s' Must Either Be Assigned A Value Or Set As Optional", flag.name))
	}

	flag.value = value
	return flag
}

func (flag *Flag) Enum() []any {
	return flag.enum
}

func (flag *Flag) SetConvert(convert func(value any) any) *Flag {
	if convert == nil {
		panic(fmt.Sprintf("Flag '%s' Must Have A Type Conversion Handler", flag.name))
	}

	flag.covert = func(value any) any {
		result := convert(value)

		if flag.validate != nil {
			return flag.validate(result)
		}

		return result
	}

	return flag
}

func (flag *Flag) SetValidate(validate func(value any) any) *Flag {
	flag.validate = validate
	return flag
}

func (flag *Flag) MustValidate() {
	if len(flag.name) == 0 {
		panic("Every Flag Must Have A Name")
	}

	if len(flag.tagline) == 0 {
		panic(fmt.Sprintf("Flag '%s' Must Have A Tagline", flag.name))
	}

	if len(flag.description) == 0 {
		panic(fmt.Sprintf("Flag '%s' Must Have A Tagline", flag.name))
	}

	if flag.covert == nil {
		panic(fmt.Sprintf("Flag '%s' Must Have A Type Conversion Handler", flag.name))
	}
}
