// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/277/A
// https://codeforces.com/problemset/status/277/problem/A
func TestCF277A(t *testing.T) {
	testCases := [][2]string{
		{
			`5 5
1 2
2 2 3
2 3 4
2 4 5
1 5`,
			`0`,
		},
		{
			`8 7
0
3 1 2 3
1 1
2 5 4
2 6 7
1 3
2 7 4
1 1`,
			`2`,
		},
		{
			`2 2
1 2
0`,
			`1`,
		},
		{
			`2 2
0
0`,
			`2`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, CF277A)
}
