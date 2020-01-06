package multimain

import (
	"fmt"
	"os"
	"strings"
)

type cmdlineSelector struct{}

func (s cmdlineSelector) Pop() (selected string, ok bool) {
	if len(os.Args) < 2 {
		return "", false
	}

	selected = os.Args[1]
	os.Args = append(os.Args[0:1], os.Args[2:]...)

	return selected, true
}

func (s cmdlineSelector) String() string {
	parts := strings.Split(os.Args[0], string(os.PathSeparator))
	return fmt.Sprintf("command line parameter (e.g. %v build)", parts[len(parts)-1])
}

func FromCmdline() Selector {
	return cmdlineSelector{}
}
