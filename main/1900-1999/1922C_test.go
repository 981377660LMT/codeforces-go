// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1922/C
// https://codeforces.com/problemset/status/1922/problem/C
func Test_cf1922C(t *testing.T) {
	testCases := [][2]string{
		{
			`1
5
0 8 12 15 20
5
1 4
1 5
3 4
3 2
5 1`,
			`3
8
1
4
14`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1922C)
}
