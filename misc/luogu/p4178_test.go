// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://www.luogu.com.cn/problem/P4178
func Test_p4178(t *testing.T) {
	testCases := [][2]string{
		{
			`7
1 6 13 
6 3 9 
3 5 7 
4 1 3 
2 4 20 
4 7 2 
10`,
			`5`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, p4178)
}
