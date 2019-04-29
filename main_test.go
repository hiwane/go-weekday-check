package main

import (
	"testing"
)

func TestCheckLine(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{input: "2019/04/28(Sun)", expect: "2019/04/28(Sun)"},
		{input: "2019/04/28(Sat)", expect: "2019/04/28(Sun)"},
		{input: "2019/ 4/28(Sat)", expect: "2019/ 4/28(Sun)"},
		{input: "2019/04/28 (Sat)", expect: "2019/04/28 (Sun)"},
	}

	Init()
	for _, test := range tests {
		s, b := doCheckLine(test.input)
		if s != test.expect {
			t.Fatalf("expect %q, but actual %q: %s", test.expect, s, test.input)
		}
		if b != (test.input != test.expect) {
			t.Fatalf("unexpected bool %v: %s", b, test.input)
		}
	}
}

func Testatoi(t *testing.T) {
	tests := []struct {
		input  string
		expect int
	}{
		{input: "204", expect: 204},
		{input: "  204", expect: 204},
		{input: "04", expect: 04},
	}
	for _, test := range tests {
		a := atoi(test.input)
		if a != test.expect {
			t.Fatalf("expect %q, but actual %q: %s", test.expect, a, test.input)
		}
	}
}
