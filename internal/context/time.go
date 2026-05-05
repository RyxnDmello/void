package context

import (
	"time"
)

func getTimeValue(value any) *time.Time {
	result, ok := value.(time.Time)

	if !ok {
		return nil
	}

	return &result
}

func getTimeListValue(value any) *[]time.Time {
	result, ok := value.([]time.Time)

	if !ok {
		return nil
	}

	return &result
}
