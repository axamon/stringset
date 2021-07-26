package stringset

import "sync"

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
