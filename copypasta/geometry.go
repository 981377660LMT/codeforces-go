package copypasta

import (
	. "fmt"
	"io"
	"math"
	"sort"
)

// 由于浮点默认是 %g，输出时应使用 Fprintf(out, "%.12f", ans)，这样还可以方便测试

const eps = 1e-6

type vec struct {
	x, y int64
}

func read(in io.Reader) vec {
	var x, y int64
	Fscan(in, &x, &y)
	return vec{x, y}
}

//func (a vec) equals(b vec) bool { return math.Abs(a.x-b.x) < eps && math.Abs(a.y-b.y) < eps }
//func (a vec) less(b vec) bool   { return a.x+eps < b.x || math.Abs(a.x-b.x) < eps && a.y+eps < b.y }
func (a vec) less(b vec) bool { return a.x < b.x || a.x == b.x && a.y < b.y }
func (a vec) add(b vec) vec   { return vec{a.x + b.x, a.y + b.y} }
func (a vec) sub(b vec) vec   { return vec{a.x - b.x, a.y - b.y} }
func (a vec) mul(k int64) vec { return vec{a.x * k, a.y * k} }
func (a vec) div(k int64) vec { return vec{a.x / k, a.y / k} }
func (a vec) len() float64    { return math.Hypot(float64(a.x), float64(a.y)) }
func (a vec) len2() int64     { return a.x*a.x + a.y*a.y }
func (a vec) dot(b vec) int64 { return a.x*b.x + a.y*b.y }
func (a vec) reverse() vec    { return vec{-a.x, -a.y} }
func (a vec) up() vec {
	if a.y < 0 || a.y == 0 && a.x < 0 {
		return a.reverse()
	}
	return a
}
func (a vec) det(b vec) int64 { return a.x*b.y - a.y*b.x }

// + b在a左侧
// - b在a右侧
// 0 ab平行或重合（共基线）
// up() 后按逆时针排序: sort.Slice(ps, func(i, j int) bool { return ps[i].det(ps[j]) > 0 })
// 排序后保证共线的向量是相邻的（因为范围是 [0, 180) ）

func (a vec) mulVec(b vec) vec { return vec{a.x*b.x - a.y*b.y, a.x*b.y + b.x*a.y} }

type line struct {
	p1, p2 vec
}

// 点 a 是否在线段 p1-p2 上（a-p1 与 a-p2 共线且方向相反）
func (a vec) onSeg(l line) bool {
	p1 := l.p1.sub(a)
	p2 := l.p2.sub(a)
	return p1.det(p2) == 0 && p1.dot(p2) <= 0
}

func (a line) intersection(b line) vec {
	va, vb := a.p2.sub(a.p1), b.p2.sub(b.p1)
	k := vb.det(b.p1.sub(a.p1)) / vb.det(a.p2.sub(a.p1))
	return a.p1.add(va.mul(k))
}

func geometryCollection() {
	// 凸包
	// qs[0] == qs[-1]
	convexHull := func(ps []vec) []vec {
		n := len(ps)
		sort.Slice(ps, func(i, j int) bool {
			a, b := ps[i], ps[j]
			return a.x < b.x || a.x == b.x && a.y < b.y
		})
		qs := make([]vec, 2*n)
		for _, pi := range ps {
			for {
				sz := len(qs)
				if sz <= 1 || qs[sz-1].sub(qs[sz-2]).det(pi.sub(qs[sz-1])) > 0 {
					break
				}
				qs = qs[:sz-1]
			}
			qs = append(qs, pi)
		}
		downSize := len(qs)
		for i := n - 2; i >= 0; i-- {
			pi := ps[i]
			for {
				sz := len(qs)
				if sz <= downSize || qs[sz-1].sub(qs[sz-2]).det(pi.sub(qs[sz-1])) > 0 {
					break
				}
				qs = qs[:sz-1]
			}
			qs = append(qs, pi)
		}
		return qs
	}

	// 凸包周长
	convexHullLength := func(ps []vec) (res float64) {
		qs := convexHull(ps)
		for i := 1; i < len(qs); i++ {
			res += qs[i].sub(qs[i-1]).len()
		}
		return
	}

	// 如果是线段的话，还需要判断恰好有四个点，并且没有严格交叉（含重合）
	// tests if angle abc is a right angle
	isOrthogonal := func(a, b, c vec) bool {
		return a.sub(b).dot(c.sub(b)) == 0
	}

	isRectangle := func(a, b, c, d vec) bool {
		return isOrthogonal(a, b, c) &&
			isOrthogonal(b, c, d) &&
			isOrthogonal(c, d, a)
	}

	isRectangleAnyOrder := func(a, b, c, d vec) bool {
		return isRectangle(a, b, c, d) ||
			isRectangle(a, b, d, c) ||
			isRectangle(a, c, b, d)
	}

	_ = []interface{}{convexHullLength, isRectangleAnyOrder}
}

//

type vec3 struct {
	x, y, z int
	idx     int
}

func vec3Collections() {
	var ps []vec3
	sort.Slice(ps, func(i, j int) bool {
		pi, pj := ps[i], ps[j]
		return pi.x < pj.x || pi.x == pj.x && (pi.y < pj.y || pi.y == pj.y && pi.z < pj.z)
	})
}
