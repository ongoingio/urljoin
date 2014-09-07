package urljoin

import (
	"testing"
)

var tests = []struct {
	in  []string
	out string
}{
	{in: []string{"/"}, out: "/"},
	{in: []string{"http://example.com"}, out: "http://example.com"},
	{in: []string{"http://example.com/"}, out: "http://example.com/"},
	{in: []string{"http://example.com", "foo"}, out: "http://example.com/foo"},
	{in: []string{"http://example.com/", "foo"}, out: "http://example.com/foo"},
	{in: []string{"http://example.com/", "/foo"}, out: "http://example.com/foo"},
	{in: []string{"http://example.com/", "foo/"}, out: "http://example.com/foo/"},
}

func TestJoin(t *testing.T) {
	for _, test := range tests {
		url := Join(test.in...)
		if url != test.out {
			t.Fatalf("url should be %s, is %s", test.out, url)
		}
	}
}
