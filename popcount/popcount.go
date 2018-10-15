// Package popcount returns the number of set bits (bits whose value is 1) in a uint64 value (population count).
package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PCLoop also returns to population count of x but using a loop
func PCLoop(x uint64) int {
	popCount := pc[byte(x>>(0*8))]
	for i := 1; i < 8; i++ {
		popCount += pc[byte(x>>(uint(i*8)))]
	}
	return int(popCount)
}
