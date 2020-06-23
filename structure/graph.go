package structure

import (
	"fmt"

	"mklive.zhangwuh.com/utils"
)

type Graph interface {
	V() int                          //vertex amount
	E() int                          //edge amount
	AddEdge(v1, v2 interface{})      //add edge between v1 and v2
	Adj(v interface{}) []interface{} // all vertex connected to v
}

type adjNode struct {
	edges []interface{}
}

// adjacency-lists based graph
type AdjGraph struct {
	vertex  int
	edge    int
	adjList []*adjNode
}

type vertex struct{}

type edge struct {
	v, w int // v :edge from, w:edge to
}

func NewAdjGraph(v int, edges []edge) *AdjGraph {
	g := &AdjGraph{
		vertex:  v,
		adjList: newAdjNodes(v),
	}
	for _, e := range edges {
		g.AddEdge(e.v, e.w)
	}
	return g
}

func newAdjNodes(v int) []*adjNode {
	var nodes []*adjNode
	for i := 0; i < v; i++ {
		nodes = append(nodes, &adjNode{})
	}
	return nodes
}

func (g *AdjGraph) V() int {
	return g.vertex
}

func (g *AdjGraph) E() int {
	return g.edge
}

func (g *AdjGraph) Adj(v interface{}) []interface{} {
	return g.adjList[v.(int)].edges
}

func (g *AdjGraph) AddEdge(v, w interface{}) {
	if v.(int) >= g.V() || w.(int) >= g.V() {
		return
	}
	defer func() {
		g.edge++
	}()
	g.adjList[v.(int)].edges = append(g.adjList[v.(int)].edges, w.(int))
	g.adjList[w.(int)].edges = append(g.adjList[w.(int)].edges, v.(int))
}

type GraphResolver struct {
	g Graph
}

func NewGraphResolver(g Graph) *GraphResolver {
	if g == nil {
		return nil
	}
	return &GraphResolver{g}
}

func (u *GraphResolver) Degree(v int) int {
	return len(u.g.Adj(v))
}

func (u *GraphResolver) MaxDegree() int {
	var max int
	for i := 0; i < u.g.V(); i++ {
		if d := u.Degree(i); d > max {
			max = d
		}
	}
	return max
}

func (u *GraphResolver) AvgDegree(g Graph) float64 {
	v := g.V()
	if v == 0 {
		return 0
	}
	return 2 * float64(g.E()) / float64(g.V())
}

func (u *GraphResolver) NumberOfSelfLoops(g Graph) int {
	var amount int
	for i := 0; i < g.V(); i++ {
		for _, j := range g.Adj(i) {
			if i == j {
				amount++
			}
		}
	}
	return amount / 2
}

func (u *GraphResolver) Desc() string {
	desc := fmt.Sprintf("vertex:%d edge:%d\n", u.g.V(), u.g.E())
	for i := 0; i < u.g.V(); i++ {
		desc += fmt.Sprintf("%d connected to:%v\n", i, u.g.Adj(i))
	}
	return desc
}

type GraphSearcher interface {
	Marked(v, w int) bool //is v connected to w
	Count(v int) int      //amount of nodes connected to v
}

//Weighted quick-union with path compression.
type ufSearcher struct {
	g  Graph
	wf [][]int // count of each tree
	uf []int
}

func newUfSearcher(g Graph) *ufSearcher {
	us := &ufSearcher{g: g, uf: make([]int, g.V()), wf: make([][]int, g.V())}
	for i := 0; i < g.V(); i++ {
		us.uf[i] = i
		us.wf[i] = []int{i}
	}
	for v := 0; v < g.V(); v++ {
		for _, j := range g.Adj(v) {
			us.union(v, j.(int))
		}
	}
	return us
}

func (u *ufSearcher) union(v int, j int) {
	p := u.root(v)
	q := u.root(j)

	if p == q {
		return
	}
	if len(u.wf[p]) > len(u.wf[q]) || (len(u.wf[p]) == len(u.wf[q]) && p < q) {
		u.setRoot(j, p)
	} else {
		u.setRoot(v, q)
	}

}

func (u *ufSearcher) setRoot(v int, root int) {
	if v == root {
		return
	}
	defer func() {
		u.uf[v] = root
		u.wf[v] = []int{}
		if !utils.ContainsInt(u.wf[root], v) {
			u.wf[root] = append(u.wf[root], v)
		}
	}()
	parent := u.uf[v]
	if v == parent {
		return
	}
	u.setRoot(parent, root)
}

