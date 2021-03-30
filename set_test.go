package stringutil_test

import (
	"testing"

	"github.com/carlmjohnson/stringutil"
)

func TestSymDiffSets(t *testing.T) {
	tests := map[string]struct {
		Old, New       stringutil.Set
		Added, Removed string
	}{
		"empty":  {map[string]bool{}, nil, "", ""},
		"falsey": {map[string]bool{"a": false}, nil, "", ""},
		"falsey-dual": {
			map[string]bool{"a": false}, map[string]bool{"a": false}, "", ""},
		"falsey-diff": {
			map[string]bool{"a": false}, map[string]bool{"b": false}, "", ""},
		"some-true": {
			map[string]bool{"a": true, "b": true, "c": false},
			map[string]bool{"b": true, "c": true, "d": false},
			"c", "a",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			gotadded, gotremoved := stringutil.SymDiffSets(tc.Old, tc.New)

			wantadded := stringutil.ToSet(split(tc.Added))
			wantremoved := stringutil.ToSet(split(tc.Removed))
			assert(t, stringutil.EqSet(wantadded, gotadded),
				"added %#v != %#v", wantadded, gotadded)
			assert(t, stringutil.EqSet(wantremoved, gotremoved),
				"removed %#v != %#v", wantremoved, gotremoved)
		})
	}
}

func TestEqSet(t *testing.T) {
	tests := map[string]struct {
		a, b stringutil.Set
		eq   bool
	}{
		"empty": {nil, nil, true},
		"same": {
			map[string]bool{"a": true},
			map[string]bool{"a": true},
			true},
		"some-false": {
			map[string]bool{"a": true, "b": false},
			map[string]bool{"a": true, "c": false},
			true},
		"diff-length": {
			map[string]bool{"a": true},
			map[string]bool{"a": true, "c": false},
			true},
		"not-eq": {
			map[string]bool{"a": true},
			map[string]bool{"b": true},
			false},
		"no-overlap": {
			map[string]bool{"a": false},
			map[string]bool{"b": true},
			false},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert(t, stringutil.EqSet(tc.a, tc.b) == tc.eq,
				"EqSet(%#v, %#v) != %v", tc.a, tc.b, tc.eq)
		})
	}
}
