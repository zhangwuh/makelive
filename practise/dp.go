package practise

import "math"

//给你 k 种面值的硬币，面值分别为 c1, c2 ... ck，每种硬币的数量无限，再给一个总金额 amount，问你最少需要几枚硬币凑出这个金额，如果不可能凑出，算法返回 -1
var conins = []int{1, 2, 5}

func dp(amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	var res = math.MaxInt8
	for _, c := range conins {
		var sub = dp(amount - c)
		if sub == -1 {
			continue
		}
		res = min(res, sub+1)
	}
	if res == math.MaxInt8 {
		return -1
	}
	return res
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}
