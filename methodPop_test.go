package stringset

import "testing"

func TestStringSet_Pop(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		s     *StringSet
		want  string
		want1 bool
	}{
		{"Pop", NewStringSet("pippo"), "pippo", true},
		{"Pop", NewStringSet(), "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Pop()
			if got != tt.want {
				t.Errorf("StringSet.Pop() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("StringSet.Pop() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
