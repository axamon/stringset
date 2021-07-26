package stringset

import "testing"

func TestStringSet_Len(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		s    *StringSet
		want int
	}{
		{"Len", NewStringSet("pippo"), 1},
		{"Len", NewStringSet(), 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); got != tt.want {
				t.Errorf("StringSet.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}
