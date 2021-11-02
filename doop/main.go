package main

import (
	"os"
)

func main() {
	tab := os.Args[1:]
	result := 0
	if len(tab) == 3 {
		if len(tab[0]) > 17 || len(tab[2]) > 17 {
			return
		}
		if IsNumeric(tab[0]) == false {
			return
		} else if IsNumeric(tab[2]) == false {
			return
		}
		if tab[1] == "+" {
			result = Atoi(tab[0]) + Atoi(tab[2])
		} else if tab[1] == "-" {
			result = Atoi(tab[0]) - Atoi(tab[2])
		} else if tab[1] == "/" {
			if tab[2] == "0" {
				os.Stderr.WriteString("No division by 0\n")
				os.Exit(0)
				return
			}
			result = Atoi(tab[0]) / Atoi(tab[2])

		} else if tab[1] == "*" {
			result = Atoi(tab[0]) * Atoi(tab[2])
		} else if tab[1] == "%" {
			if tab[2] == "0" {
				os.Stderr.WriteString("No modulo by 0\n")
				os.Exit(0)
				return
			}
			result = Atoi(tab[0]) % Atoi(tab[2])
		} else {
			return
		}
	} else {
		return
	}
	os.Stderr.WriteString(Itoa(result) + "\n")
}

const fastSmalls = true // enable fast path for small integers

// FormatUint returns the string representation of i in the given base,
// for 2 <= base <= 36. The result uses the lower-case letters 'a' to 'z'
// for digit values >= 10.
func FormatUint(i uint64, base int) string {
	if fastSmalls && i < nSmalls && base == 10 {
		return small(int(i))
	}
	_, s := formatBits(nil, i, base, false, false)
	return s
}

// FormatInt returns the string representation of i in the given base,
// for 2 <= base <= 36. The result uses the lower-case letters 'a' to 'z'
// for digit values >= 10.
func FormatInt(i int64, base int) string {
	if fastSmalls && 0 <= i && i < nSmalls && base == 10 {
		return small(int(i))
	}
	_, s := formatBits(nil, uint64(i), base, i < 0, false)
	return s
}

// Itoa is equivalent to FormatInt(int64(i), 10).
func Itoa(i int) string {
	return FormatInt(int64(i), 10)
}

// AppendInt appends the string form of the integer i,
// as generated by FormatInt, to dst and returns the extended buffer.
func AppendInt(dst []byte, i int64, base int) []byte {
	if fastSmalls && 0 <= i && i < nSmalls && base == 10 {
		return append(dst, small(int(i))...)
	}
	dst, _ = formatBits(dst, uint64(i), base, i < 0, true)
	return dst
}

// AppendUint appends the string form of the unsigned integer i,
// as generated by FormatUint, to dst and returns the extended buffer.
func AppendUint(dst []byte, i uint64, base int) []byte {
	if fastSmalls && i < nSmalls && base == 10 {
		return append(dst, small(int(i))...)
	}
	dst, _ = formatBits(dst, i, base, false, true)
	return dst
}

// small returns the string for an i with 0 <= i < nSmalls.
func small(i int) string {
	if i < 10 {
		return digits[i : i+1]
	}
	return smallsString[i*2 : i*2+2]
}

const nSmalls = 100

const smallsString = "00010203040506070809" +
	"10111213141516171819" +
	"20212223242526272829" +
	"30313233343536373839" +
	"40414243444546474849" +
	"50515253545556575859" +
	"60616263646566676869" +
	"70717273747576777879" +
	"80818283848586878889" +
	"90919293949596979899"

const host32bit = ^uint(0)>>32 == 0

const digits = "0123456789abcdefghijklmnopqrstuvwxyz"

// formatBits computes the string representation of u in the given base.
// If neg is set, u is treated as negative int64 value. If append_ is
// set, the string is appended to dst and the resulting byte slice is
// returned as the first result value; otherwise the string is returned
// as the second result value.
//
func formatBits(dst []byte, u uint64, base int, neg, append_ bool) (d []byte, s string) {
	if base < 2 || base > len(digits) {
		panic("strconv: illegal AppendInt/FormatInt base")
	}
	// 2 <= base && base <= len(digits)

	var a [64 + 1]byte // +1 for sign of 64bit value in base 2
	i := len(a)

	if neg {
		u = -u
	}

	// convert bits
	// We use uint values where we can because those will
	// fit into a single register even on a 32bit machine.
	if base == 10 {
		// common case: use constants for / because
		// the compiler can optimize it into a multiply+shift

		if host32bit {
			// convert the lower digits using 32bit operations
			for u >= 1e9 {
				// Avoid using r = a%b in addition to q = a/b
				// since 64bit division and modulo operations
				// are calculated by runtime functions on 32bit machines.
				q := u / 1e9
				us := uint(u - q*1e9) // u % 1e9 fits into a uint
				for j := 4; j > 0; j-- {
					is := us % 100 * 2
					us /= 100
					i -= 2
					a[i+1] = smallsString[is+1]
					a[i+0] = smallsString[is+0]
				}

				// us < 10, since it contains the last digit
				// from the initial 9-digit us.
				i--
				a[i] = smallsString[us*2+1]

				u = q
			}
			// u < 1e9
		}

		// u guaranteed to fit into a uint
		us := uint(u)
		for us >= 100 {
			is := us % 100 * 2
			us /= 100
			i -= 2
			a[i+1] = smallsString[is+1]
			a[i+0] = smallsString[is+0]
		}

		// us < 100
		is := us * 2
		i--
		a[i] = smallsString[is+1]
		if us >= 10 {
			i--
			a[i] = smallsString[is]
		}

	} else if isPowerOfTwo(base) {
		// Use shifts and masks instead of / and %.
		// Base is a power of 2 and 2 <= base <= len(digits) where len(digits) is 36.
		// The largest power of 2 below or equal to 36 is 32, which is 1 << 5;
		// i.e., the largest possible shift count is 5. By &-ind that value with
		// the constant 7 we tell the compiler that the shift count is always
		// less than 8 which is smaller than any register width. This allows
		// the compiler to generate better code for the shift operation.
		shift := uint(TrailingZeros(uint(base))) & 7
		b := uint64(base)
		m := uint(base) - 1 // == 1<<shift - 1
		for u >= b {
			i--
			a[i] = digits[uint(u)&m]
			u >>= shift
		}
		// u < base
		i--
		a[i] = digits[uint(u)]
	} else {
		// general case
		b := uint64(base)
		for u >= b {
			i--
			// Avoid using r = a%b in addition to q = a/b
			// since 64bit division and modulo operations
			// are calculated by runtime functions on 32bit machines.
			q := u / b
			a[i] = digits[uint(u-q*b)]
			u = q
		}
		// u < base
		i--
		a[i] = digits[uint(u)]
	}

	// add sign, if any
	if neg {
		i--
		a[i] = '-'
	}

	if append_ {
		d = append(dst, a[i:]...)
		return
	}
	s = string(a[i:])
	return
}

