// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1272/C
// https://codeforces.com/problemset/status/1272/problem/C
func Test_cf1272C(t *testing.T) {
	testCases := [][2]string{
		{
			`7 2
abacaba
a b`,
			`12`,
		},
		{
			`10 3
sadfaasdda
f a d`,
			`21`,
		},
		{
			`7 1
aaaaaaa
b`,
			`0`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1272C)
}
