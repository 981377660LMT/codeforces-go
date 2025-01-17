// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 题目：https://atcoder.jp/contests/abc252/tasks/abc252_e
// 提交：https://atcoder.jp/contests/abc252/submit?taskScreenName=abc252_e
// 对拍：https://atcoder.jp/contests/abc252/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc252_e&orderBy=source_length
// 最短：https://atcoder.jp/contests/abc252/submissions?f.Status=AC&f.Task=abc252_e&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [e]")
	testCases := [][2]string{
		{
			`3 3
1 2 1
2 3 2
1 3 10`,
			`1 2`,
		},
		{
			`4 6
1 2 1
1 3 1
1 4 1
2 3 1
2 4 1
3 4 1`,
			`3 1 2`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
