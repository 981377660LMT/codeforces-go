// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1875/D
// https://codeforces.com/problemset/status/1875/problem/D
func Test_cf1875D(t *testing.T) {
	testCases := [][2]string{
		{
			`4
8
5 2 1 0 3 0 4 0
2
1 2
5
1 0 2 114514 0
8
0 1 2 0 1 2 0 3`,
			`3
0
2
7`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1875D)
}
