// Package stringset allows the creation of sets for strings
// that are concurrecy safe.
package stringset

import (
	"sync"
)

// StringSet is a set of unique strings.
// The lock sync.RWMutex allows to solve concurrency issues
type StringSet struct {
	m    map[string]struct{}
	lock sync.RWMutex
}

// NewStringSet creates a new set for strings.
func NewStringSet(strings ...string) *StringSet {
	res := &StringSet{
		m: map[string]struct{}{},
	}
	for _, s := range strings {
		res.Add(s)
	}
	return res
}

// Add adds a string to the set.
// If string is already in the set, it has no effect.
func (s *StringSet) Add(str string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.m[str] = struct{}{}
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
	defer s.lock.Unlock()
	delete(s.m, str)
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
	for _, k := range other.Strings() {
		if !s.Exists(k) {
			return false
		}
	}
	return true
}

// Union returns a new set with contains all elements of the previous ones.
func (s *StringSet) Union(other *StringSet) (union *StringSet) {
	otherlist := other.Strings()
	union = s
	for _, k := range otherlist {
		union.Add(k)
	}
	return union
}

// Len returns the number of items in the set.
// Cannot be used in for loops.
func (s *StringSet) Len() int {
	s.lock.Lock()
	n := len(s.m)
	s.lock.Unlock()
	return n
}

// Pop removes and returns an arbitrary element from the set and removes it from the
// set. If the set was empty, this returns ("", false).
func (s *StringSet) Pop() (str string, ok bool) {
	if s.Len() != 0 {
		for _, str = range s.Strings() {
			s.Delete(str) // deletes only one value from the set and than exits
			return str, true
		}
	}
	return "", false
}

// Difference returns a new set with all elements from the first set and no elements from the latter.
func (s *StringSet) Difference(other *StringSet) (diff *StringSet) {
	toremove := other.Strings()
	diff = s
	for _, k := range toremove {
		diff.Delete(k)
	}
	return diff
}

// Intersect returns a new set wich contains only the elemets shared by both input sets.
func (s *StringSet) Intersect(other *StringSet) (intersection *StringSet) {

	slen := s.Len()
	otherlen := other.Len()

	var smaller, greater *StringSet
	if slen > otherlen {
		smaller = other
		greater = s
	}

	if slen <= otherlen {
		smaller = s
		greater = other
	}

	smallerslice := smaller.Strings()

	for _, element := range smallerslice {
		if greater.Exists(element) {
			continue
		}
		smaller.Delete(element)
	}

	return smaller
}
