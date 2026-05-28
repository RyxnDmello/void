package void

import (
	"fmt"
	"net"
	"slices"
	"strings"
)

const (
	InvalidIP       ErrorCode = "INVALID_IP"
	InvalidIPList   ErrorCode = "INVALID_IP_LIST"
	InvalidIPOption ErrorCode = "INVALID_IP_OPTION"
)

func NewIPFlag(name string) *Flag {
	return NewCustomFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil, net.IP:
			return v

		case string:
			if ip := net.ParseIP(v); ip != nil {
				return ip
			}

			panic(
				NewUsageError(InvalidIP, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid IP address detected.",
					).
					WithExpected(
						"Expected a valid IPv4 or IPv6 address.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %s", v),
					).
					WithSuggestion(
						"Provide a valid IPv4 or IPv6 address.",
					),
			)

		default:
			panic(
				NewUsageError(InvalidIP, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid IP address detected.",
					).
					WithExpected(
						"Expected a valid IPv4 or IPv6 address.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid IPv4 or IPv6 address.",
					),
			)
		}
	})
}

func NewIPListFlag(name string) *Flag {
	return NewCustomListFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil, []net.IP:
			return v

		case []string:
			ips := make([]net.IP, 0)

			for i := range v {
				ip := net.ParseIP(v[i])

				if ip == nil {
					panic(
						NewUsageError(InvalidIPList, nil).
							WithFlag(flag).
							WithMessage(
								"Invalid IP address detected.",
							).
							WithExpected(
								"Expected a valid IPv4 or IPv6 address within the IP address list.",
							).
							WithReceived(
								fmt.Sprintf("Invalid IP address: %s", v[i]),
							).
							WithSuggestion(
								"Provide valid IPv4 or IPv6 addresses.",
							),
					)
				}

				ips = append(ips, ip)
			}

			return ips

		default:
			panic(
				NewUsageError(InvalidIPList, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid IP address list detected.",
					).
					WithExpected(
						fmt.Sprintf("Flag '%s' must receive a list of valid IPv4 or IPv6 addresses.", flag.Name()),
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid list of IPv4 or IPv6 addresses.",
					),
			)
		}
	})
}

func NewIPEnumFlag(name string, enum []string) *Flag {
	flag := NewCustomEnumFlag(name, enum, func(flag flagView, value any) any {
		if value == nil {
			return nil
		}

		_, enum := flag.Enumerable()

		options := make([]string, 0)

		for i := range enum {
			options = append(options, enum[i].(net.IP).String())
		}

		switch v := value.(type) {
		case net.IP:
			if slices.Contains(options, v.String()) {
				return v
			}

			panic(
				NewUsageError(InvalidIPOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported IP address detected.",
					).
					WithExpected(
						fmt.Sprintf("Expected one of: %s", strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %s", v.String()),
					).
					WithSuggestion(
						"Provide one of the supported enum IP addresses.",
					),
			)

		case string:
			ip := net.ParseIP(v)

			if ip == nil {
				panic(
					NewUsageError(InvalidIP, nil).
						WithFlag(flag).
						WithMessage(
							"Invalid IP address detected.",
						).
						WithExpected(
							"Expected a valid IPv4 or IPv6 address.",
						).
						WithReceived(
							fmt.Sprintf("Received value: %s", v),
						).
						WithSuggestion(
							"Provide a valid IPv4 or IPv6 address.",
						),
				)
			}

			if slices.Contains(options, ip.String()) {
				return ip
			}

			panic(
				NewUsageError(InvalidIPOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported IP address detected.",
					).
					WithExpected(
						fmt.Sprintf("Expected one of: %s", strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %s", ip.String()),
					).
					WithSuggestion(
						"Provide one of the supported enum IP addresses.",
					),
			)

		default:
			panic(
				NewUsageError(InvalidIPOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported IP address detected.",
					).
					WithExpected(
						fmt.Sprintf("Flag '%s' must receive one of: %s", flag.Name(), strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide one of the supported enum IP addresses.",
					),
			)
		}
	})

	flag.enum = make([]any, 0)

	for i := range enum {
		ip := net.ParseIP(enum[i])

		if ip == nil {
			NewDefinitionError(InvalidIPOption, nil).
				WithFlag(flag).
				WithMessage(
					"Invalid IP enum value detected.",
				).
				WithExpected(
					"All enum options must be valid IPv4 or IPv6 addresses.",
				).
				WithReceived(
					fmt.Sprintf("Invalid enum value: %s", enum[i]),
				).
				WithSuggestion(
					"Provide valid IPv4 or IPv6 addresses.",
				).
				Fatal()
		}

		flag.enum = append(flag.enum, ip)
	}

	return flag
}