func (u *ufSearcher) root(w int) (root int) {
	return u.uf[w]
}

func (u *ufSearcher) Marked(v, w int) bool {
	return u.root(v) == u.root(w)
}

func (u *ufSearcher) Count(v int) int {
	return len(u.wf[u.root(v)]) - 1 //exclude self
}

type dfsSearcher struct {
	g        Graph
	marked   [][]int
	edgeTo   [][]int
	hasCycle bool
	isBlack  []bool
	twoColor bool
}

func newDfsSearcher(g Graph) *dfsSearcher {
	searcher := &dfsSearcher{g: g, isBlack: make([]bool, g.V()), twoColor: true}
	for i := 0; i < g.V(); i++ {
		var l []int
		var e []int
		for i := 0; i < g.V(); i++ {
			e = append(e, -1)
		}
		searcher.dfs(&l, &e, i, i)
		searcher.marked = append(searcher.marked, l)
		searcher.edgeTo = append(searcher.edgeTo, e)
	}
	return searcher
}

func (u *dfsSearcher) dfs(marked *[]int, edgeTo *[]int, w int, parent int) {
	*marked = append(*marked, w)
	for _, n := range u.g.Adj(w) {
		if !utils.ContainsInt(*marked, n.(int)) {
			(*edgeTo)[n.(int)] = w
			u.isBlack[n.(int)] = !u.isBlack[w]
			u.dfs(marked, edgeTo, n.(int), w)
		} else {
			if n != parent {
				u.hasCycle = true
			}
			if !u.isBlack[n.(int)] == u.isBlack[w] {
				u.twoColor = false
			}
		}
	}
}

func (u *dfsSearcher) isBipartite() bool {
	return u.twoColor
}

func (u *dfsSearcher) pathTo(s, v int) []int {
	edges := u.edgeTo[s]
	stack := &LinkedStack{}
	for v != s {
		stack.Push(v)
		path := edges[v]
		if path == -1 {
			return []int{}
		}
		v = path
	}
	stack.Push(s)
	var path []int
	for !stack.IsEmpty() {
		path = append(path, stack.Pop().(int))
	}
	return path
}

func (u *dfsSearcher) Marked(v, w int) bool {
	return utils.ContainsInt(u.marked[v], w)
}

func (u *dfsSearcher) Count(v int) int {
	return len(u.marked[v]) - 1
}

type bfsSearcher struct {
	g      Graph
	marked [][]int
	edgeTo [][]int
}

func newBfsSearcher(g Graph) *bfsSearcher {
	searcher := &bfsSearcher{g: g}
	for i := 0; i < g.V(); i++ {
		var l []int
		var e []int
		for i := 0; i < g.V(); i++ {
			e = append(e, -1)
		}
		searcher.bfs(&l, &e, i)
		searcher.marked = append(searcher.marked, l)
		searcher.edgeTo = append(searcher.edgeTo, e)
	}
	return searcher
}

func (u *bfsSearcher) bfs(marked *[]int, edgeTo *[]int, v int) {
	q := NewQueue()
	q.Enqueue(v)
	*marked = append(*marked, v)
	for {
		val := q.Dequeue()
		if val == nil {
			break
		}
		w := val.(int)
		for _, n := range u.g.Adj(w) {
			if !utils.ContainsInt(*marked, n.(int)) {
				(*edgeTo)[n.(int)] = w
				*marked = append(*marked, n.(int))
				q.Enqueue(n)
			}
		}

	}
}

func (u *bfsSearcher) pathTo(s, v int) []int {
	edges := u.edgeTo[s]
	stack := &LinkedStack{}
	for v != s {
		stack.Push(v)
		path := edges[v]
		if path == -1 {
			return []int{}
		}
		v = path
	}
	stack.Push(s)
	var path []int
	for !stack.IsEmpty() {
		path = append(path, stack.Pop().(int))
	}
	return path
}

func (u *bfsSearcher) Marked(v, w int) bool {
	return utils.ContainsInt(u.marked[v], w)
}

func (u *bfsSearcher) Count(v int) int {
	return len(u.marked[v]) - 1
}
