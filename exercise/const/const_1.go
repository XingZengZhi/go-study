package _const

const Ln2 = 0.693147180559945309417232121458176568075500134360255254120680009
const Log2E = 1 / Ln2 // this is a precise reciprocal
const Billion = 1e9   // float constant
const hardEight = (1 << 100) >> 97

//并行赋值
const beef, two, c = "eat", 2, "veg"
const (
	Monday, Tuesday, Wednesday = 1, 2, 3
	Thursday, Friday, Saturday = 4, 5, 6
)

//用作枚举
const (
	Unknown = 0
	Female  = 1
	Male    = 2
)
