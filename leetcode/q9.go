package leetcode

//判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	var reverse int
	copy := x
	for copy != 0 {
		num := copy % 10
		if reverse > 214748364 || reverse == 214748364 && num > 7 { //overflow
			return false
		}
		reverse = reverse*10 + num
		copy /= 10
	}
	return reverse == x
}
