package context

func getFloatValue(value any) *float64 {
	result, ok := value.(float64)

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
