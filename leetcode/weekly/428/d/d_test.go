// Generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	testutil2 "github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_d(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, makeStringGood, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode.cn/contest/weekly-contest-428/problems/minimum-operations-to-make-character-frequencies-equal/
// https://leetcode.cn/problems/minimum-operations-to-make-character-frequencies-equal/

func TestCompareInf(_t *testing.T) {
	//return
	testutil.DebugTLE = 0
	rg := testutil2.NewRandGenerator()
	inputGenerator := func() (s string) {
		//return
		rg.Clear()
		s = rg.Str(1,40,'o','z')
		return
	}

	testutil.CompareInf(_t, inputGenerator, makeStringGoodAC, makeStringGood)
}
