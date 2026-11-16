package customset

import (
	"fmt"
	"strings"
)

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
type Set struct {
	m    map[string]struct{}
	size int
}

func New() Set {
	return Set{m: make(map[string]struct{}), size: 0}
}

func NewFromSlice(l []string) Set {
	newSet := New()
	for _, e := range l {
		newSet.Add(e)
	}
	return newSet
}

func (s Set) String() string {
	output := make([]string, 0, s.size)
	for k, _ := range s.m {
		output = append(output, fmt.Sprintf("\"%s\"", k))
	}
	return fmt.Sprintf("{%s}", strings.Join(output, ", "))
}

func (s Set) IsEmpty() bool {
	return s.size == 0
}

func (s Set) Has(elem string) bool {
	_, ok := s.m[elem]
	return ok
}

func (s *Set) Add(elem string) {
	_, ok := s.m[elem]
	if ok {
		return
	}

	s.m[elem] = struct{}{}
	s.size++
}

func Subset(s1, s2 Set) bool {
	for elem, _ := range s1.m {
		if !s2.Has(elem) {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	for elem, _ := range s1.m {
		if s2.Has(elem) {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	for elem, _ := range s1.m {
		if !s2.Has(elem) {
			return false
		}
	}
	for elem, _ := range s2.m {
		if !s1.Has(elem) {
			return false
		}
	}
	return true
}

func Intersection(s1, s2 Set) Set {
	s := New()
	for elem, _ := range s1.m {
		if s2.Has(elem) {
			s.Add(elem)
		}
	}
	return s
}

func Difference(s1, s2 Set) Set {
	s := New()
	for elem, _ := range s1.m {
		if !s2.Has(elem) {
			s.Add(elem)
		}
	}
	return s
}

func Union(s1, s2 Set) Set {
	s := New()
	for elem, _ := range s1.m {
		s.Add(elem)
	}
	for elem, _ := range s2.m {
		s.Add(elem)
	}
	return s
}
