package multimain

import (
	"fmt"
	"strings"
)

type Selector interface {
	Pop() (selected string, ok bool)
}

type Mapping struct {
	m map[string]func()
}

func MappingFromMap(m map[string]func()) *Mapping {
	return &Mapping{
		m: m,
	}
}

func Main(mapping *Mapping, selector Selector, selectors ...Selector) {
	selectors = append([]Selector{selector}, selectors...)

	var selectorStrs []string

	for _, selector := range selectors {
		selected, ok := selector.Pop()
		if ok {
			f := mapping.m[selected]
			if f == nil {
				panic(fmt.Errorf("selected command (%v) doesn't exist", selected))
			}

			f()
			return
		}

		selectorStrs = append(selectorStrs, fmt.Sprint(selector))
	}

	panic(fmt.Errorf("need to select a command through one of the following: %v",
		strings.Join(selectorStrs, ", ")))
}
