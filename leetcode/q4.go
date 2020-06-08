package leetcode

func findMedianSortedArrays(n1 []int, n2 []int) float64 {
	merged := merge(n1, n2)
	if len(merged)%2 == 1 {
		return float64(merged[len(merged)/2])
	}
	middle := len(merged)/2 - 1
	return float64(merged[middle]+merged[middle+1]) / 2
}

//O(m+n)
func merge(n1 []int, n2 []int) []int {
	var merged []int
	var m, n int
	for {
		if m == len(n1) && n < len(n2) {
			merged = append(merged, n2[n:]...)
			break
		}
		if m < len(n1) && n == len(n2) {
			merged = append(merged, n1[m:]...)
			break
		}
		if n1[m] <= n2[n] {
			merged = append(merged, n1[m])
			m++
		} else {
			merged = append(merged, n2[n])
			n++
		}
	}
	return merged
}
