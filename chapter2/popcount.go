package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + pc[i&i]
	}
}

func popCount(x uint64) int {
	
}
