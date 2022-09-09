package st

type Set[T comparable] struct {
	m map[T]bool
}

func New[T comparable]() *Set[T] {
	return &Set[T]{m: map[T]bool{}}
}

// Length returns the length of the set.
func (s *Set[T]) Length() int {
	return len(s.m)
}

// Add adds elem to the calling set.
func (s *Set[T]) Add(elem T) {
	s.m[elem] = true
}

// Remove removes elem from the calling set.
func (s *Set[T]) Remove(elem T) {
	delete(s.m, elem)
}

// Has returns true if the calling set contains elem.
func (s *Set[T]) Has(elem T) bool {
	_, ok := s.m[elem]
	return ok
}

// IsSubset returns True if the calling set is subset of s2.
func (s *Set[T]) IsSubset(s2 *Set[T]) bool {
	for e, _ := range s.m {
		if _, ok := s2.m[e]; !ok {
			return false
		}
	}
	return true
}

// FromSlice returns a new set of comparable elements from the slice s.
func FromSlice[T comparable](s []T) *Set[T] {
	newSet := Set[T]{m: map[T]bool{}}
	for _, elem := range s {
		newSet.m[elem] = true
	}
	return &newSet
}

// Union returns a set containing all the elements from the given sets.
func Union[T comparable](sets ...*Set[T]) *Set[T] {
	union := New[T]()
	for _, s := range sets {
		for elem := range s.m {
			union.m[elem] = true
		}
	}
	return union
}

// Intersection returns a set containing only the elements which are present in all sets.
func Intersection[T comparable](sets ...*Set[T]) *Set[T] {
	if len(sets) == 0 {
		return New[T]()
	}
	intersection := sets[0]
	for i := 1; i < len(sets); i++ {
		// Minor optimization - intersection with the empty set is the empty set, we can return.
		if intersection.Length() == 0 || sets[i].Length() == 0 {
			break
		}
		intersection = twoSetsIntersection[T](intersection, sets[i])

	}

	return intersection
}

// Helpers

func twoSetsIntersection[T comparable](s1, s2 *Set[T]) *Set[T] {
	var (
		ss *Set[T]
		ls *Set[T]
	)
	intersection := New[T]()

	if s1.Length() < s2.Length() {
		ss = s1
		ls = s2
	} else {
		ss = s2
		ls = s1
	}

	for v, _ := range ss.m {
		if _, ok := ls.m[v]; ok {
			intersection.Add(v)
		}
	}
	return intersection
}
