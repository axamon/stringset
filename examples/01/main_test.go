package main

import (
	"testing"

	"github.com/axamon/stringset"
)

func BenchmarkHavyAdd(b *testing.B) {
	tt1 := stringset.NewStringSet("pippo", "pluto", "paperino")
	for n := 0; n < b.N; n++ {
		HeavyAdd(tt1, n)
	}
}

func BenchmarkH(b *testing.B) {
	tt1 := stringset.NewStringSet("pippo", "pluto", "paperino")
	for n := 0; n < b.N; n++ {
		Heavyload(tt1, n)
	}
}
