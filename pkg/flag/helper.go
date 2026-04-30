package flag

import (
	"fmt"
	"regexp"
	"strings"
)

func mustCleanName(name string) string {
	name = strings.ToLower(strings.TrimSpace(name))

	regex := regexp.MustCompile(`^(?:--)?([a-zA-Z]+)$`)

	if matches := regex.FindStringSubmatch(name); len(matches) != 0 {
		return matches[1]
	}

	panic(fmt.Sprintf("Flag Cannot Have A Name Such As '%s'", name))
}
