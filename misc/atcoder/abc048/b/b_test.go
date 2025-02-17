// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 题目：https://atcoder.jp/contests/abc048/tasks/abc048_b
// 提交：https://atcoder.jp/contests/abc048/submit?taskScreenName=abc048_b
// 对拍：https://atcoder.jp/contests/abc048/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc048_b&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [b]")
	testCases := [][2]string{
		{
			`4 8 2`,
			`3`,
		},
		{
			`0 5 1`,
			`6`,
		},
		{
			`9 9 2`,
			`0`,
		},
		{
			`1 1000000000000000000 3`,
			`333333333333333333`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
