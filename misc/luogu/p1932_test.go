// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://www.luogu.com.cn/problem/P1932
func Test_p1932(t *testing.T) {
	testCases := [][2]string{
		{
			`1
1`,
			`2
0
1
1
0`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, p1932)
}
