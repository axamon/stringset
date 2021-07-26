package stringset

import "testing"

func TestStringSet_Contains(t *testing.T) {
	t.Parallel()
	type args struct {
		other *StringSet
	}
	tests := []struct {
		name string
		s    *StringSet
		args args
		want bool
	}{
		{"Contains", NewStringSet("pippo", "pluto", "paperino"), args{NewStringSet("pippo", "paperino")}, true},
		{"Contains", NewStringSet("pippo", "pluto", "paperino"), args{NewStringSet("pluto", "minnie")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Contains(tt.args.other); got != tt.want {
				t.Errorf("StringSet.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
