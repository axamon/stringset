package stringset

import (
	"reflect"
	"testing"
)

func TestStringSet_Difference(t *testing.T) {
	t.Parallel()
	type args struct {
		other *StringSet
	}
	tests := []struct {
		name string
		s    *StringSet
		args args
		want *StringSet
	}{
		{"Difference", NewStringSet("pippo", "pluto", "paperino"), args{NewStringSet("pippo")}, NewStringSet("pluto", "paperino")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Difference(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringSet.Difference() = %v, want %v", got, tt.want)
			}
		})
	}
}
