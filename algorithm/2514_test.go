package algorithm

import (
	"fmt"
	"testing"
)

func TestSortDomains(t *testing.T) {
	ds := []*Domain{
		NewDomain("ee.princeton.edu"),
		NewDomain("cs.princeton.edu"),
		NewDomain("princeton.edu"),
		NewDomain("cnn.com"),
		NewDomain("google.com"),
		NewDomain("apple.com"),
		NewDomain("www.cs.princeton.edu"),
		NewDomain("bolle.cs.princeton.edu"),
	}
	output := SortDomains(ds)
	for _, v := range output {
		fmt.Println(v.(*Domain).url)
	}
}
