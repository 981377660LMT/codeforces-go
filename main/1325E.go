package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1325E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)

	const mx, inf int = 1e6, 1e9
	lpf := [mx + 1]int{1: 1}
	for i := 2; i <= mx; i++ {
		if lpf[i] == 0 {
			for j := i; j <= mx; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
			}
		}
	}

	g := make([][]int, mx)
	var n, v int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		ps := make([]int, 0, 2)
		for v > 1 {
			p := lpf[v]
			cnt := 1
			for v /= p; lpf[v] == p; v /= p {
				cnt++
			}
			if cnt&1 == 1 {
				ps = append(ps, p)
			}
		}
		if len(ps) == 0 {
			Fprint(_w, 1)
			return
		}
		if len(ps) == 1 {
			ps = append(ps, 1)
		}
		v, w := ps[0], ps[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	ans := inf
	dist := make([]int, mx)
	for i := range dist {
		dist[i] = -1
	}
	type pair struct{ v, fa int }
	var p pair
	for st := 1; st < 1000; st++ {
		vs := []int{st}
		dist[st] = 0
		q := []pair{{st, -1}}
		for len(q) > 0 {
			p, q = q[0], q[1:]
			v, fa := p.v, p.fa
			for _, w := range g[v] {
				if dist[w] == -1 {
					dist[w] = dist[v] + 1
					q = append(q, pair{w, v})
					vs = append(vs, w)
				} else if w != fa {
					if l := dist[w] + dist[v] + 1; l < ans {
						ans = l
					}
				}
			}
		}
		for _, v := range vs {
			dist[v] = -1
		}
	}
	if ans == inf {
		ans = -1
	}
	Fprint(_w, ans)
}

//func main() { CF1325E(os.Stdin, os.Stdout) }
