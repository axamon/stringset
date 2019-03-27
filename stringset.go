package stringset

import (
	"sync"
)

// StringSet is a set of unique strings
// lock allows to solve concurrency issues
type StringSet struct {
	m    map[string]struct{}
	lock sync.RWMutex
}

// NewStringSet creates a new set for strings
func NewStringSet(strings ...string) *StringSet {
	res := &StringSet{
		m: map[string]struct{}{},
	}
	for _, s := range strings {
		res.Add(s)
	}
	return res
}

// Add adds a string to the set. If string is already in the set, it has no effect
func (s *StringSet) Add(str string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.m[str] = struct{}{}
}

// Exists checks if string exists in the set
func (s *StringSet) Exists(str string) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	_, exists := s.m[str]
	return exists
}

// Delete removes a string from the set
func (s *StringSet) Delete(str string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.m, str)
}

// Strings returns strings in the set
func (s *StringSet) Strings() []string {
	s.lock.Lock()
	n := len(s.m)
	s.lock.Unlock()
	if n == 0 {
		return nil
	}
	// for efficiency, pre-allocate the array with known, final capacity
	// this avoids re-allocating underlying array in append
	res := make([]string, 0, n)
	s.lock.Lock()
	for str := range s.m {
		res = append(res, str)
	}
	s.lock.Unlock()
	return res
}
