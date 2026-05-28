package void

import (
	"fmt"
	"net/mail"
	"slices"
	"strings"
)

const (
	InvalidEmailAddress       ErrorCode = "INVALID_EMAIL_ADDRESS"
	InvalidEmailAddressList   ErrorCode = "INVALID_EMAIL_ADDRESS_LIST"
	InvalidEmailAddressOption ErrorCode = "INVALID_EMAIL_ADDRESS_OPTION"
)

func NewEmailFlag(name string) *Flag {
	return NewCustomFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil, *mail.Address:
			return v

		case mail.Address:
			return &v

		case string:
			if address, err := mail.ParseAddress(v); err == nil {
				return address
			}

			panic(
				NewUsageError(InvalidEmailAddress, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid email address detected.",
					).
					WithExpected(
						"Expected a valid email address.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %s", v),
					).
					WithSuggestion(
						"Provide a valid email address such as 'user@example.com'.",
					),
			)

		default:
			panic(
				NewUsageError(InvalidEmailAddress, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid email address detected.",
					).
					WithExpected(
						"Expected a valid email address.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid email address.",
					),
			)
		}
	})
}

func NewEmailListFlag(name string) *Flag {
	return NewCustomListFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil, []*mail.Address:
			return v

		case []mail.Address:
			addresses := make([]*mail.Address, 0)

			for i := range v {
				addresses = append(addresses, &v[i])
			}

			return addresses

		case []string:
			addresses := make([]*mail.Address, 0)

			for i := range v {
				address, err := mail.ParseAddress(v[i])

				if err != nil {
					panic(
						NewUsageError(InvalidEmailAddressList, nil).
							WithFlag(flag).
							WithMessage(
								"Invalid email address detected.",
							).
							WithExpected(
								"Expected a valid email address within the email address list.",
							).
							WithReceived(
								fmt.Sprintf("Invalid email address: %s", v[i]),
							).
							WithSuggestion(
								"Provide valid email addresses such as 'user@example.com'.",
							),
					)
				}

				addresses = append(addresses, address)
			}

			return addresses

		default:
			panic(
				NewUsageError(InvalidEmailAddressList, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid email address list detected.",
					).
					WithExpected(
						"Expected a list of valid email addresses.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid list of email addresses.",
					),
			)
		}
	})
}

func NewEmailEnumFlag(name string, enum []string) *Flag {
	flag := NewCustomEnumFlag(name, enum, func(flag flagView, value any) any {
		if value == nil {
			return nil
		}

		_, enum := flag.Enumerable()

		options := make([]string, 0)

		for i := range enum {
			options = append(options, enum[i].(*mail.Address).String())
		}

		switch v := value.(type) {
		case *mail.Address:
			if slices.Contains(options, v.String()) {
				return v
			}

			panic(
				NewUsageError(InvalidEmailAddressOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported email address detected.",
					).
					WithExpected(
						fmt.Sprintf("Expected one of: %s", strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %s", v.String()),
					).
					WithSuggestion(
						"Provide one of the supported enum email addresses.",
					),
			)

		case mail.Address:
			if slices.Contains(options, v.String()) {
				return &v
			}

			panic(
				NewUsageError(InvalidEmailAddressOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported email address detected.",
					).
					WithExpected(
						fmt.Sprintf("Expected one of: %s", strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %s", v.String()),
					).
					WithSuggestion(
						"Provide one of the supported enum email addresses.",
					),
			)

		case string:
			address, err := mail.ParseAddress(v)

			if err != nil {
				panic(
					NewUsageError(InvalidEmailAddress, nil).
						WithFlag(flag).
						WithMessage(
							"Invalid email address detected.",
						).
						WithExpected(
							"Expected a valid email address.",
						).
						WithReceived(
							fmt.Sprintf("Received value: %s", v),
						).
						WithSuggestion(
							"Provide a valid email address such as 'user@example.com'.",
						),
				)
			}

			if slices.Contains(options, address.String()) {
				return address
			}

			panic(
				NewUsageError(InvalidEmailAddressOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported email address detected.",
					).
					WithExpected(
						fmt.Sprintf("Expected one of: %s", strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %s", address.String()),
					).
					WithSuggestion(
						"Provide one of the supported enum email addresses.",
					),
			)

		default:
			panic(
				NewUsageError(InvalidEmailAddressOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported email address detected.",
					).
					WithExpected(
						fmt.Sprintf("Flag '%s' must receive one of: %s", flag.Name(), strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide one of the supported enum email addresses.",
					),
			)
		}
	})

	flag.enum = make([]any, 0)

	for i := range enum {
		address, err := mail.ParseAddress(enum[i])

		if err != nil {
			NewDefinitionError(InvalidEmailAddressOption, nil).
				WithFlag(flag).
				WithMessage(
					"Invalid email enum value detected.",
				).
				WithExpected(
					"All enum options must be valid email addresses.",
				).
				WithReceived(
					fmt.Sprintf("Invalid enum value: %s", enum[i]),
				).
				WithSuggestion(
					"Provide valid email addresses such as 'user@example.com'.",
				).
				Fatal()
		}

		flag.enum = append(flag.enum, address)
	}

	return flag
}
