package copypasta

/*
注意 11100001 的情况（特判 lowbit = 1）

Precedence    Operator
    5         *  /  %  <<  >>  &  &^
    4         +  -  |  ^
    3         ==  !=  <  <=  >  >=
    2         &&
    1         ||

a|b = a^b + a&b
*/

// See "math/bits" for more functions
// NOTE: bits.Len(0) returns 0 not 1
func bitsCollection() {
	// ^n+1 = (-1-n)+1 = -n
	lowbit := func(n int64) int64 { return n & -n }

	_ = []interface{}{lowbit}
}

// https://halfrost.com/go_s2_de_bruijn/
