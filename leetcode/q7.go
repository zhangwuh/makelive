package leetcode

const (
	ceil  = 214748364
	floor = -214748364
)

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	var res int
	for x != 0 {
		low := x % 10
		if res > ceil || (res == ceil && low > 7) {
			return 0
		}
		if res < floor || (res == floor && low < -8) {
			return 0
		}
		x /= 10
		res = res*10 + low
	}
	return res
}
