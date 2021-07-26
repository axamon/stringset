package stringset

// Contains restituisce true se tutti gli elementi in other sono presenti in s.
func (s *StringSet) Contains(other *StringSet) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	other.lock.Lock()
	defer other.lock.Unlock()
	for str := range other.m {
		if _, exists := s.m[str]; !exists {
			return false
		}
	}
	return true
}
