package copypasta

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// 带有 IO 缓冲区的输入输出，适用于绝大多数题目
func bufferIO(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)

	Fprintln(out, n)
}

// 快读，适用于输入量巨大的题目
// 相比 Fscan，每读入 1e6 个 int 可以加速约 400~450ms（Codeforces/AtCoder）
func fastIO(_r io.Reader, _w io.Writer) {
	in := bufio.NewScanner(_r)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// 读一个非负整数
	r := func() (x int) {
		in.Scan()
		for _, b := range in.Bytes() {
			x = x*10 + int(b&15)
		}
		return
	}

	// 读一个整数，支持负数
	r = func() (x int) {
		in.Scan()
		data := in.Bytes()
		if data[0] == '-' {
			for _, b := range data[1:] {
				x = x*10 + int(b&15)
			}
			return -x
		}
		for _, b := range data {
			x = x*10 + int(b&15)
		}
		return
	}

	// 读一个浮点数
	rf := func() float64 {
		in.Scan()
		s := in.Bytes()
		neg := false
		if s[0] == '-' {
			neg = true
			s = s[1:]
		}
		dotPos := len(s) - 1
		f := int64(0)
		for i, b := range s {
			if b == '.' {
				dotPos = i
			} else {
				f = f*10 + int64(b&15)
			}
		}
		if neg {
			f = -f
		}
		return float64(f) / math.Pow10(len(s)-1-dotPos)
	}

	// NOTE: bufio.Scanner 在读长字符串的情况下可能会有奇怪的 bug，所以还是用下面的 fasterIO 吧！（CF827A WA5）
	in.Buffer(nil, 1e9)
	rs := func() []byte { in.Scan(); return in.Bytes() }

	_ = []interface{}{r, rf, rs}
}

// 超快读
// 以 CF1276B（1e6 数据量）为例，测试结果如下：
// bufferIO  670 ms
// fastIO    296 ms
// fasterIO  202 ms
// fasterIO  202 ms (use syscall.Read(syscall.Stdin, buf))
// 选择 4KB 作为缓存块大小的原因 https://stackoverflow.com/questions/6578394/whats-so-special-about-4kb-for-a-buffer-length
// NOTE: 如果只有数字的话，只需要判断字符与 '0' 的关系就行了；有小写字母的话，与 'z' 的大小判断可以省去（对运行耗时无影响）
func fasterIO(_r io.Reader, _w io.Writer) {
	const eof = 0
	out := bufio.NewWriter(_w)
	defer out.Flush()
	_i, _n, buf := 0, 0, make([]byte, 1<<12) // 4KB

	// 读一个字符
	rc := func() byte {
		if _i == _n {
			_n, _ = _r.Read(buf)
			if _n == 0 { // EOF
				return eof
			}
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}

	// 读一个非负整数
	r := func() (x int) {
		b := rc()
		for ; '0' > b || b > '9'; b = rc() {
			// 某些多组数据的题目，不告诉有多少组数据，那么需要额外判断是否读到了 EOF
			if b == eof {
				return
			}
		}
		for ; '0' <= b && b <= '9'; b = rc() {
			x = x*10 + int(b&15)
		}
		return
	}

	// 读一个整数，支持负数
	r = func() (x int) {
		b := rc()
		neg := false
		for ; '0' > b || b > '9'; b = rc() {
			// 某些多组数据的题目，不告诉有多少组数据，那么需要额外判断是否读到了 EOF
			if b == eof {
				return
			}
			if b == '-' {
				neg = true
			}
		}
		for ; '0' <= b && b <= '9'; b = rc() {
			x = x*10 + int(b&15)
		}
		if neg {
			return -x
		}
		return
	}

	// 读一个数字或字母
	r1 := func() byte {
		b := rc()
		for ; '0' > b; b = rc() {
		}
		return b
	}

	// 读一个仅包含小写字母的字符串
	rs := func() (s []byte) {
		b := rc()
		for ; 'a' > b || b > 'z'; b = rc() {
		}
		for ; 'a' <= b && b <= 'z'; b = rc() {
			s = append(s, b)
		}
		return
	}

	// 读一个长度为 n 的仅包含小写字母的字符串
	rsn := func(n int) []byte {
		b := rc()
		for ; 'a' > b || b > 'z'; b = rc() {
		}
		s := make([]byte, 0, n)
		s = append(s, b)
		for i := 1; i < n; i++ {
			s = append(s, rc())
		}
		return s
	}

	// 如果只有/还剩下一个长度未知的字符串（仅包含小写字母）
	readStringUntilEOF := func() (s []byte) {
		// 若之前 Read 过……
		for _i < len(buf) && buf[_i] < 'a' {
			_i++
		}
		s = append(s, buf[_i:]...)

		// 核心是这一段
		for {
			n, _ := _r.Read(buf)
			if n == 0 {
				break
			}
			s = append(s, buf[:n]...)
		}

		// 注意末尾有 \r \n 的情况
		for ; s[len(s)-1] < 'a'; s = s[:len(s)-1] {
		}
		return
	}

	_ = []interface{}{r, r1, rs, rsn, readStringUntilEOF}
}

// 如果题目按照行来读入更方便的话……
func lineIO(_r io.Reader, _w io.Writer) {
	in := bufio.NewScanner(_r)
	in.Buffer(nil, 1e9) // default maxTokenSize is 65536
	out := bufio.NewWriter(_w)
	defer out.Flush()

	for in.Scan() {
		line := in.Bytes()
		// ...
		_ = line
	}
}
