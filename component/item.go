package component

import "strings"

type Items map[string][]string

func (i *Items) Append(at string, items ...string) bool {
	if *i == nil {
		*i = Items{}
	}
	(*i)[at] = append((*i)[at], items...)
	return true
}

func (i *Items) Pop(item string) bool {
	for at, items := range *i {
		for j := range items {
			if item == items[j] {
				items = append(items[:j], items[j+1:]...)
				if len(items) > 0 {
					(*i)[at] = items
				} else {
					delete(*i, at)
				}
				return true
			}
		}
	}
	return false
}

func (i *Items) String() string {
	var q bool
	var b strings.Builder
	for on, items := range *i {
		if q {
			b.WriteString("; ")
		}
		b.WriteString(on)
		b.WriteString(": ")
		b.WriteString(strings.Join(items, ", "))
		q = true
	}
	return b.String()
}

func (i *Items) Empty() bool {
	return len(*i) == 0
}
