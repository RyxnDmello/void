package context

func getIntValue(value any) *int {
	result, ok := value.(int)

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
