package leetcode

//给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

func longestPalindrome(s string) string {
	size := len(s)
	runes := []rune(s)
	var rr []rune
	for i := size - 1; i >= 0; i-- {
		rr = append(rr, runes[i])
	}
	reverse := string(rr)
	var res string
	for start := 0; start < size; start++ {
		tail := size - 1
		for tail >= start {
			if s[start] == s[tail] {
				s1 := s[start : tail+1]
				s2 := reverse[size-tail-1 : size-start]
				if s1 == s2 {
					current := s[start : tail+1]
					if len(res) < len(current) {
						res = current
					}
					break
				}
			}
			tail--
		}
	}
	return res
}
