package stringset

import "sync"

// Intersect restituisce un nuovo *StringSet con solo gli elementi presenti
// in entrambi s e other: s ∩ other
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
	// copia il set meno numeroso in ret
	smaller.lock.Lock()
	for str := range smaller.m {
		ret.m[str] = struct{}{}
	}
	smaller.lock.Unlock()

	greater.lock.Lock()
	defer greater.lock.Unlock()
	for element := range ret.m {
		// se l'elemento in smaller esiste anche in greater prosegue
		if _, exists := greater.m[element]; exists {
			continue
		}
		// altrimenti lo rimuove da ret.
		ret.lock.Lock()
		delete(ret.m, element)
		ret.lock.Unlock()
	}
	return ret
}
