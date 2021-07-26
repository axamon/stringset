package stringset

import (
	"reflect"
	"testing"
)

func TestStringSet_Intersect(t *testing.T) {
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
		{
			"Intersect",
			NewStringSet("pippo", "pluto", "paperino"),
			args{NewStringSet("pluto", "paperino", "minnie")},
			NewStringSet("paperino", "pluto")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Intersect(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringSet.Intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}
