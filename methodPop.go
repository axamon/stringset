package stringset

// Pop returns and removes an arbitrary element from the set.
// If the set is empty it returns "", false.
func (s *StringSet) Pop() (string, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	var str string
	switch len(s.m) {
	case 0:
		return "", false
	default:
		// deletes only one value from the set and than exits
		for str = range s.m {
			delete(s.m, str)
			break
		}
	}
	return str, true
}
