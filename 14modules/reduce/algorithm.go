package _reduce

func ReduceIntToInt(f func(int, int) int, a []int) int {
	if len(a) == 0 {
		return 0
	}
	b := a[0]
	for i := 1; i < len(a); i++ {
		b = f(b, a[i])
	}
	return b
}
