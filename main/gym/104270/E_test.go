// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/gym/104270/problem/E
// https://codeforces.com/gym/104270/status/E?friends=on
func Test_cfE(t *testing.T) {
	testCases := [][2]string{
		{
			`2
4 8
3 2 6 6
3 9
10 10 1`,
			`6
4`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, runE)
}
