// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1730/problem/D
// https://codeforces.com/problemset/status/1730/problem/D?friends=on
func Test_cf1730D(t *testing.T) {
	testCases := [][2]string{
		{
			`7
3
cbc
aba
5
abcaa
cbabb
5
abcaa
cbabz
1
a
a
1
a
b
6
abadaa
adaaba
8
abcabdaa
adabcaba`,
			`YES
YES
NO
YES
NO
NO
YES`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1730D)
}
