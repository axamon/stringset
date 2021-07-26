package stringset

// Exists restituisce true se str è presente in s.
func (s *StringSet) Exists(str string) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	_, exists := s.m[str]
	return exists
}
