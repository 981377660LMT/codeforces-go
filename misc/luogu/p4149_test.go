// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://www.luogu.com.cn/problem/P4149
func Test_p4149(t *testing.T) {
	testCases := [][2]string{
		{
			`4 3
0 1 1
1 2 2
1 3 4`,
			`2`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, p4149)
}
