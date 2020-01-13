package stringset

import "sync"

// StringSet is a set of unique strings.
// The lock sync.RWMutex allows to solve concurrency issues.
type StringSet struct {
	m    map[string]struct{}
	lock sync.RWMutex
}
