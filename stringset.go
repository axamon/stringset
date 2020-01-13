// Package stringset allows the creation of sets for strings
// that are concurrecy safe.
package stringset

import (
	"sync"
)

// New creates a new instance of type *Stringset
func New() *StringSet {
	res := &StringSet{
		m: map[string]struct{}{},
	}

	return res
}

// NewStringSet creates a new set for strings.
func NewStringSet(strings ...string) *StringSet {
	res := &StringSet{
		m: map[string]struct{}{},
	}
	res.lock.Lock()
	for _, s := range strings {
		res.m[s] = struct{}{}
	}
	res.lock.Unlock()

	return res
}

// Exists checks if string exists in the set.
func (s *StringSet) Exists(str string) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	_, exists := s.m[str]
	return exists
}

// Delete removes a string from the set.
func (s *StringSet) Delete(str string) {
	s.lock.Lock()
	delete(s.m, str)
	s.lock.Unlock()
}

// Strings returns a slice of strings in the set.
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

// Contains returns true if the given set contains all elements from the other set.
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

// Union returns a new set which contains all elements of the previous ones.
func (s *StringSet) Union(other *StringSet) *StringSet {
	var slen, otherlen int
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		s.lock.Lock()
		slen = len(s.m)
		s.lock.Unlock()
	}()
	go func() {
		defer wg.Done()
		other.lock.Lock()
		otherlen = len(other.m)
		other.lock.Unlock()
	}()
	wg.Wait()

	l := slen + otherlen
	ret := &StringSet{
		m: make(map[string]struct{}, l),
	}

	wg.Add(2)
	go func() {
		defer wg.Done()
		s.lock.Lock()
		for str := range s.m {
			ret.lock.Lock()
			ret.m[str] = struct{}{}
			ret.lock.Unlock()
		}
		s.lock.Unlock()
	}()
	go func() {
		defer wg.Done()
		other.lock.Lock()
		for str := range other.m {
			ret.lock.Lock()
			ret.m[str] = struct{}{}
			ret.lock.Unlock()
		}
		other.lock.Unlock()
	}()
	wg.Wait()

	return ret
}

// Len returns the number of items in the set.
// Cannot be used in for loops.
func (s *StringSet) Len() int {
	s.lock.Lock()
	n := len(s.m)
	s.lock.Unlock()
	return n
}

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

// Difference returns a new set with all elements from the first set less
// all elements from the second one.
func (s *StringSet) Difference(other *StringSet) *StringSet {
	diff := &StringSet{
		m: map[string]struct{}{},
	}
	s.lock.Lock()
	for str := range s.m {
		diff.m[str] = struct{}{}
	}
	s.lock.Unlock()
	other.lock.Lock()
	for str := range other.m {
		delete(diff.m, str)
	}
	other.lock.Unlock()
	return diff
}
