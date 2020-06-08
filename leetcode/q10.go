package leetcode

import "mklive.zhangwuh.com/structure"

/*
Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and '*'.

'.' Matches any single character.
'*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).

Note:

s could be empty and contains only lowercase letters a-z.
p could be empty and contains only lowercase letters a-z, and characters like . or *.
*/

const (
	anyChar  = '.'
	wildChar = '*'
)

func isMatch(s string, p string) bool {
	ss := &structure.LinkedStack{}
	ps := &structure.LinkedStack{}
	for _, r := range s {
		ss.Push(r)
	}
	for _, r := range p {
		ps.Push(r)
	}

	return doMatch(ss, ps)
}

func doMatch(ss, ps structure.Stack) bool {
	for !ps.IsEmpty() {
		pattern := ps.Pop()
		if pattern == nil {
			break
		}
		if pattern == anyChar {
			ss.Pop()
			continue
		} else if pattern == wildChar {
			wildChar := ps.Pop()
			if wildChar == nil {
				return false // no character before '*'
			}
			return wildMatch(ss, wildChar.(rune))
		} else { //pattern is character
			source := ss.Pop()
			if source == nil || source != pattern {
				return false
			}
		}
	}
	return ss.IsEmpty()
}

func wildMatch(stack structure.Stack, wildChar rune) bool {
	for !stack.IsEmpty() {
		source := stack.Pop()
		if source == nil {
			return false
		}

		if wildChar == anyChar {
			return true
		}
		if source != wildChar {
			stack.Push(source)
			return false
		}
	}
	return true
}
