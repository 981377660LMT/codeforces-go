// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_a(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithFile(t, findTheLongestBalancedSubstring, "a.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode.cn/contest/weekly-contest-339/problems/find-the-longest-balanced-substring-of-a-binary-string/
// https://leetcode.cn/problems/find-the-longest-balanced-substring-of-a-binary-string/
