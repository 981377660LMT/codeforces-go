// Generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_d(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, countGoodArrays, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode.cn/contest/weekly-contest-430/problems/count-the-number-of-arrays-with-k-matching-adjacent-elements/
// https://leetcode.cn/problems/count-the-number-of-arrays-with-k-matching-adjacent-elements/