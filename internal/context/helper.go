package context

func getValue(values map[string]any, name string) any {
	if len(values) == 0 {
		return nil
	}

	return values[name]
}

func getStringValue(value any) *string {
	result, ok := value.(string)

	if !ok {
		return nil
	}

	return &result
}

func getIntValue(value any) *int {
	result, ok := value.(int)

	if !ok {
		return nil
	}

	return &result
}

func getFloatValue(value any) *float64 {
	result, ok := value.(float64)

	if !ok {
		return nil
	}

	return &result
}

func getBoolValue(value any) *bool {
	result, ok := value.(bool)

	if !ok {
		return nil
	}

	return &result
}

func getStringListValue(value any) *[]string {
	result, ok := value.([]string)

	if !ok {
		return nil
	}

	return &result
}

func getIntListValue(value any) *[]int {
	result, ok := value.([]int)

	if !ok {
		return nil
	}

	return &result
}

func getFloatListValue(value any) *[]float64 {
	result, ok := value.([]float64)

	if !ok {
		return nil
	}

	return &result
}

func getBoolListValue(value any) *[]bool {
	result, ok := value.([]bool)

	if !ok {
		return nil
	}

	return &result
}
