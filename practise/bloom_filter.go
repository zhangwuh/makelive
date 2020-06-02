package practise

import "math"

type BloomFilter interface {
	Set(data []byte)
	Get(data []byte) bool
}

type bloomFilter struct {
	m int //size = - n log p / (log2 )^2
	bitset []bool
	k int // k = ln2 * m /n
	faultRate float64
}

func NewBloomFilter(faultRate float64, size int) *bloomFilter {
	m := -math.Log(faultRate) * float64(size) / math.Ln2*math.Ln2
	k := math.Ln2 * m / float64(size)
	return &bloomFilter{
		m: int(m),
		k: int(k),
		bitset: make([]bool, int(m)),
		faultRate: faultRate,
	}
}

func (filter *bloomFilter) Set(data []byte) {
	for i := 0;i < filter.k;i++ {
		hash := filter.position(data, i)
		filter.bitset[hash] = true
	}
}

func (filter *bloomFilter) Get(data []byte) bool  {
	for i := 0;i < filter.k;i++ {
		hash := filter.position(data, i)
		if !filter.bitset[hash] {
			return false
		}
	}
	return true
}

func (filter *bloomFilter) position(data []byte, i int) int {
	hs := filter.hash(data, i)
	return int(math.Abs(float64(hs % (filter.m - 1))))
}

func (filter *bloomFilter) hash(data []byte, i int) int {
	hash := 17
	for _, d := range data {
		hash = hash * 31 + int(d) + i
	}
	return hash
}