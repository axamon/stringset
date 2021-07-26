package stringset

// Difference restituisce un nuovo set con gli elementi in s meno quelli in o.
func (s *StringSet) Difference(o *StringSet) *StringSet {
	diff := &StringSet{
		m: map[string]struct{}{},
	}
	s.lock.Lock()
	for str := range s.m {
		diff.m[str] = struct{}{}
	}
	s.lock.Unlock()
	o.lock.Lock()
	for str := range o.m {
		delete(diff.m, str)
	}
	o.lock.Unlock()
	return diff
}
