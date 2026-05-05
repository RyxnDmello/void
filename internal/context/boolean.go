package context

func getBoolValue(value any) *bool {
	result, ok := value.(bool)

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
