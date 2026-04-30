package context

import (
	"fmt"
)

func getValue(values map[string]any, name string) any {
	if len(values) == 0 {
		return nil
	}

	return values[name]
}

func mustGetValue(values map[string]any, name string) any {
	if len(values) == 0 {
		panic(fmt.Sprintf("'%s' Does Not Exist", name))
	}

	return values[name]
}

func mustGetString(values map[string]any, name string) string {
	value := mustGetValue(values, name)

	result, ok := value.(string)

	if !ok {
		panic(fmt.Sprintf("'%s' Cannot Be Converted Into A String", name))
	}

	return result
}

func mustGetInteger(values map[string]any, name string) int {
	value := mustGetValue(values, name)

	result, ok := value.(int)

	if !ok {
		panic(fmt.Sprintf("'%s' Cannot Be Converted Into An Integer", name))
	}

	return result
}

func mustGetFloat(values map[string]any, name string) float64 {
	value := mustGetValue(values, name)

	result, ok := value.(float64)

	if !ok {
		panic(fmt.Sprintf("'%s' Cannot Be Converted Into A Float", name))
	}

	return result
}

func mustGetBool(values map[string]any, name string) bool {
	value := mustGetValue(values, name)

	result, ok := value.(bool)

	if !ok {
		panic(fmt.Sprintf("'%s' Cannot Be Converted Into A Boolean", name))
	}

	return result
}

func mustGetList(values map[string]any, name string) []string {
	value := mustGetValue(values, name)

	result, ok := value.([]string)

	if !ok {
		panic(fmt.Sprintf("'%s' Cannot Be Converted Into A List Of Strings", name))
	}

	return result
}

func mustGetIntList(values map[string]any, name string) []int {
	value := mustGetValue(values, name)

	result, ok := value.([]int)

	if !ok {
		panic(fmt.Sprintf("'%s' Cannot Be Converted Into A List Of Integers", name))
	}

	return result
}

func mustGetFloatList(values map[string]any, name string) []float64 {
	value := mustGetValue(values, name)

	result, ok := value.([]float64)

	if !ok {
		panic(fmt.Sprintf("'%s' Cannot Be Converted Into A List Of Floats", name))
	}

	return result
}

func mustGetBoolList(values map[string]any, name string) []bool {
	value := mustGetValue(values, name)

	result, ok := value.([]bool)

	if !ok {
		panic(fmt.Sprintf("'%s' Cannot Be Converted Into A List Of Booleans", name))
	}

	return result
}
