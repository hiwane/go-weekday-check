package main

import (
	"strconv"
	"testing"
	"time"
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
		{input: "2019-04-28 (Sat)", expect: "2019-04-28 (Sun)"},
		{input: "2019- 4-28 (Tue)", expect: "2019- 4-28 (Tue)"},
		{input: "2019年4月28日 (水)", expect: "2019年4月28日 (日)"},
		{input: "19年4月28日 (水)", expect: "19年4月28日 (日)"},
		{input: "019年4月28日 (水)", expect: "019年4月28日 (水)"},
		{input: "a19年4月28日 (水)", expect: "a19年4月28日 (水)"},
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

func TestAtoi(t *testing.T) {
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

func TestGuessYear(t *testing.T) {
	// now month expect
	tests := [][]int{
		{1, 1, 0},
		{1, 8, 0},
		{1, 9, 0},
		{1, 10, -1},
		{1, 12, -1},
		{2, 1, 0},
		{2, 8, 0},
		{2, 10, 0},
		{2, 11, -1},
		{2, 12, -1},
		{4, 1, 1},
		{4, 2, 0},
		{4, 11, 0},
		{4, 12, 0},
		{5, 1, 1},
		{5, 2, 1},
		{5, 3, 0},
		{5, 11, 0},
		{5, 12, 0},
		{6, 1, 1},
		{6, 3, 1},
		{6, 4, 0},
		{6, 12, 0},
	}
	for _, test := range tests {
		y := guessYear(test[0], test[1])
		if y != test[2] {
			t.Fatalf("expect %d, but actual %d: [%d, %d]", test[2], y, test[0], test[1])
		}
	}
}

func TestGetYMD(t *testing.T) {
	// now month expect
	tm := time.Now()
	tests := [][]int{
		{2019, 4, 30, 2019},
		{2019, 5, 1, 2019},
		{19, 5, 1, 2019},
		{20, 5, 1, 2020},
		{0, 5, 1, 2000},
		{-1, int(tm.Month()), 1, tm.Year()},
	}
	for _, test := range tests {
		var ys string
		if test[0] < 0 {
			ys = ""
		} else {
			ys = strconv.Itoa(test[0])
		}
		y, m, d := getYMD(ys, strconv.Itoa(test[1]), strconv.Itoa(test[2]))
		if y != test[3] {
			t.Fatalf("y: expect %d, but actual %d: [%d, %d, %d]", test[3], y, test[0], test[1], test[2])
		}
		if m != test[1] {
			t.Fatalf("m: expect %d, but actual %d: [%d, %d, %d]", test[1], m, test[0], test[1], test[2])
		}
		if d != test[2] {
			t.Fatalf("d: expect %d, but actual %d: [%d, %d, %d]", test[2], d, test[0], test[1], test[2])
		}
	}
}
