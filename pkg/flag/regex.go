package flag

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

func NewRegexFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil:
			return v

		case string:
			if exp, err := regexp.Compile(v); err == nil {
				return exp
			}

			panic(fmt.Sprintf("'%s' Not A Valid Regular Expression", v))

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A Regular Expression", flag.name))
		}
	})

	return flag
}

func NewRegexListFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, []*regexp.Regexp:
			return v

		case []regexp.Regexp:
			exps := make([]*regexp.Regexp, 0)

			for i := range v {
				exps = append(exps, &v[i])
			}

			return exps

		case []string:
			exps := make([]*regexp.Regexp, 0)

			for i := range v {
				exp, err := regexp.Compile(v[i])

				if err != nil {
					panic(fmt.Sprintf("'%s' Not A Valid Regular Expression", v[i]))
				}

				exps = append(exps, exp)
			}

			return exps

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A List Of Regular Expressions", flag.name))
		}
	})

	flag.variadic = true

	return flag
}

func NewRegexEnumFlag(name string, enum []string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	if len(enum) == 0 {
		panic(fmt.Sprintf("Flag '%s' Must Have At Least 1 Option", flag.name))
	}

	for i := range enum {
		exp, err := regexp.Compile(enum[i])

		if err != nil {
			panic(fmt.Sprintf("Flag '%s' Has An Invalid Option: '%s'", flag.name, enum[i]))
		}

		flag.enum = append(flag.enum, exp)
	}

	flag.SetConvert(func(value any) any {
		if value == nil {
			return nil
		}

		options := make([]string, 0)

		for i := range flag.enum {
			options = append(options, flag.enum[i].(*regexp.Regexp).String())
		}

		switch v := value.(type) {
		case *regexp.Regexp:
			if slices.Contains(options, v.String()) {
				return v
			}

			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))

		case regexp.Regexp:
			if slices.Contains(options, v.String()) {
				return &v
			}

			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))

		case string:
			exp, err := regexp.Compile(v)

			if err != nil {
				panic(fmt.Sprintf("'%s' Not A Valid Regular Expression", v))
			}

			if slices.Contains(enum, v) {
				return exp
			}

			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))
		}
	})

	return flag
}
