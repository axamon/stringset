package stringset

// Add adds a string to the set.
// If string is already in the set, it has no effect.
func (s *StringSet) Add(str string) {
	s.lock.Lock()
	s.m[str] = struct{}{}
	s.lock.Unlock()
}

// AddSlice adds the elements of the slice to the set.
func (s *StringSet) AddSlice(slice []string) *StringSet {
	s.lock.Lock()
	for _, str := range slice {
		s.m[str] = struct{}{}
	}
	s.lock.Unlock()

	return s
}
