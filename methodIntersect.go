package stringset

import "sync"

// Intersect returns a new set which contains only the elemets shared
// by both sets.
func (s *StringSet) Intersect(other *StringSet) *StringSet {

	var wg sync.WaitGroup

	var slen, otherlen int

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

	var intersection *StringSet
	switch {
	case slen >= otherlen:
		intersection = createIntersect(otherlen, other, s)

	default:
		intersection = createIntersect(slen, s, other)
	}
	return intersection
}

func createIntersect(smallerlen int, smaller, greater *StringSet) *StringSet {
	ret := &StringSet{
		m: make(map[string]struct{}, smallerlen),
	}
	// Copy smaller set in ret
	smaller.lock.Lock()
	for str := range smaller.m {
		ret.m[str] = struct{}{}
	}
	smaller.lock.Unlock()

	greater.lock.Lock()
	defer greater.lock.Unlock()
	for element := range ret.m {
		// If element in smaller exists also in greater moves along
		if _, exists := greater.m[element]; exists {
			continue
		}
		// otherwise deletes it also from ret
		ret.lock.Lock()
		delete(ret.m, element)
		ret.lock.Unlock()
	}
	return ret
}
