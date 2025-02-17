// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 题目：https://atcoder.jp/contests/abc100/tasks/abc100_d
// 提交：https://atcoder.jp/contests/abc100/submit?taskScreenName=abc100_d
// 对拍：https://atcoder.jp/contests/abc100/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc100_d&orderBy=source_length
// 最短：https://atcoder.jp/contests/abc100/submissions?f.Status=AC&f.Task=abc100_d&orderBy=source_length
func Test_d(t *testing.T) {
	testCases := [][2]string{
		{
			`5 3
3 1 4
1 5 9
2 6 5
3 5 8
9 7 9`,
			`56`,
		},
		{
			`5 3
1 -2 3
-4 5 -6
7 -8 -9
-10 11 -12
13 -14 15`,
			`54`,
		},
		{
			`10 5
10 -80 21
23 8 38
-94 28 11
-26 -2 18
-69 72 79
-26 -86 -54
-72 -50 59
21 65 -32
40 -94 87
-62 18 82`,
			`638`,
		},
		{
			`3 2
2000000000 -9000000000 4000000000
7000000000 -5000000000 3000000000
6000000000 -1000000000 8000000000`,
			`30000000000`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
