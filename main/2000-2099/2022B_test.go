// Generated by copypasta/template/generator_test.go
package main

import (
	"fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

// https://codeforces.com/problemset/problem/2022/B
// https://codeforces.com/problemset/status/2022/problem/B?friends=on
func Test_cf2022B(t *testing.T) {
	testCases := [][2]string{
		{
			`4
3 2
3 1 2
3 3
2 1 3
5 3
2 2 1 9 2
7 4
2 5 3 3 5 2 5`,
			`3
3
9
7`,
		},
		{
			`1
2 1
1 2`,
			`3`,
		},
		{
			`1
4 3
1 5 5 1`,
			`5`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf2022B)
}

func TestCompare_cf2022B(_t *testing.T) {
	//return
	testutil.DebugTLE = 0
	rg := testutil.NewRandGenerator()
	inputGenerator := func() string {
		//return ``
		rg.Clear()
		rg.One()
		n := rg.Int(1, 8)
		rg.Int(1, 8)
		rg.NewLine()
		rg.IntSlice(n, 1, 8)
		return rg.String()
	}

	testutil.AssertEqualRunResultsInf(_t, inputGenerator, solve_cf2022B, cf2022B)
}

func solve_cf2022B(in io.Reader, out io.Writer) {
	var n, x, sum, mx int
	fmt.Fscan(in, &n, &n, &x)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		sum += a[i]
		mx = max(mx, a[i])
	}
	ans := max(mx, (sum+x-1)/x)
	fmt.Fprintln(out, ans)
}