package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type WeekDate struct {
	re *regexp.Regexp
}

var res []WeekDate

func i2m(m int) time.Month {
	month := []time.Month{
		time.January,
		time.January,
		time.February,
		time.March,
		time.April,
		time.May,
		time.June,
		time.July,
		time.August,
		time.September,
		time.October,
		time.November,
		time.December}
	return month[m]
}

var weekstr [][]string

func Init() {
	// YYYY-MM-DD [ISO 8601]
	re := regexp.MustCompile(`([12]\d\d\d)-([01]\d)-([0-3]\d) ?\(([^)]+)\)`)
	res = append(res, WeekDate{re: re})
	// YYYY/MM/DD
	re = regexp.MustCompile(`([12]\d\d\d)/([01 ]?\d)/([0-3 ]?\d) ?\(([^)]+)\)`)
	res = append(res, WeekDate{re: re})
	re = regexp.MustCompile(`\b(\d\d)/([01 ]?\d)/([0-3 ]?\d) ?\(([^)]+)\)`)
	res = append(res, WeekDate{re: re})
	re = regexp.MustCompile(`^()([01 ]?\d)/([0-3 ]?\d) ?\(([^)]+)\)`) // MM/DD
	res = append(res, WeekDate{re: re})
	re = regexp.MustCompile(`[^0-9/]()([01 ]?\d)/([0-3 ]?\d) ?\(([^)]+)\)`) // MM/DD
	res = append(res, WeekDate{re: re})
	// YYYY年MM月DD日
	re = regexp.MustCompile(`([12]\d\d\d)年([01 ]?\d)月([0-3 ]?\d)日 ?\(([^)]+)\)`)
	res = append(res, WeekDate{re: re})
	re = regexp.MustCompile(`\b(\d\d)年([01 ]?\d)月([0-3 ]?\d)日 ?\(([^)]+)\)`)
	res = append(res, WeekDate{re: re})
	re = regexp.MustCompile(`^()([01 ]?\d)月([0-3 ]?\d)日 ?\(([^)]+)\)`) // MM月DD日
	res = append(res, WeekDate{re: re})
	re = regexp.MustCompile(`[^年]()([01 ]?\d)月([0-3 ]?\d)日 ?\(([^)]+)\)`) // MM月DD日
	res = append(res, WeekDate{re: re})

	weekstr = [][]string{
		{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"},
		{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
		{"日", "月", "火", "水", "木", "金", "土"},
		{"日曜日", "月曜日", "火曜日", "水曜日", "木曜日", "金曜日", "土曜日"},
	}
}

func atoi(s string) int {
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			a, _ := strconv.Atoi(s[i:])
			return a
		}
	}
	return 0
}

func guessYear(now, month int) int {
	// mm-3 .. mm+9
	//       *
	// a b c 1 2 3 4 5 6 7 8 9
	// b c 1 2 3 4 5 6 7 8 9 A
	// c 1 2 3 4 5 6 7 8 9 A B
	// 1 2 3 4 5 6 7 8 9 A B C
	// 2 3 4 5 6 7 8 9 A B C 1
	DF := 3
	if now <= DF {
		if month >= 12-DF+now {
			return -1
		}
	} else {
		if month <= now-DF {
			return +1
		}
	}
	return 0
}

func getYMD(ys, ms, ds string) (int, int, int) {
	y := atoi(ys)
	m := atoi(ms)
	d := atoi(ds)
	if ys == "" {
		// year 未指定. 月から推測
		t := time.Now()
		mm := int(t.Month())
		y = t.Year() + guessYear(m, mm)
	} else if 0 <= y && y < 100 {
		// 2 桁指定
		if y < 90 {
			y += 2000
		} else {
			y += 1900
		}
	}
	return y, m, d
}

func checkWeek(line string, match []string) (string, bool) {
	y, m, d := getYMD(match[1], match[2], match[3])
	// fmt.Printf("org: %d/%d/%d %s\n", y, m, d, w)

	if m < 0 || m > 12 || d < 0 || d > 31 {
		// y 指定されている場合には何かの間違いでしょう...
		return line, len(match[1]) == 4 && 1990 < y && y < 2200
	}
	if y < 1900 || y > 2100 {
		return line, false
	}

	w := match[4]
	t := time.Date(y, i2m(m), d, 12, 0, 0, 0, time.UTC)
	for _, ws := range weekstr {
		for _, ww := range ws {
			if w == ww {
				mx := strings.Replace(match[0], w, ws[t.Weekday()], -1)
				line = strings.Replace(line, match[0], mx, -1)
				// fmt.Printf("new: %d/%d/%d %s %s %v\n", y, m, d, ws[t.Weekday()], t.Weekday(), match[4] != ws[t.Weekday()])
				return line, match[4] != ws[t.Weekday()]
			}
		}
	}
	return line, false
}

func doCheckLine(line string) (string, bool) {
	b := false
	for _, wd := range res {
		matches := wd.re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			ln, bx := checkWeek(line, match)
			b = b || bx
			line = ln
		}
	}

	return line, b
}

func doCheck(fname string, fix bool) (int, error) {
	var f *os.File
	if fname == "" {
		f = os.Stdin
	} else {
		f, err := os.Open(fname)
		if err != nil {
			return 0, err
		}
		defer f.Close()
	}

	reader := bufio.NewReaderSize(f, 4096)
	lines := make([]string, 0)
	b := false
	for i := 1; ; i++ {
		line, _, err := reader.ReadLine()
		// fmt.Printf("in %s\n", string(line))
		if err == io.EOF {
			break
		} else if err != nil {
			return 0, err
		}

		linex, bx := doCheckLine(string(line))
		if bx {
			fmt.Fprintf(os.Stderr, "%s:%d: invalid weekday %s\n", fname, i, string(line))
		}
		lines = append(lines, linex)
		b = b || bx
	}

	if fix && b {
		// 書き込み
		f = nil
		if fname == "" {
			f = os.Stdout
		} else {
			f, err := os.Create(fname)
			if err != nil {
				return 0, err
			}
			defer f.Close()
		}

		for i := 0; i < len(lines); i++ {
			f.WriteString(lines[i] + "\n")
		}
	}
	if b {
		return 1, nil
	} else {
		return 0, nil
	}
}

func showVersion() {
	fmt.Fprintf(os.Stderr, "go-weekday-check v0.0.1\n")
}

var (
	fix     = flag.Bool("fix", false, "")
	version = flag.Bool("v", false, "show version")
)

func run() int {
	flag.Parse()
	if *version {
		showVersion()
		return 0
	}

	args := flag.Args()
	if len(args) == 0 {
		r, err := doCheck("", *fix)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 3
		}
		return r
	}

	ret := 0
	for _, file := range args {
		r, err := doCheck(file, *fix)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 3
		}
		if r != 0 {
			ret = 1
		}
	}

	return ret
}

func main() {
	Init()
	ret := run()
	os.Exit(ret)
}
