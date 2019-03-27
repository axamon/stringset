package stringset

import (
	"reflect"
	"sync"
	"testing"
)

func TestStringSet_Delete(t *testing.T) {
	type fields struct {
		m    map[string]struct{}
		lock sync.RWMutex
	}
	type args struct {
		str string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{"topolino", fields{}, args{"topolino"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StringSet{
				m:    tt.fields.m,
				lock: tt.fields.lock,
			}
			s.Delete(tt.args.str)
		})
	}
}

func TestNewStringSet(t *testing.T) {
	type args struct {
		strings []string
	}
	tests := []struct {
		name string
		args args
		want *StringSet
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStringSet(tt.args.strings...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStringSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
