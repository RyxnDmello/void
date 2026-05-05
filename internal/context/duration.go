package context

import (
	"time"
)

func getDurationValue(value any) *time.Duration {
	result, ok := value.(time.Duration)

	if !ok {
		return nil
	}

	return &result
}

func getDurationListValue(value any) *[]time.Duration {
	result, ok := value.([]time.Duration)

	if !ok {
		return nil
	}

	return &result
}
