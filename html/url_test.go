package html

import (
	"net/url"
	"testing"
)

func u(raw string) *url.URL {
	u, err := url.Parse(raw)
	if err != nil {
		panic(err)
	}

	return u
}

func TestCanonicalizeURL(t *testing.T) {
	for _, test := range []struct {
		desc   string
		base   *url.URL
		raw    string
		result string
		err    bool
	}{
		{
			desc:   "simple",
			base:   u("https://www.example.com"),
			raw:    "index.html",
			result: "https://www.example.com/index.html",
		},
		{
			desc:   "absolute",
			base:   u("https://www.example.com"),
			raw:    "https://www.other.com",
			result: "https://www.other.com",
		},
		{
			desc:   "no anchor",
			base:   u("https://www.example.com"),
			raw:    "index.html#test-anchor",
			result: "https://www.example.com/index.html",
		},
		{
			desc:   "root",
			base:   u("https://www.example.com/directory/"),
			raw:    "/image.png",
			result: "https://www.example.com/image.png",
		},
		{
			desc: "error",
			base: u("http://www.example.com"),
			raw:  ":blah",
			err:  true,
		},
	} {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			result, err := CanonicalizeURL(test.base, test.raw)

			if err != nil != test.err {
				t.Errorf("got error '%#v', wanted '%#v'", err, test.err)
			}

			if err != nil {
				return
			}

			if result != test.result {
				t.Errorf("got '%s', wanted '%s'", result, test.result)
			}
		})
	}
}
