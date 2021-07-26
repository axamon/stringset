package stringset

import (
	"reflect"
	"testing"
)

func TestStringSet_Union(t *testing.T) {
	type args struct {
		other *StringSet
	}
	tests := []struct {
		name string
		s    *StringSet
		args args
		want *StringSet
	}{
		{"Union", NewStringSet("pippo", "pluto"), args{NewStringSet("paperino")}, NewStringSet("pippo", "pluto", "paperino")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Union(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringSet.Union() = %v, want %v", got, tt.want)
			}
		})
	}
}
