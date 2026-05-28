package void

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

const (
	InvalidMemory         ErrorCode = "INVALID_MEMORY"
	InvalidMemoryList     ErrorCode = "INVALID_MEMORY_LIST"
	InvalidMemoryUnit     ErrorCode = "INVALID_MEMORY_UNIT"
	InvalidMemoryOption   ErrorCode = "INVALID_MEMORY_OPTION"
	InvalidMemoryQuantity ErrorCode = "INVALID_MEMORY_QUANTITY"
)

var (
	memoryRegexp  = regexp.MustCompile(`^([0-9]+(?:\.[0-9]+)?)([A-Za-z]+)$`)
	unitsToMemory = map[string]int64{
		"B":   1,
		"KB":  1000,
		"MB":  1000 * 1000,
		"GB":  1000 * 1000 * 1000,
		"TB":  1000 * 1000 * 1000 * 1000,
		"PB":  1000 * 1000 * 1000 * 1000 * 1000,
		"KiB": 1024,
		"MiB": 1024 * 1024,
		"GiB": 1024 * 1024 * 1024,
		"TiB": 1024 * 1024 * 1024 * 1024,
		"PiB": 1024 * 1024 * 1024 * 1024 * 1024,
	}
)

func NewMemoryFlag(name string) *Flag {
	return NewCustomFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil:
			return v

		case string:
			return mustConvMemory(v)

		default:
			panic(
				NewUsageError(InvalidMemory, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid memory value detected.",
					).
					WithExpected(
						"Expected a valid memory size value.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide a valid memory value such as '512MB', '2GiB', or '1TB'.",
					),
			)
		}
	})
}

func NewMemoryListFlag(name string) *Flag {
	return NewCustomListFlag(name, func(flag flagView, value any) any {
		switch v := value.(type) {
		case nil:
			return v

		case []string:
			memories := make([]float64, 0)

			for i := range v {
				memories = append(memories, mustConvMemory(v[i]))
			}

			return memories

		default:
			panic(
				NewUsageError(InvalidMemoryList, nil).
					WithFlag(flag).
					WithMessage(
						"Invalid memory list detected.",
					).
					WithExpected(
						"Expected a list of valid memory size values.",
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide valid memory values such as '512MB', '2GiB', or '1TB'.",
					),
			)
		}
	})
}

func NewMemoryEnumFlag(name string, enum []string) *Flag {
	flag := NewCustomEnumFlag(name, enum, func(flag flagView, value any) any {
		if value == nil {
			return nil
		}

		_, enum := flag.Enumerable()

		options := make([]string, 0)

		for i := range enum {
			options = append(options, fmt.Sprintf("%v", enum[i].(float64)))
		}

		switch v := value.(type) {
		case string:
			if memory := mustConvMemory(v); slices.Contains(options, fmt.Sprintf("%v", memory)) {
				return memory
			}

			panic(
				NewUsageError(InvalidMemoryOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported memory value detected.",
					).
					WithExpected(
						fmt.Sprintf("Expected one of: %s", strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %s", v),
					).
					WithSuggestion(
						"Provide one of the supported enum memory values.",
					),
			)

		default:
			panic(
				NewUsageError(InvalidMemoryOption, nil).
					WithFlag(flag).
					WithMessage(
						"Unsupported memory value detected.",
					).
					WithExpected(
						fmt.Sprintf("Flag '%s' must receive one of: %s", flag.Name(), strings.Join(options, ", ")),
					).
					WithReceived(
						fmt.Sprintf("Received value: %v", v),
					).
					WithSuggestion(
						"Provide one of the supported enum memory values.",
					),
			)
		}
	})

	flag.enum = make([]any, 0)

	for i := range enum {
		flag.enum = append(flag.enum, mustConvMemory(enum[i]))
	}

	return flag
}

func mustConvMemory(value string) float64 {
	matches := memoryRegexp.FindStringSubmatch(value)

	if len(matches) == 0 {
		NewUsageError(InvalidMemory, nil).
			WithMessage(
				"Invalid memory value detected.",
			).
			WithExpected(
				"Expected a valid memory size value.",
			).
			WithReceived(
				fmt.Sprintf("Received value: %s", value),
			).
			WithSuggestion(
				"Use valid memory formats such as '512MB', '2GiB', or '1TB'.",
			).
			Fatal()
	}

	quant := matches[1]
	units := matches[2]

	multiplier, ok := unitsToMemory[units]

	if !ok {
		NewUsageError(InvalidMemoryUnit, nil).
			WithMessage(
				"Invalid memory unit detected.",
			).
			WithExpected(
				"Expected a supported memory unit.",
			).
			WithReceived(
				fmt.Sprintf("Received unit: %s", units),
			).
			WithSuggestion(
				"Use supported units such as B, KB, MB, GB, TB, KiB, MiB, or GiB.",
			).
			Fatal()
	}

	if result, err := strconv.ParseFloat(quant, 64); err == nil {
		return result * float64(multiplier)
	}

	NewUsageError(InvalidMemoryQuantity, nil).
		WithMessage(
			"Invalid memory quantity detected.",
		).
		WithExpected(
			"Expected a valid numeric memory quantity.",
		).
		WithReceived(
			fmt.Sprintf("Received quantity: %s", quant),
		).
		WithSuggestion(
			"Provide a valid numeric memory quantity such as '512', '1.5', or '2048'.",
		).
		Fatal()

	return 0
}
