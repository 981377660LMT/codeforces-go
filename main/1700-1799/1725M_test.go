// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1725/M
// https://codeforces.com/problemset/status/1725/problem/M
func Test_cf1725M(t *testing.T) {
	testCases := [][2]string{
		{
			`5 7
1 2 2
2 4 1
4 1 4
2 5 3
5 4 1
5 2 4
2 1 1`,
			`1 -1 3 4`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1725M)
}
