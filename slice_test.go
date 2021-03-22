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
