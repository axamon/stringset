package stringset

import "testing"

func TestStringSet_Delete(t *testing.T) {
	t.Parallel()
	type args struct {
		str string
	}
	tests := []struct {
		name string
		s    *StringSet
		args args
	}{
		{"Delete", NewStringSet("pippo", "pluto"), args{"pippo"}},
		{"Delete", NewStringSet("pippo", "pluto"), args{"paperino"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Delete(tt.args.str)
		})
	}
}