func isPowerOfTwo(x int) bool {
	return x&(x-1) == 0
}

const deBruijn32 = 0x077CB531

var deBruijn32tab = [32]byte{
	0, 1, 28, 2, 29, 14, 24, 3, 30, 22, 20, 15, 25, 17, 4, 8,
	31, 27, 13, 23, 21, 19, 16, 7, 26, 12, 18, 6, 11, 5, 10, 9,
}

const deBruijn64 = 0x03f79d71b4ca8b09

var deBruijn64tab = [64]byte{
	0, 1, 56, 2, 57, 49, 28, 3, 61, 58, 42, 50, 38, 29, 17, 4,
	62, 47, 59, 36, 45, 43, 51, 22, 53, 39, 33, 30, 24, 18, 12, 5,
	63, 55, 48, 27, 60, 41, 37, 16, 46, 35, 44, 21, 52, 32, 23, 11,
	54, 26, 40, 15, 34, 20, 31, 10, 25, 14, 19, 9, 13, 8, 7, 6,
}

func TrailingZeros(x uint) int {
	if UintSize == 32 {
		return TrailingZeros32(uint32(x))
	}
	return TrailingZeros64(uint64(x))
}

const uintSize = 32
const UintSize = uintSize

// TrailingZeros16 returns the number of trailing zero bits in x; the result is 16 for x == 0.
func TrailingZeros16(x uint16) int {
	if x == 0 {
		return 16
	}
	// see comment in TrailingZeros64
	return int(deBruijn32tab[uint32(x&-x)*deBruijn32>>(32-5)])
}

// TrailingZeros32 returns the number of trailing zero bits in x; the result is 32 for x == 0.
func TrailingZeros32(x uint32) int {
	if x == 0 {
		return 32
	}
	// see comment in TrailingZeros64
	return int(deBruijn32tab[(x&-x)*deBruijn32>>(32-5)])
}

// TrailingZeros64 returns the number of trailing zero bits in x; the result is 64 for x == 0.
func TrailingZeros64(x uint64) int {
	if x == 0 {
		return 64
	}
	// If popcount is fast, replace code below with return popcount(^x & (x - 1)).
	//
	// x & -x leaves only the right-most bit set in the word. Let k be the
	// index of that bit. Since only a single bit is set, the value is two
	// to the power of k. Multiplying by a power of two is equivalent to
	// left shifting, in this case by k bits. The de Bruijn (64 bit) constant
	// is such that all six bit, consecutive substrings are distinct.
	// Therefore, if we have a left shifted version of this constant we can
	// find by how many bits it was shifted by looking at which six bit
	// substring ended up at the top of the word.
	// (Knuth, volume 4, section 7.3.1)
	return int(deBruijn64tab[(x&-x)*deBruijn64>>(64-6)])
}

func Atoi(s string) int {
	var b, y int
	b = 1
	runes := []rune(s)
	nombres := 0
	if !(len(runes)-1 < 0) {
		if runes[0] == 43 {
			runes[0] = 48
		}
		if runes[0] == 45 {
			runes[0] = 48
			y = 1

		}
	}

	for i := 0; i < len(runes); i++ {
		if (int(runes[i]-48) != 0) && (int(runes[i]-48) != 1) && (int(runes[i]-48) != 2) && (int(runes[i]-48) != 3) && (int(runes[i]-48) != 4) && (int(runes[i]-48) != 5) && (int(runes[i]-48) != 6) && (int(runes[i]-48) != 7) && (int(runes[i]-48) != 8) && (int(runes[i]-48) != 9) {
			b = 0
			break

		}
	}

	for i := 0; i < len(runes)-1 && b == 1; i++ {

		if len(runes)-1 < 0 {
			break
		}
		nombres += int(runes[i] - 48)
		nombres = nombres * 10
	}
	if !(len(runes)-1 < 0) {
		nombres += int(runes[len(runes)-1] - 48)
	}
	if y == 1 {
		nombres = nombres * -1
	}
	if b == 1 {
		return (nombres)

	} else {
		return (0)
	}

}

func IsNumeric(s string) bool {
	tab := []rune(s)
	for i := range tab {
		if !(tab[i] > 47 && tab[i] < 58) && (tab[i] != 45) {
			return false
		}
	}
	return true
}
