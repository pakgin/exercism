package genericset

// Implement Set[T comparable] as a map[T]struct{}.
type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return Set[T]{}
}

// Add an element to a Set.
func (s Set[T]) Add(elem T) {
	s[elem] = struct{}{}
}

// Check if an element is in a Set.
func (s Set[T]) Contains(elem T) bool {
	_, ok := s[elem]
	return ok
}

// Check if two sets have the same elements.
func (s Set[T]) Equals(other Set[T]) bool {
	for k, _ := range s {
		if !other.Contains(k) {
			return false
		}
	}
	for k, _ := range other {
		if !s.Contains(k) {
			return false
		}
	}
	return true
}
