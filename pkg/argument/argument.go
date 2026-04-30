package argument

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Argument struct {
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

func NewArgument(name string) *Argument {
	argument := NewCustomArgument(name, func(value any) any {
		return nil
	})

	argument.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, string:
			return v

		default:
			return fmt.Sprintf("%v", v)
		}
	})

	return argument
}

func NewIntArgument(name string) *Argument {
	argument := NewCustomArgument(name, func(value any) any {
		return nil
	})

	argument.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, int:
			return v

		case string:
			if result, err := strconv.Atoi(v); err == nil {
				return result
			}

			panic(fmt.Sprintf("Argument '%s' Must Be An Integer", argument.name))

		default:
			panic(fmt.Sprintf("Argument '%s' Must Be An Integer", argument.name))
		}
	})

	return argument
}

func NewFloatArgument(name string) *Argument {
	argument := NewCustomArgument(name, func(value any) any {
		return nil
	})

	argument.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, float64:
			return v

		case int, string:
			raw := fmt.Sprintf("%v", v)

			if result, err := strconv.ParseFloat(raw, 64); err == nil {
				return result
			}

			panic(fmt.Sprintf("Argument '%s' Must Be A Float", argument.name))

		default:
			panic(fmt.Sprintf("Argument '%s' Must Be A Float", argument.name))
		}
	})

	return argument
}

func NewBoolArgument(name string) *Argument {
	argument := NewCustomArgument(name, func(value any) any {
		return nil
	})

	argument.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, bool:
			return v

		case string:
			if v == argument.name {
				return true
			}

			if result, err := strconv.ParseBool(v); err == nil {
				return result
			}

			panic(fmt.Sprintf("Argument '%s' Must Be A Boolean", argument.name))

		default:
			panic(fmt.Sprintf("Argument '%s' Must Be A Boolean", argument.name))
		}
	})

	return argument
}

func NewEnumArgument(name string, enum []string) *Argument {
	argument := NewCustomArgument(name, func(value any) any {
		return nil
	})

	if len(enum) == 0 {
		panic(fmt.Sprintf("Argument '%s' Of Type Enum Must Have At Least 1 Option", argument.name))
	}

	for i := range enum {
		argument.enum = append(argument.enum, enum[i])
	}

	argument.SetConvert(func(value any) any {
		if value == nil {
			return nil
		}

		enum := make([]string, 0)

		for i := range argument.enum {
			enum = append(enum, argument.enum[i].(string))
		}

		switch v := value.(type) {
		case string:
			if !slices.Contains(enum, v) {
				panic(fmt.Sprintf("Argument '%s' Must Be One Of %s", argument.name, strings.Join(enum, ", ")))
			}

			return v

		default:
			panic(fmt.Sprintf("Argument '%s' Must Be One Of %s", argument.name, strings.Join(enum, ", ")))
		}
	})

	return argument
}

func NewIntEnumArgument(name string, enum []int) *Argument {
	argument := NewCustomArgument(name, func(value any) any {
		return nil
	})

	if len(enum) == 0 {
		panic(fmt.Sprintf("Argument '%s' Of Type Enum Must Have At Least 1 Option", argument.name))
	}

	for i := range enum {
		argument.enum = append(argument.enum, enum[i])
	}

	argument.SetConvert(func(value any) any {
		if value == nil {
			return nil
		}

		enum := make([]string, 0)

		for i := range argument.enum {
			enum = append(enum, fmt.Sprintf("%v", argument.enum[i]))
		}

		switch v := value.(type) {
		case int, string:
			raw := fmt.Sprintf("%v", v)

			result, err := strconv.Atoi(raw)

			if err != nil {
				panic(fmt.Sprintf("Argument '%s' Must Be An Integer", argument.name))
			}

			if !slices.Contains(enum, raw) {
				panic(fmt.Sprintf("Argument '%s' Must Be One Of %s", argument.name, strings.Join(enum, ", ")))
			}

			return result

		default:
			panic(fmt.Sprintf("Argument '%s' Must Be One Of %s", argument.name, strings.Join(enum, ", ")))
		}
	})

	return argument
}

func NewFloatEnumArgument(name string, enum []float64) *Argument {
	argument := NewCustomArgument(name, func(value any) any {
		return nil
	})

	if len(enum) == 0 {
		panic(fmt.Sprintf("Argument '%s' Of Type Enum Must Have At Least 1 Option", argument.name))
	}

	for i := range enum {
		argument.enum = append(argument.enum, enum[i])
	}

	argument.SetConvert(func(value any) any {
		if value == nil {
			return nil
		}

		enum := make([]string, 0)

		for i := range argument.enum {
			enum = append(enum, fmt.Sprintf("%v", argument.enum[i]))
		}

		switch v := value.(type) {
		case int, float64, string:
			raw := fmt.Sprintf("%v", v)

			result, err := strconv.ParseFloat(raw, 64)

			if err != nil {
				panic(fmt.Sprintf("Argument '%s' Must Be A Float", argument.name))
			}

			if !slices.Contains(enum, raw) {
				panic(fmt.Sprintf("Argument '%s' Must Be One Of %s", argument.name, strings.Join(enum, ", ")))
			}

			return result

		default:
			panic(fmt.Sprintf("Argument '%s' Must Be One Of %s", argument.name, strings.Join(enum, ", ")))
		}
	})

	return argument
}

