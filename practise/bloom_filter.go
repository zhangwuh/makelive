package practise

type BloomFilter interface {
	Put(data []byte)
	IsExists(data []byte) bool
}

type bloomFilter struct {
	size int16
	hasher 
}