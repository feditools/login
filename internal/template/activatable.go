package template

import "regexp"

// ActivableSlice a slice where each element has an active bit which can be set based on a string
type ActivableSlice interface {
	GetChildren(i int) ActivableSlice
	GetMatcher(i int) *regexp.Regexp
	SetActive(i int, a bool)
	Len() int
}

// SetActive sets an active bit in a slice
func SetActive(a ActivableSlice, s string) bool {
	found := false
	for i := 0; i < a.Len(); i++ {
		matcher := a.GetMatcher(i)
		if matcher != nil {
			if matcher.Match([]byte(s)) {
				a.SetActive(i, true)
				found = true
			}
		}
		children := a.GetChildren(i)
		if children != nil {
			foundChild := SetActive(children, s)
			if foundChild {
				a.SetActive(i, true)
				found = true
			}
		}
	}
	return found
}
