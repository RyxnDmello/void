package context

func getStringValue(value any) *string {
	result, ok := value.(string)

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
