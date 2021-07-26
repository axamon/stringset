package stringset

// Len returns the number of items in the set.
// Cannot be used in for loops.
func (s *StringSet) Len() int {
	s.lock.Lock()
	n := len(s.m)
	s.lock.Unlock()
	return n
}