func NewListArgument(name string) *Argument {
	argument := NewCustomArgument(name, func(value any) any {
		return nil
	})

	argument.variadic = true

	argument.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, []string:
			return v

		default:
			panic(fmt.Sprintf("Argument '%s' Must Be A List Of Strings", argument.name))
		}
	})

	return argument
}

func NewIntListArgument(name string) *Argument {
	argument := NewCustomArgument(name, func(value any) any {
		return nil
	})

	argument.variadic = true

	argument.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, []int:
			return v

		case []string:
			result := make([]int, 0)

			for i := range v {
				value, err := strconv.Atoi(v[i])

				if err != nil {
					panic(fmt.Sprintf("Argument '%s' Must Be A List Of Integers", argument.name))
				}

				result = append(result, value)
			}

			return result

		default:
			panic(fmt.Sprintf("Argument '%s' Must Be A List Of Integers", argument.name))
		}
	})

	return argument
}

func NewFloatListArgument(name string) *Argument {
	argument := NewCustomArgument(name, func(value any) any {
		return nil
	})

	argument.variadic = true

	argument.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, []float64:
			return v

		case []int:
			result := make([]float64, 0)

			for i := range v {
				raw := fmt.Sprintf("%v", v[i])

				value, err := strconv.ParseFloat(raw, 64)

				if err != nil {
					panic(fmt.Sprintf("Argument '%s' Must Be A List Of Floats", argument.name))
				}

				result = append(result, value)
			}

			return result

		case []string:
			result := make([]float64, 0)

			for i := range v {
				value, err := strconv.ParseFloat(v[i], 64)

				if err != nil {
					panic(fmt.Sprintf("Argument '%s' Must Be A List Of Floats", argument.name))
				}

				result = append(result, value)
			}

			return result

		default:
			panic(fmt.Sprintf("Argument '%s' Must Be A List Of Floats", argument.name))
		}
	})

	return argument
}

func NewBoolListArgument(name string) *Argument {
	argument := NewCustomArgument(name, func(value any) any {
		return nil
	})

	argument.variadic = true

	argument.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, []bool:
			return v

		case []string:
			result := make([]bool, 0)

			for i := range v {
				value, err := strconv.ParseBool(v[i])

				if err != nil {
					panic(fmt.Sprintf("Argument '%s' Must Be A List Of Booleans", argument.name))
				}

				result = append(result, value)
			}

			return result

		default:
			panic(fmt.Sprintf("Argument '%s' Must Be A List Of Booleans", argument.name))
		}
	})

	return argument
}

func NewCustomArgument(name string, convert func(value any) any) *Argument {
	name = mustCleanName(name)

	if convert == nil {
		panic(fmt.Sprintf("Argument '%s' Must Have A Type Conversion Handler", name))
	}

	return &Argument{
		name:        name,
		tagline:     "",
		description: "",
		required:    true,
		variadic:    false,
		value:       nil,
		enum:        nil,
		covert:      nil,
		validate:    nil,
	}
}

func (argument *Argument) Name() string {
	return argument.name
}

func (argument *Argument) Tagline() string {
	return argument.tagline
}

func (argument *Argument) SetTagline(tagline string) *Argument {
	argument.tagline = strings.Trim(tagline, " ")
	return argument
}

func (argument *Argument) Description() string {
	return argument.description
}

func (argument *Argument) SetDescription(description string) *Argument {
	argument.description = strings.Trim(description, " ")
	return argument
}

func (argument *Argument) IsRequired() bool {
	return argument.required
}

func (argument *Argument) Optional(value any) *Argument {
	argument.value = value
	argument.required = false
	return argument
}

func (argument *Argument) IsVariadic() bool {
	return argument.variadic
}

func (argument *Argument) Value() any {
	return argument.covert(argument.value)
}

func (argument *Argument) SetValue(value any) *Argument {
	if value == nil {
		panic(fmt.Sprintf("Argument '%s' Must Either Be Assigned A Value Or Set As Optional", argument.name))
	}

	argument.value = value
	return argument
}

func (argument *Argument) Enum() []any {
	return argument.enum
}

func (argument *Argument) SetConvert(convert func(value any) any) *Argument {
	if convert == nil {
		panic(fmt.Sprintf("Argument '%s' Must Have A Type Conversion Handler", argument.name))
	}

	argument.covert = func(value any) any {
		result := convert(value)

		if argument.validate != nil {
			return argument.validate(result)
		}

		return result
	}

	return argument
}

func (argument *Argument) SetValidate(validate func(value any) any) *Argument {
	argument.validate = validate
	return argument
}

func (argument *Argument) MustValidate() {
	if len(argument.name) == 0 {
		panic("Every Argument Must Have A Name")
	}

	if len(argument.tagline) == 0 {
		panic(fmt.Sprintf("Argument '%s' Must Have A Tagline", argument.name))
	}

	if len(argument.description) == 0 {
		panic(fmt.Sprintf("Argument '%s' Must Have A Tagline", argument.name))
	}

	if argument.covert == nil {
		panic(fmt.Sprintf("Argument '%s' Must Have A Type Conversion Handler", argument.name))
	}
}
