package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF282E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	type node struct {
		ch [2]*node
		c  int
	}

	var n int
	var ans, pre int64
	Fscan(in, &n)
	a := make([]int64, n)
	root := &node{}
	for i := range a {
		Fscan(in, &a[i])
		// 插入前缀（除了最后一个）
		for j, o := 39, root; j >= 0; j-- {
			b := pre >> j & 1
			if o.ch[b] == nil {
				o.ch[b] = &node{}
			}
			o = o.ch[b]
			o.c++
		}
		pre ^= a[i]
		if pre > ans {
			ans = pre // 前缀最大值
		}
	}
	suf := int64(0)
	for i := n - 1; i >= 0; i-- {
		suf ^= a[i]
		res := int64(0)
		for j, o := 39, root; j >= 0; j-- {
			b := suf >> j & 1
			if o.ch[b^1] != nil && o.ch[b^1].c > 0 {
				res |= 1 << j
				b ^= 1
			}
			o = o.ch[b]
		}
		if res > ans {
			ans = res // 后缀^前缀最大值
		}
		// 删除前缀
		pre ^= a[i]
		for i, o := 39, root; i >= 0; i-- {
			o = o.ch[pre>>i&1]
			o.c--
		}
	}
	Fprint(out, ans)
}

//func main() { CF282E(os.Stdin, os.Stdout) }
