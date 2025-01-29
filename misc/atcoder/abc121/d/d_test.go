// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 题目：https://atcoder.jp/contests/abc121/tasks/abc121_d
// 提交：https://atcoder.jp/contests/abc121/submit?taskScreenName=abc121_d
// 对拍：https://atcoder.jp/contests/abc121/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc121_d&orderBy=source_length
// 最短：https://atcoder.jp/contests/abc121/submissions?f.Status=AC&f.Task=abc121_d&orderBy=source_length
func Test_d(t *testing.T) {
	testCases := [][2]string{
		{
			`2 4`,
			`5`,
		},
		{
			`123 456`,
			`435`,
		},
		{
			`123456789012 123456789012`,
			`123456789012`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
