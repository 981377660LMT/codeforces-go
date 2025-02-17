// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1927/problem/G
// https://codeforces.com/problemset/status/1927/problem/G
func Test_cf1927G(t *testing.T) {
	testCases := [][2]string{
		{
			`13
1
1
2
1 1
2
2 1
2
1 2
2
2 2
3
1 1 1
3
3 1 2
3
1 3 1
7
1 2 3 1 2 4 2
7
2 1 1 1 2 3 1
10
2 2 5 1 6 1 8 2 8 2
6
2 1 2 1 1 2
6
1 1 4 1 3 2`,
			`1
2
1
1
1
3
1
2
3
4
2
3
3`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1927G)
}
