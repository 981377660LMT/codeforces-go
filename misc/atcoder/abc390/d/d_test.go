// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 题目：https://atcoder.jp/contests/abc390/tasks/abc390_d
// 提交：https://atcoder.jp/contests/abc390/submit?taskScreenName=abc390_d
// 对拍：https://atcoder.jp/contests/abc390/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc390_d&orderBy=source_length
// 最短：https://atcoder.jp/contests/abc390/submissions?f.Status=AC&f.Task=abc390_d&orderBy=source_length
func Test_d(t *testing.T) {
	testCases := [][2]string{
		{
			`3
2 5 7`,
			`3`,
		},
		{
			`2
100000000000000000 100000000000000000`,
			`2`,
		},
		{
			`6
71 74 45 34 31 60`,
			`84`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
