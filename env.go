package multimain

import (
	"fmt"
	"os"
	"strings"
)

type envSelector struct {
	variable string
}

func (s envSelector) Pop() (selected string, ok bool) {
	value := os.Getenv(s.variable)
	parts := strings.Split(value, "/")

	for i, part := range parts {
		if part == "" {
			continue
		}

		// update environment with remaining parts
		newValue := strings.Join(parts[i+1:], "/")
		err := os.Setenv(s.variable, newValue)
		if err != nil && newValue != "" {
			// TODO: what do we do with this error?
			return "", false
		}

		return part, true
	}

	return "", false
}

func (s envSelector) String() string {
	return fmt.Sprintf("environment variable '%v'", s.variable)
}

func FromEnv(variable string) Selector {
	return envSelector{variable}
}
