// Package stringset permette di creare e gestire in parallelo set di stringhe.
package stringset

// New crea una istanza di tipo *Stringset
func New() *StringSet {
	res := &StringSet{
		m: map[string]struct{}{},
	}

	return res
}

// NewStringSet crea una istanza *StringSet avente strings come contenuto.
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

// Strings restituisce la lista di stringhe contenute in s.
func (s *StringSet) Strings() []string {
	s.lock.Lock()
	n := len(s.m)
	s.lock.Unlock()
	if n == 0 {
		return nil
	}
	res := make([]string, 0, n)
	s.lock.Lock()
	for str := range s.m {
		res = append(res, str)
	}
	s.lock.Unlock()
	return res
}
