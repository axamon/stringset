package stringset

// Delete elimina la stringa str da s.
func (s *StringSet) Delete(str string) {
	s.lock.Lock()
	delete(s.m, str)
	s.lock.Unlock()
}
