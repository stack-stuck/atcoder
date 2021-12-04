package main

import (
	"fmt"
	"strings"
)

var blackPointsMap map[int64]map[int64]bool

func main() {
	blackPointsMap = map[int64]map[int64]bool{}
	var n, a, b, p, q, r, s int64
	testPattern := 3
	testPattern1 := func() {
		n = 5
		a = 3
		b = 2
		p = 1
		q = 5
		r = 1
		s = 5
	}
	testPattern2 := func() {
		n = 5
		a = 3
		b = 3
		p = 4
		q = 5
		r = 2
		s = 5
	}
	testPattern3 := func() {
		n = 1000000000000000000
		a = 999999999999999999
		b = 999999999999999999
		p = 999999999999999998
		q = 1000000000000000000
		r = 999999999999999998
		s = 1000000000000000000
	}

	if testPattern == 0 {
		fmt.Scan(&n, &a, &b)
		fmt.Scan(&p, &q, &r, &s)
	} else if testPattern == 1 {
		testPattern1()
	} else if testPattern == 2 {
		testPattern2()
	} else if testPattern == 3 {
		testPattern3()
	}

	calcRule(a, b, n)
	var output []string
	for i := p; i <= q; i++ {
		var outputRow string
		if row, ok := blackPointsMap[i]; !ok {
			output = append(output, strings.Repeat(".", int(s-r+1))) // todo: overflow?
			continue
		} else {
			for j := r; j <= s; j++ {
				if _, ok := row[j]; ok {
					outputRow += "#"
				} else {
					outputRow += "."
				}
			}
			output = append(output, outputRow)
		}
	}
	for i := 0; i < len(output); i++ {
		fmt.Println(output[i])
	}
	return
}

func calcRule(a, b, n int64) {
	calcRule1(a, b, n)
	calcRule2(a, b, n)
	return
}

func calcRule1(a, b, n int64) {
	lower := max(1-a, 1-b)
	upper := min(n-a, n-b)
	if lower > upper {
		return
	}

	for i := lower; i <= upper; i++ {
		if a+i > n || b+i > n {
			break
		}
		if a+i <= 0 || b+i <= 0 {
			continue
		}
		if blackPointsMap[a+i] == nil {
			blackPointsMap[a+i] = map[int64]bool{}
		}
		blackPointsMap[a+i][b+i] = true
	}
}

func calcRule2(a, b, n int64) {
	lower := max(1-a, b-n)
	upper := min(n-a, b-1)
	if lower > upper {
		return
	}

	for i := lower; i <= upper; i++ {
		if a+i > n || b-i < 0 {
			break
		}
		if a+i <= 0 || b-i > n {
			continue
		}
		if blackPointsMap[a+i] == nil {
			blackPointsMap[a+i] = map[int64]bool{}
		}
		blackPointsMap[a+i][b-i] = true
	}
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
