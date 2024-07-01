package functionMap

func MapIntToInt(f func(int) int, a []int) []int {
	b := make([]int, len(a))
	for i, v := range a {
		b[i] = f(v)
	}
	return b
}
