package stringset

import (
	"reflect"
	"testing"
)

func TestStringSet_Add(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		s    *StringSet
		args args
		want *StringSet
	}{
		{"Add", NewStringSet("pippo", "paperino"), args{"pluto"}, NewStringSet("pippo", "paperino", "pluto")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); !reflect.DeepEqual(got, tt.want.Len()-1) {
				t.Errorf("StringSet.Difference() = %v, want %v", got, tt.want)
			}
		})
	}
}
