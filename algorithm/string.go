package algorithm

//5.1.1键索引计数法

type node struct {
	name string
	key  int
}

type indexCounterSorter struct {
	nodes   []node
	counter []int
	size    int
}

func newIndexCounterSorter(nodes []node, size int) *indexCounterSorter {
	s := &indexCounterSorter{nodes: nodes, size: size}
	s.count()
	s.toIndexes()
	return s
}

func (sorter *indexCounterSorter) count() {
	counter := make([]int, sorter.size+1)
	for _, n := range sorter.nodes {
		counter[n.key+1]++
	}
	sorter.counter = counter
}

func (sorter *indexCounterSorter) toIndexes() {
	for i := 0; i+1 < len(sorter.counter); i++ {
		sorter.counter[i+1] += sorter.counter[i]
	}
}

func (sorter *indexCounterSorter) sort() []node {
	res := make([]node, len(sorter.nodes))
	for _, n := range sorter.nodes {
		res[sorter.counter[n.key]] = n
		sorter.counter[n.key]++
	}
	return res
}
