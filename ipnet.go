package void

import (
	"fmt"
	"net"
	"slices"
	"strings"
)

const (
	InvalidIPNet       ErrorCode = "INVALID_IP_NET"
	InvalidIPNetList   ErrorCode = "INVALID_IP_NET_LIST"
	InvalidIPNetOption ErrorCode = "INVALID_IP_NET_OPTION"
)

func NewIPNetFlag(name string) *Flag {
	return NewCustomFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil, *net.IPNet:
			return v

		case net.IPNet:
			return &v

		case string:
			if _, net, err := net.ParseCIDR(v); err == nil {
				return net
			}

			panic(
				NewUsageError(InvalidIPNet, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid CIDR or network range detected.",
					).
					WithExpected(
						"Expected a valid CIDR or network range.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %s", v),
					).
					WithSuggestion(
						"Provide a valid CIDR such as '192.168.1.0/24' or '10.0.0.0/8'.",
					),
			)

		default:
			panic(
				NewUsageError(InvalidIPNet, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid CIDR or network range detected.",
					).
					WithExpected(
						"Expected a valid CIDR or network range.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid CIDR or network range.",
					),
			)
		}
	})
}

func NewIPNetListFlag(name string) *Flag {
	return NewCustomListFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil, []*net.IPNet:
			return v

		case []net.IPNet:
			nets := make([]*net.IPNet, 0)

			for i := range v {
				nets = append(nets, &v[i])
			}

			return nets

		case []string:
			nets := make([]*net.IPNet, 0)

			for i := range v {
				_, net, err := net.ParseCIDR(v[i])

				if err != nil {
					panic(
						NewUsageError(InvalidIPNetList, nil).
							WithFlag(flag).
							WithMessage(
								"Invalid CIDR or network range detected.",
							).
							WithExpected(
								"Expected a valid CIDR or network range within the network range list.",
							).
							WithReceived(
								fmt.Sprintf("Invalid network range: %s", v[i]),
							).
							WithSuggestion(
								"Provide valid CIDR values such as '192.168.1.0/24'.",
							),
					)
				}

				nets = append(nets, net)
			}

			return nets

		default:
			panic(
				NewUsageError(InvalidIPNetList, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid network range list detected.",
					).
					WithExpected(
						fmt.Sprintf("Flag '%s' must receive a list of valid CIDR or network range values.", flag.Name()),
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid list of CIDR or network range values.",
					),
			)
		}
	})
}

func NewIPNetEnumFlag(name string, enum []string) *Flag {
	flag := NewCustomEnumFlag(name, enum, func(flag flagView, value any) any {
		if value == nil {
			return nil
		}

		_, enum := flag.Enumerable()

		options := make([]string, 0)

		for i := range enum {
			options = append(options, enum[i].(*net.IPNet).String())
		}

		switch v := value.(type) {
		case *net.IPNet:
			if slices.Contains(options, v.String()) {
				return v
			}

			panic(
				NewUsageError(InvalidIPNetOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported CIDR or network range detected.",
					).
					WithExpected(
						fmt.Sprintf("Expected one of: %s", strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %s", v.String()),
					).
					WithSuggestion(
						"Provide one of the supported enum network ranges.",
					),
			)

		case net.IPNet:
			if slices.Contains(options, v.String()) {
				return &v
			}

			panic(
				NewUsageError(InvalidIPNetOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported CIDR or network range detected.",
					).
					WithExpected(
						fmt.Sprintf("Expected one of: %s", strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %s", v.String()),
					).
					WithSuggestion(
						"Provide one of the supported enum network ranges.",
					),
			)

		case string:
			_, net, err := net.ParseCIDR(v)

			if err != nil {
				panic(
					NewUsageError(InvalidIPNet, nil).
						WithFlag(flag).
						WithMessage(
							"Invalid CIDR or network range detected.",
						).
						WithExpected(
							"Expected a valid CIDR or network range.",
						).
						WithReceived(
							fmt.Sprintf("Received value: %s", v),
						).
						WithSuggestion(
							"Provide a valid CIDR such as '192.168.1.0/24'.",
						),
				)
			}

			if slices.Contains(options, net.String()) {
				return net
			}

			panic(
				NewUsageError(InvalidIPNetOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported CIDR or network range detected.",
					).
					WithExpected(
						fmt.Sprintf("Expected one of: %s", strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %s", net.String()),
					).
					WithSuggestion(
						"Provide one of the supported enum network ranges.",
					),
			)

		default:
			panic(
				NewUsageError(InvalidIPNetOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported CIDR or network range detected.",
					).
					WithExpected(
						fmt.Sprintf("Flag '%s' must receive one of: %s", flag.Name(), strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide one of the supported enum network ranges.",
					),
			)
		}
	})

	flag.enum = make([]any, 0)

	for i := range enum {
		_, net, err := net.ParseCIDR(enum[i])

		if err != nil {
			NewDefinitionError(InvalidIPNetOption, nil).
				WithFlag(flag).
				WithFlag(flag).
				WithMessage(
					"Invalid CIDR enum value detected.",
				).
				WithExpected(
					"All enum options must be valid CIDR or network range values.",
				).
				WithReceived(
					fmt.Sprintf("Invalid enum value: %s", enum[i]),
				).
				WithSuggestion(
					"Provide valid CIDR values such as '192.168.1.0/24'.",
				).
				Fatal()
		}

		flag.enum = append(flag.enum, net)
	}

	return flag
}
