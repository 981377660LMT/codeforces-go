// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_d(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithFile(t, modifiedGraphEdges, "d.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
	if err := testutil.RunFuncWithRandomInput(t, modifiedGraphEdges); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode.cn/contest/weekly-contest-346/problems/modify-graph-edge-weights/
// https://leetcode.cn/problems/modify-graph-edge-weights/
