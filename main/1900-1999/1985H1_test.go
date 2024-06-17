// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1985/problem/H1
// https://codeforces.com/problemset/status/1985/problem/H1
func Test_cf1985H1(t *testing.T) {
	testCases := [][2]string{
		{
			`6
1 1
.
4 2
..
#.
#.
.#
3 5
.#.#.
..#..
.#.#.
5 5
#...#
....#
#...#
.....
...##
6 6
.#..#.
#..#..
.#...#
#.#.#.
.#.##.
###..#
6 8
..#....#
.####.#.
###.#..#
.##.#.##
.#.##.##
#..##.#.`,
			`1
6
9
11
15
30`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1985H1)
}