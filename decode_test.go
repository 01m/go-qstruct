package qstruct

import (
	"net/url"
	"testing"
)

type Test struct {
	A string `json:"A"`
	B int
}

func TestDecode(t *testing.T) {
	u := url.Values{}
	u.Set("A", "123")

	test := &Test{}
	_ = Decode(u, test)

	if test.A != "123" {
		t.Fatalf("got: %s, expected: %s\n", test.A, "123")
	}
}
