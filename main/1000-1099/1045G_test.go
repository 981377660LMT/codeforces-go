// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1045/G
// https://codeforces.com/problemset/status/1045/problem/G?friends=on
func Test_cf1045G(t *testing.T) {
	testCases := [][2]string{
		{
			`3 2
3 6 1
7 3 10
10 5 8`,
			`1`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1045G)
}
