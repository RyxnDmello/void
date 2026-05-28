package void

import (
	"fmt"
	"net/url"
	"slices"
	"strings"
)

const (
	InvalidURL       ErrorCode = "INVALID_URL"
	InvalidURLList   ErrorCode = "INVALID_URL_LIST"
	InvalidURLOption ErrorCode = "INVALID_URL_OPTION"
)

func NewURLFlag(name string) *Flag {
	return NewCustomFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil, *url.URL:
			return v

		case url.URL:
			return &v

		case string:
			if url, err := url.Parse(v); err == nil {
				return url
			}

			panic(
				NewUsageError(InvalidURL, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid URL detected.",
					).
					WithExpected(
						"Expected a valid URL.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %s", v),
					).
					WithSuggestion(
						"Provide a valid URL such as 'https://example.com'.",
					),
			)

		default:
			panic(
				NewUsageError(InvalidURL, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid URL detected.",
					).
					WithExpected(
						"Expected a valid URL.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid URL.",
					),
			)
		}
	})
}

func NewURLListFlag(name string) *Flag {
	return NewCustomListFlag(name, func(flag flagView, value any) any {
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
					panic(
						NewUsageError(InvalidURLList, nil).
							WithFlag(flag).
							WithMessage(
								"Invalid URL detected.",
							).
							WithExpected(
								"Expected a valid URL within the URL list.",
							).
							WithReceived(
								fmt.Sprintf("Invalid URL: %s", v[i]),
							).
							WithSuggestion(
								"Provide valid URLs such as 'https://example.com'.",
							),
					)
				}

				urls = append(urls, url)
			}

			return urls

		default:
			panic(
				NewUsageError(InvalidURLList, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid URL list detected.",
					).
					WithExpected(
						"Expected a list of valid URLs.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid list of URLs.",
					),
			)
		}
	})
}

func NewURLEnumFlag(name string, enum []string) *Flag {
	flag := NewCustomEnumFlag(name, enum, func(flag flagView, value any) any {
		if value == nil {
			return nil
		}

		_, enum := flag.Enumerable()

		options := make([]string, 0)

		for i := range enum {
			options = append(options, enum[i].(*url.URL).String())
		}

		switch v := value.(type) {
		case *url.URL:
			if slices.Contains(options, v.String()) {
				return v
			}

			panic(
				NewUsageError(InvalidURLOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported URL detected.",
					).
					WithExpected(
						fmt.Sprintf("Expected one of: %s", strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %s", v.String()),
					).
					WithSuggestion(
						"Provide one of the supported enum URLs.",
					),
			)

		case url.URL:
			if slices.Contains(options, v.String()) {
				return &v
			}

			panic(
				NewUsageError(InvalidURLOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported URL detected.",
					).
					WithExpected(
						fmt.Sprintf("Expected one of: %s", strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %s", v.String()),
					).
					WithSuggestion(
						"Provide one of the supported enum URLs.",
					),
			)

		case string:
			url, err := url.Parse(v)

			if err != nil {
				panic(
					NewUsageError(InvalidURL, nil).
						WithFlag(flag).
						WithMessage(
							"Invalid URL detected.",
						).
						WithExpected(
							"Expected a valid URL.",
						).
						WithReceived(
							fmt.Sprintf("Received value: %s", v),
						).
						WithSuggestion(
							"Provide a valid URL such as 'https://example.com'.",
						),
				)
			}

			if slices.Contains(options, url.String()) {
				return url
			}

			panic(
				NewUsageError(InvalidURLOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported URL detected.",
					).
					WithExpected(
						fmt.Sprintf("Expected one of: %s", strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %s", url.String()),
					).
					WithSuggestion(
						"Provide one of the supported enum URLs.",
					),
			)

		default:
			panic(
				NewUsageError(InvalidURLOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported URL detected.",
					).
					WithExpected(
						fmt.Sprintf("Flag '%s' must receive one of: %s", flag.Name(), strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide one of the supported enum URLs.",
					),
			)
		}
	})

	flag.enum = make([]any, 0)

	for i := range enum {
		url, err := url.Parse(enum[i])

		if err != nil {
			NewDefinitionError(InvalidURLOption, nil).
				WithFlag(flag).
				WithMessage(
					"Invalid URL enum value detected.",
				).
				WithExpected(
					"All enum options must be valid URLs.",
				).
				WithReceived(
					fmt.Sprintf("Invalid enum value: %s", enum[i]),
				).
				WithSuggestion(
					"Provide valid URLs such as 'https://example.com'.",
				).
				Fatal()
		}

		flag.enum = append(flag.enum, url)
	}

	return flag
}
