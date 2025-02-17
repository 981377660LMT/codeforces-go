// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1931/problem/G
// https://codeforces.com/problemset/status/1931/problem/G
func Test_cf1931G(t *testing.T) {
	testCases := [][2]string{
		{
			`11
1 1 1 1
1 2 5 10
4 6 100 200
900000 900000 900000 900000
0 0 0 0
0 0 566 239
1 0 0 0
100 0 100 0
0 0 0 4
5 5 0 2
5 4 0 5`,
			`4
66
0
794100779
1
0
1
0
1
36
126`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1931G)
}
