package algorithm

import (
	"strings"

	"mklive.zhangwuh.com/sort"
)

type Domain struct {
	url  string
	subs []string
}

func SortDomains(ds []*Domain) []sort.Comparable {
	qs := sort.NewQuickSort{}
	var input []sort.Comparable
	for _, v := range ds {
		input = append(input, v)
	}
	qs.Sort(input)
	return input
}

func NewDomain(url string) *Domain {
	subs := strings.Split(url, ".")
	d := &Domain{
		url: url,
	}
	for i := len(subs) - 1; i >= 0; i-- {
		d.subs = append(d.subs, subs[i])
	}
	return d
}

func (d *Domain) Compare(to sort.Comparable) int {
	dt := to.(*Domain)
	min := min(d, dt)
	for i := 0; i < len(min.subs); i++ {
		if d.subs[i] < dt.subs[i] {
			return -1
		} else if d.subs[i] > dt.subs[i] {
			return 1
		}
	}
	if min == d {
		return -1
	}
	return 1
}

func min(d *Domain, j *Domain) *Domain {
	if len(d.subs) >= len(j.subs) {
		return j
	}
	return d
}
