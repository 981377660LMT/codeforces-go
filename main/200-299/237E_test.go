// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/237/E
// https://codeforces.com/problemset/status/237/problem/E
func Test_cf237E(t *testing.T) {
	testCases := [][2]string{
		{
			`bbaze
3
bzb 2
aeb 3
ba 10`,
			`8`,
		},
		{
			`abacaba
4
aba 2
bcc 1
caa 2
bbb 5`,
			`18`,
		},
		{
			`xyz
4
axx 8
za 1
efg 4
t 1`,
			`-1`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf237E)
}
