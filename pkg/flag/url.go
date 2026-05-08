package flag

import (
	"fmt"
	"net/url"
	"slices"
	"strings"
)

func NewUrlFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, *url.URL:
			return v

		case url.URL:
			return &v

		case string:
			if url, err := url.Parse(v); err == nil {
				return url
			}

			panic(fmt.Sprintf("%s Not A Valid URL", v))

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A URL", flag.name))
		}
	})

	return flag
}

func NewUrlListFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, []*url.URL:
			return v

		case []url.URL:
			urls := make([]*url.URL, 0)

			for i := range v {
				urls = append(urls, &v[i])
			}

			return urls

		case []string:
			urls := make([]*url.URL, 0)

			for i := range v {
				url, err := url.Parse(v[i])

				if err != nil {
					panic(fmt.Sprintf("%s Not A Valid URL", v[i]))
				}

				urls = append(urls, url)
			}

			return urls

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A List Of URLs", flag.name))
		}
	})

	flag.variadic = true

	return flag
}

func NewUrlEnumFlag(name string, enum []string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	if len(enum) == 0 {
		panic(fmt.Sprintf("Flag '%s' Must Have At Least 1 Option", flag.name))
	}

	for i := range enum {
		url, err := url.Parse(enum[i])

		if err != nil {
			panic(fmt.Sprintf("Flag '%s' Has An Invalid Option: %s", flag.name, enum[i]))
		}

		flag.enum = append(flag.enum, url)
	}

	flag.SetConvert(func(value any) any {
		if value == nil {
			return nil
		}

		options := make([]string, 0)

		for i := range flag.enum {
			options = append(options, flag.enum[i].(*url.URL).String())
		}

		switch v := value.(type) {
		case *url.URL:
			if slices.Contains(options, v.String()) {
				return v
			}

			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))

		case url.URL:
			if slices.Contains(options, v.String()) {
				return v
			}

			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))

		case string:
			url, err := url.Parse(v)

			if err != nil {
				panic(fmt.Sprintf("%s Not A Valid URL", v))
			}

			if !slices.Contains(options, url.String()) {
				panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))
			}

			return url

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))
		}
	})

	return flag
}
