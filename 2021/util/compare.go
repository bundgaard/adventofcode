package util

func MaxMap(in map[byte]int) byte {
	val := Max(in['0'], in['1'])
	for k, v := range in {
		if v == val {
			return k
		}
	}
	return 0
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MinMap(in map[byte]int) byte {
	val := Min(in['0'], in['1'])
	for k, v := range in {
		if v == val {
			return k
		}
	}
	return 1
}
