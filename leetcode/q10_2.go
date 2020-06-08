package leetcode

type matcher struct {
	f    func(r rune) (match bool, pskip bool, sskip bool)
	text string
}

func isMatch2(s string, p string) bool {
	sq := &deque{}
	for _, r := range s {
		sq.Add(r)
	}

	pq := resolveMatchers(p)

	for !pq.IsEmpty() {
		m := pq.next()
		sv := sq.next()
		if sv == nil {
			sv = int32(0)
		}
		ss := sv.(rune)
		match, pskip, sskip := m.(*matcher).f(ss)
		if !match {
			return false
		}
		if !pskip {
			pq.putback(m)
		}
		if !sskip {
			sq.putback(sv)
		}
		pq.inverse()
		sq.inverse()
	}
	return sq.IsEmpty()
}

func resolveMatchers(s string) *deque {
	ms := &deque{stackMode: true}
	for i := len(s) - 1; i >= 0; i-- {
		r := rune(s[i])
		if r == anyChar {
			ms.Add(&matcher{func(rr rune) (bool, bool, bool) {
				if rr == 0 {
					return false, true, true
				}
				return true, true, true
			}, string(r)})
		} else if r == wildChar {
			i--
			wild := rune(s[i])
			ms.Add(&matcher{func(rr rune) (bool, bool, bool) {
				if rr == 0 {
					return true, true, true
				}
				if wild == anyChar {
					return true, false, true
				}

				if wild != rr {
					return true, true, false
				}
				return true, false, true
			}, "*" + string(wild)})
		} else {
			ms.Add(&matcher{func(rr rune) (bool, bool, bool) {
				if rr == 0 {
					return false, true, true
				}
				if rr == r {
					return true, true, true
				}
				return false, true, true
			}, string(r)})
		}
	}
	return ms
}
