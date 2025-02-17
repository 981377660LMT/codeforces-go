// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1100/F
// https://codeforces.com/problemset/status/1100/problem/F
func Test_cf1100F(t *testing.T) {
	testCases := [][2]string{
		{
			`4
7 2 3 4
3
1 4
2 3
1 3`,
			`7
3
7`,
		},
		{
			`5
12 14 23 13 7
15
1 1
1 2
1 3
1 4
1 5
2 2
2 3
2 4
2 5
3 3
3 4
3 5
4 4
4 5
5 5`,
			`12
14
27
27
31
14
25
26
30
23
26
29
13
13
7`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1100F)
}
