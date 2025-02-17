// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/914/C
// https://codeforces.com/problemset/status/914/problem/C
func Test_cf914C(t *testing.T) {
	testCases := [][2]string{
		{
			`110
2`,
			`3`,
		},
		{
			`111111011
2`,
			`169`,
		},
		{
			`1
1`,
			`0`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf914C)
}
