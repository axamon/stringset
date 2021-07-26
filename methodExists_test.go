package stringset

import "testing"

func TestStringSet_Exists(t *testing.T) {
	t.Parallel()
	type args struct {
		str string
	}
	tests := []struct {
		name string
		s    *StringSet
		args args
		want bool
	}{
		{"Exists", NewStringSet("pippo", "pluto"), args{"pippo"}, true},
		{"Exists", NewStringSet("pippo", "pluto"), args{"paperino"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Exists(tt.args.str); got != tt.want {
				t.Errorf("StringSet.Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}
