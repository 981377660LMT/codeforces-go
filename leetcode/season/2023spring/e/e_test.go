// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_e(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithFile(t, getSchemeCount, "e.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
	if err := testutil.RunFuncWithRandomInput(t, getSchemeCount); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode.cn/contest/season/2023-spring/problems/1ybDKD/
// https://leetcode.cn/problems/1ybDKD/
