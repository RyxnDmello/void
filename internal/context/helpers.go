package context

import (
	"fmt"
)

func mustGetValue(values map[string]any, name string) any {
	if len(values) == 0 {
		panic(fmt.Sprintf("%s' Does Not Exist", name))
	}

	return values[name]
}

func mustGetString(values map[string]any, name string) string {
	value := mustGetValue(values, name)

	result, ok := value.(string)

	if !ok {
		panic(fmt.Sprintf("%s' Cannot Be Converted Into A String", name))
	}

	return result
}

func mustGetNumber(values map[string]any, name string) float64 {
	value := mustGetValue(values, name)

	result, ok := value.(float64)

	if !ok {
		panic(fmt.Sprintf("%s' Cannot Be Converted Into A Number", name))
	}

	return result
}

func mustGetBool(values map[string]any, name string) bool {
	value := mustGetValue(values, name)

	result, ok := value.(bool)

	if !ok {
		panic(fmt.Sprintf("%s' Cannot Be Converted Into A Boolean", name))
	}

	return result
}

func mustGetList(values map[string]any, name string) []string {
	value := mustGetValue(values, name)

	result, ok := value.([]string)

	if !ok {
		panic(fmt.Sprintf("%s' Cannot Be Converted Into A List", name))
	}

	return result
}
