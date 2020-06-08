package algorithm

type UnionFind interface {
	union(p, q int)
	find(p int) int
	connected(p, q int) bool
	unionCount() int
}

type unionFind struct {
	points []int // points[p] = unionId
	count  int
}

func (uf *unionFind) union(p, q int) {
	pid := uf.find(p)
	qid := uf.find(q)
	if pid == qid {
		return
	}
	for i, v := range uf.points {
		if v == pid {
			uf.points[i] = qid
		}
	}
	uf.count--
}

func (uf *unionFind) find(p int) int {
	return uf.points[p]
}

func (uf *unionFind) connected(p, q int) bool {
	return uf.find(p) == uf.find(q)
}

func (uf *unionFind) unionCount() int {
	return uf.count
}

//  0 < point <= size
func NewUnionFind(size int) UnionFind {
	uf := &unionFind{}
	for i := 0; i < size; i++ {
		uf.points = append(uf.points, i)
	}
	uf.count = size
	return uf
}

type forestUnionFind struct {
	points []int // points[p] = unionId
	count  int
}

//  0 < point <= size
func NewTreeUnionFind(size int) UnionFind {
	uf := &forestUnionFind{}
	for i := 0; i < size; i++ {
		uf.points = append(uf.points, i)
	}
	uf.count = size
	return uf
}

func (uf *forestUnionFind) union(p, q int) {
	pid := uf.find(p)
	qid := uf.find(q)
	if pid == qid {
		return
	}
	uf.points[pid] = qid
	uf.count--
}

func (uf *forestUnionFind) find(p int) (root int) {
	id := uf.points[p]
	if id == p { //root
		return p
	}
	defer func() {
		uf.points[p] = root
	}()
	root = uf.find(id) //flat all the nodes to root
	return
}

func (uf *forestUnionFind) connected(p, q int) bool {
	return uf.find(p) == uf.find(q)
}

func (uf *forestUnionFind) unionCount() int {
	return uf.count
}
