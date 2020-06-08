package leetcode

import "math"

/*请你来实现一个 atoi 函数，使其能将字符串转换成整数。

首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。接下来的转化规则如下：

如果第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字字符组合起来，形成一个有符号整数。
假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成一个整数。
该字符串在有效的整数部分之后也可能会存在多余的字符，那么这些字符可以被忽略，它们对函数不应该造成影响。
注意：假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换，即无法进行有效转换。

在任何情况下，若函数不能进行有效的转换时，请返回 0 。

提示：

本题中的空白字符只包括空格字符 ' ' 。
假设我们的环境只能存储 32 位大小的有符号整数，那么其数值范围为 [−231,  231 − 1]。如果数值超过这个范围，请返回  INT_MAX (231 − 1) 或 INT_MIN (−231) 。
*/

func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	for i, c := range str {
		if isSpace(c) {
			continue
		}
		if c == '+' || c == '-' {
			if len(str) <= i+1 {
				return 0
			}
			return atoi(str[i+1:], c == '+')
		}
		if !isNumber(c) {
			return 0
		} else {
			return atoi(str[i:], true)
		}
	}
	return 0
}

func atoi(s string, isPositive bool) int {
	if !isNumber(rune(s[0])) {
		return 0
	}
	//s start with number
	var numbers []rune
	for _, r := range s {
		if !isNumber(r) {
			break
		}
		numbers = append(numbers, r)
	}
	return convertToInt(numbers, isPositive)
}

func convertToInt(numbers []rune, isPositive bool) int {
	var res int
	for _, r := range numbers {
		num := toInt(r)
		if isPositive {
			if res > 214748364 || res == 214748364 && num > 7 { //overflow
				return math.MaxInt32
			}
		} else {
			if res > 214748364 || res == 214748364 && num > 8 { //overflow
				return math.MinInt32
			}
		}
		res = res*10 + num
	}
	if !isPositive {
		return -res
	}
	return res
}

func toInt(r rune) int {
	return int(r - '0')
}

func isSpace(c rune) bool {
	return c == ' '
}

func isNumber(c rune) bool {
	return c <= '9' && c >= '0'
}
