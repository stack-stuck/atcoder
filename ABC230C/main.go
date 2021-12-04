package main

import (
	"fmt"
	"strings"
)

var blackPointsMap map[int64]map[int64]bool

func initMap() {
	blackPointsMap = map[int64]map[int64]bool{}
}

func main() {
	var n, a, b, p, q, r, s int64

	fmt.Scan(&n, &a, &b)
	fmt.Scan(&p, &q, &r, &s)

	calc(a, b, n)
	print(p, q, r, s, false)
	return
}

func print(p, q, r, s int64, forTest bool) *[]string {
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
	if forTest {
		return &output
	}

	for i := 0; i < len(output); i++ {
		fmt.Println(output[i])
	}
	return nil
}

func calc(a, b, n int64) {
	initMap()
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
