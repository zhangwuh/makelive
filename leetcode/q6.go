package leetcode

/* 将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N */

func convert(s string, numRows int) string {
	var x, y int
	var res [][]rune
	for i, v := range s {
		if numRows <= 1 {
			return s
		}
		if x%(numRows-1) == 0 {
			if y < numRows {
				res = appendCol(res, y, x)
				res[y][x] = v
				y++
				if y < numRows {
					continue
				}
			}
		} else {
			y = i - (2 * x)
			res = appendCol(res, y, x)
			res[y][x] = v
		}
		y = 0
		x++
	}
	var out []rune
	for _, i := range res {
		for _, j := range i {
			if j != ' ' {
				out = append(out, j)
			}
		}
	}
	return string(out)
}

func appendCol(res [][]rune, x, y int) [][]rune {
	if len(res) < x+1 {
		res = append(res, []rune{})
	}
	for len(res[x]) < y+1 {
		res[x] = append(res[x], ' ')
	}
	return res
}
