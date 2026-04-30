package argument

import (
	"fmt"
	"regexp"
	"strings"
)

func mustCleanName(name string) string {
	name = strings.ToLower(strings.TrimSpace(name))

	regex := regexp.MustCompile(`^[a-zA-Z]+$`)

	if valid := regex.MatchString(name); !valid {
		panic(fmt.Sprintf("Argument Cannot Have A Name Such As '%s'", name))
	}

	return name
}
