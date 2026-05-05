package context

func getValue(values map[string]any, name string) any {
	if len(values) == 0 {
		return nil
	}

	return values[name]
}
