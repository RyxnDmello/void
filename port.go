package void

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

const (
	InvalidPort       ErrorCode = "INVALID_PORT"
	InvalidPortList   ErrorCode = "INVALID_PORT_LIST"
	InvalidPortOption ErrorCode = "INVALID_PORT_OPTION"
)

const (
	minPort = 1
	maxPort = 65535
)

func NewPortFlag(name string) *Flag {
	return NewCustomFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil:
			return v

		case int:
			if minPort <= v && v <= maxPort {
				return v
			}

			panic(
				NewUsageError(InvalidPort, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid network port detected.",
					).
					WithExpected(
						fmt.Sprintf("Expected an integer between %d and %d.", minPort, maxPort),
					).
					WithReceived(
						fmt.Sprintf("Received port: %d", v),
					).
					WithSuggestion(
						"Provide a valid network port within the supported range.",
					),
			)

		case string:
			port, err := strconv.Atoi(v)

			if err != nil {
				panic(
					NewUsageError(InvalidPort, nil).
						WithFlag(flag).
						WithMessage(
							"Invalid network port detected.",
						).
						WithExpected(
							fmt.Sprintf("Expected an integer between %d and %d.", minPort, maxPort),
						).
						WithReceived(
							fmt.Sprintf("Received value: %s", v),
						).
						WithSuggestion(
							"Provide a valid integer network port.",
						),
				)
			}

			if minPort <= port && port <= maxPort {
				return port
			}

			panic(
				NewUsageError(InvalidPort, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid network port detected.",
					).
					WithExpected(
						fmt.Sprintf("Expected an integer between %d and %d.", minPort, maxPort),
					).
					WithReceived(
						fmt.Sprintf("Received port: %d", port),
					).
					WithSuggestion(
						"Provide a valid network port within the supported range.",
					),
			)

		default:
			panic(
				NewUsageError(InvalidPort, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid network port detected.",
					).
					WithExpected(
						"Expected a valid network port.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid integer network port.",
					),
			)
		}
	})
}

func NewPortListFlag(name string) *Flag {
	return NewCustomListFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil:
			return v

		case []int:
			ports := make([]int, 0)

			for i := range v {
				if v[i] < minPort || v[i] > maxPort {
					panic(
						NewUsageError(InvalidPortList, nil).
							WithFlag(flag).
							WithMessage(
								"Invalid network port detected.",
							).
							WithExpected(
								fmt.Sprintf("Expected an integer between %d and %d.", minPort, maxPort),
							).
							WithReceived(
								fmt.Sprintf("Received port: %d", v[i]),
							).
							WithSuggestion(
								"Provide valid network ports within the supported range.",
							),
					)
				}

				ports = append(ports, v[i])
			}

			return ports

		case []string:
			ports := make([]int, 0)

			for i := range v {
				port, err := strconv.Atoi(v[i])

				if err != nil {
					panic(
						NewUsageError(InvalidPortList, nil).
							WithFlag(flag).
							WithMessage(
								"Invalid network port detected.",
							).
							WithExpected(
								fmt.Sprintf("Expected an integer between %d and %d.", minPort, maxPort),
							).
							WithReceived(
								fmt.Sprintf("Received value: %s", v[i]),
							).
							WithSuggestion(
								"Provide valid integer network ports.",
							),
					)
				}

				if port < minPort || port > maxPort {
					panic(
						NewUsageError(InvalidPortList, nil).
							WithFlag(flag).
							WithMessage(
								"Invalid network port detected.",
							).
							WithExpected(
								fmt.Sprintf("Expected an integer between %d and %d.", minPort, maxPort),
							).
							WithReceived(
								fmt.Sprintf("Received port: %d", port),
							).
							WithSuggestion(
								"Provide valid network ports within the supported range.",
							),
					)
				}

				ports = append(ports, port)
			}

			return ports

		default:
			panic(
				NewUsageError(InvalidPortList, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid network port list detected.",
					).
					WithExpected(
						"Expected a list of valid network ports.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid list of integer network ports.",
					),
			)
		}
	})
}

func NewPortEnumFlag(name string, enum []int) *Flag {
	flag := NewCustomEnumFlag(name, enum, func(flag flagView, value any) any {
		if value == nil {
			return nil
		}

		_, enum := flag.Enumerable()

		options := make([]string, 0)

		for i := range enum {
			options = append(options, fmt.Sprintf("%v", enum[i]))
		}

		switch v := value.(type) {
		case int, string:
			raw := fmt.Sprintf("%v", v)

			port, err := strconv.Atoi(raw)

			if err != nil {
				panic(
					NewUsageError(InvalidPortOption, nil).
						WithFlag(flag).
						WithMessage(
							"Invalid network port detected.",
						).
						WithExpected(
							"Expected a valid integer network port.",
						).
						WithReceived(
							fmt.Sprintf("Received value: %s", raw),
						).
						WithSuggestion(
							"Provide a valid integer network port.",
						),
				)
			}

			if slices.Contains(options, raw) {
				return port
			}

			panic(
				NewUsageError(InvalidPortOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported network port detected.",
					).
					WithExpected(
						fmt.Sprintf("Expected one of: %s", strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received port: %d", port),
					).
					WithSuggestion(
						"Provide one of the supported enum network ports.",
					),
			)

		default:
			panic(
				NewUsageError(InvalidPortOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported network port detected.",
					).
					WithExpected(
						fmt.Sprintf("Flag '%s' must receive one of: %s", flag.Name(), strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide one of the supported enum network ports.",
					),
			)
		}
	})

	for i := range enum {
		if enum[i] < minPort || enum[i] > maxPort {
			NewDefinitionError(InvalidPortOption, nil).
				WithFlag(flag).
				WithMessage(
					"Invalid network port enum value detected.",
				).
				WithExpected(
					fmt.Sprintf("All enum options must be between %d and %d.", minPort, maxPort),
				).
				WithReceived(
					fmt.Sprintf("Invalid enum port: %d", enum[i]),
				).
				WithSuggestion(
					"Provide valid network ports within the supported range.",
				).
				Fatal()
		}
	}

	return flag
}
