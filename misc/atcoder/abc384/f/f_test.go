// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 题目：https://atcoder.jp/contests/abc384/tasks/abc384_f
// 提交：https://atcoder.jp/contests/abc384/submit?taskScreenName=abc384_f
// 对拍：https://atcoder.jp/contests/abc384/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc384_f&orderBy=source_length
// 最短：https://atcoder.jp/contests/abc384/submissions?f.Status=AC&f.Task=abc384_f&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [f]")
	testCases := [][2]string{
		{
			`2
4 8`,
			`5`,
		},
		{
			`3
51 44 63`,
			`384`,
		},
		{
			`8
577752 258461 183221 889769 278633 577212 392309 326001`,
			`20241214`,
		},
		{
			`3
1 2 3`,
			`14`,
		},
		{
			`3
2 3 4`,
			`20`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
