// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	examples := [][]string{
		{
			`[-1,0,0,2]`, `[1,2,3,4]`, 
			`[5,1,1,1]`,
		},
		{
			`[-1,0,1,0,3,3]`, `[5,4,6,2,1,3]`, 
			`[7,1,1,4,2,1]`,
		},
		{
			`[-1,2,3,0,2,4,1]`, `[2,3,4,5,6,7,8]`, 
			`[1,1,1,1,1,1,1]`,
		},
		{
			`[-1,0,0,0,2]`, `[6,4,3,2,1]`,
			`[5,1,2,1,2]`,
		},
	}
	targetCaseNum :=  -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, smallestMissingValueSubtree, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-258/problems/smallest-missing-genetic-value-in-each-subtree/
