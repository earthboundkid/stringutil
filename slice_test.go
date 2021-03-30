package stringutil_test

import (
	"strings"
	"testing"

	"github.com/carlmjohnson/stringutil"
)

func assert(t *testing.T, assertion bool, format string, args ...interface{}) {
	t.Helper()
	if !assertion {
		t.Fatalf(format, args...)
	}
}

func TestUnique(t *testing.T) {
	cases := map[string]struct {
		input, want []string
	}{
		"empty": {[]string{}, []string{}},
		"one":   {strings.Fields("a"), strings.Fields("a")},
		"two":   {strings.Fields("b a"), strings.Fields("b a")},
		"dupes": {strings.Fields("a a b a b"), strings.Fields("a b")},
	}
	for name, tc := range cases {
		t.Run(name+"-nil", func(t *testing.T) {
			got := stringutil.Unique(nil, tc.input)
			assert(t, stringutil.EqSlice(got, tc.want), "want %v; got %v", tc.want, got)
		})
	}
	for name, tc := range cases {
		t.Run(name+"-notnil", func(t *testing.T) {
			base := []string{"a"}
			got := stringutil.Unique(base, tc.input)
			want := append(base, tc.want...)
			assert(t, stringutil.EqSlice(got, want), "want %v; got %v", want, got)
		})
	}
	for name, tc := range cases {
		t.Run(name+"-extracap", func(t *testing.T) {
			base := make([]string, 1, 100)
			got := stringutil.Unique(base, tc.input)
			want := append(base, tc.want...)
			assert(t, stringutil.EqSlice(got, want), "want %v; got %v", want, got)
		})
	}
	for name, tc := range cases {
		t.Run(name+"-inplace", func(t *testing.T) {
			input := stringutil.Copy(tc.input)
			got := stringutil.Unique(input[:0], input)
			assert(t, stringutil.EqSlice(got, tc.want), "want %v; got %v", tc.want, got)
		})
	}
}

var sink []string

func Benchmark(b *testing.B) {
	var inputs [][]string
	data := strings.Fields("a b c d e f a a a b b g g h")
	for i := 0; i < b.N; i++ {
		inputs = append(inputs, stringutil.Copy(data))
	}
	b.ResetTimer()
	m := make(stringutil.Set)
	for i := 0; i < b.N; i++ {
		stringutil.Zero(m)
		sink = stringutil.UniqueVia(inputs[i][:0], inputs[i], m)
	}
}

func TestEqSlice(t *testing.T) {
	tests := map[string]struct {
		a, b string
		eq   bool
	}{
		"empty":      {"", "", true},
		"blank":      {"-", "-", true},
		"blankfirst": {",a", ",a", true},
		"difflength": {"a,b", "a", false},
		"difforder":  {"a,b", "b,a", false},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			as := split(tc.a)
			bs := split(tc.b)
			assert(t, stringutil.EqSlice(as, bs) == tc.eq,
				"EqSlice(%#v, %#v) != %v", as, bs, tc.eq)
		})
	}
}

func split(s string) []string {
	ss := strings.Split(s, ",")
	if s == "" {
		return nil
	}
	if s == "-" {
		return []string{""}
	}
	return ss
}

func TestSymDiff(t *testing.T) {
	tests := map[string]struct {
		Old, New, Added, Removed string
	}{
		"empty":       {"", "", "", ""},
		"none":        {"a,b,c", "a,b,c", "", ""},
		"adding":      {"", "a", "a", ""},
		"removing":    {"a", "", "", "a"},
		"swap":        {"a", "b", "b", "a"},
		"some_remain": {"a,b,c", "b,c,d", "d", "a"},
		"doubles":     {"a,a,b,c,c", "b,b,c,d,d", "d", "a"},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			old := split(tc.Old)
			new := split(tc.New)
			wantadded := split(tc.Added)
			wantremoved := split(tc.Removed)
			gotadded, gotremoved := stringutil.SymDiff(old, new)
			stringutil.Sort(gotadded)
			assert(t, stringutil.EqSlice(wantadded, gotadded),
				"%#v != %#v", wantadded, gotadded)
			stringutil.Sort(gotremoved)
			assert(t, stringutil.EqSlice(wantremoved, gotremoved),
				"%#v != %#v", wantremoved, gotremoved)
		})
	}
}
